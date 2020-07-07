// weather
// Copyright (C) 2020 Brandon LeBlanc <brandon@leblanc.codes>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package router

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"github.com/demosdemon/weather/pkg/meteonook"
	"github.com/demosdemon/weather/pkg/meteonook/enums"
)

const (
	RFC3339Date = "2006-01-02"
)

type FeedQuery struct {
	URL struct {
		Seed       uint32 `schema:"seed,required"`
		Hemisphere string `schema:"hemisphere,required"`
		Name       string `schema:"name,required"`
		Date       string `schema:"date"`
	}

	Query struct {
		Timezone  string `schema:"timezone"`
		FirstDate string `schema:"first_date"`
		LastDate  string `schema:"last_date"`
	}
}

func (f FeedQuery) tz() (*time.Location, error) {
	return time.LoadLocation(f.Query.Timezone)
}

func (f FeedQuery) parse(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}

	loc, err := f.tz()
	if err != nil {
		return time.Time{}, err
	}

	return time.ParseInLocation(RFC3339Date, s, loc)
}

func (f FeedQuery) date() (time.Time, error) {
	date, err := f.parse(f.URL.Date)
	if err != nil {
		return date, err
	}
	if date.IsZero() {
		tz, _ := f.tz()
		date = truncate(time.Now().In(tz))
	}
	return date, nil
}

func (f FeedQuery) first() (date time.Time, err error) {
	if date, err = f.parse(f.Query.FirstDate); err != nil {
		return date, err
	}
	if date.IsZero() {
		if date, err = f.date(); err != nil {
			return date, err
		}
		date = date.AddDate(0, -3, 0)
	}
	return date, nil
}

func (f FeedQuery) last() (date time.Time, err error) {
	if date, err = f.parse(f.Query.LastDate); err != nil {
		return date, err
	}
	if date.IsZero() {
		if date, err = f.first(); err != nil {
			return date, err
		}
		date = date.AddDate(1, 0, 0)
	}
	return date, nil
}

func vars(r *http.Request) url.Values {
	v := mux.Vars(r)
	rv := make(url.Values, len(v))
	for k, v := range v {
		rv.Set(k, v)
	}
	return rv
}

func getQuery(r *http.Request) (*FeedQuery, *time.Location, *meteonook.Island, error) {
	var query FeedQuery

	dec := schema.NewDecoder()
	if err := dec.Decode(&query.URL, vars(r)); err != nil {
		return nil, nil, nil, newError(http.StatusBadRequest, "invalid url parameters", err)
	}
	if err := dec.Decode(&query.Query, r.URL.Query()); err != nil {
		return nil, nil, nil, newError(http.StatusBadRequest, "invalid query parameters", err)
	}

	loc, err := query.tz()
	if err != nil {
		return nil, loc, nil, newError(http.StatusBadRequest, "invalid timezone", err)
	}

	var hemisphere enums.Hemisphere
	switch query.URL.Hemisphere {
	case "N", "n":
		hemisphere = enums.Northern
	case "S", "s":
		hemisphere = enums.Southern
	default:
		return nil, loc, nil, newError(http.StatusBadRequest, "invalid hemisphere", nil)
	}

	island := meteonook.Island{
		Name:       query.URL.Name,
		Hemisphere: hemisphere,
		Seed:       query.URL.Seed,
		Timezone:   meteonook.Timezone{Location: loc},
	}

	return &query, loc, &island, nil
}

func getDate(r *http.Request) (*meteonook.Island, *meteonook.Day, *time.Location, error) {
	query, loc, island, err := getQuery(r)
	if err != nil {
		return island, nil, loc, err
	}

	today, err := query.date()
	if err != nil {
		return island, nil, loc, newError(http.StatusBadRequest, "invalid date", err)
	}

	day, err := island.NewDay(today)
	if err != nil {
		return island, nil, loc, newError(http.StatusBadRequest, "error with weather engine", err)
	}

	return island, day, loc, nil
}

func getFeed(r *http.Request) (*meteonook.Island, []*meteonook.Day, *time.Location, error) {
	query, loc, island, err := getQuery(r)
	if err != nil {
		return island, nil, loc, err
	}

	first, err := query.first()
	if err != nil {
		return island, nil, loc, newError(http.StatusBadRequest, "invalid first_date", err)
	}

	last, err := query.last()
	if err != nil {
		return island, nil, loc, newError(http.StatusBadRequest, "invalid last_date", err)
	}

	if last.Before(first) {
		return island, nil, loc, newError(http.StatusBadRequest, "last is before first", nil)
	}

	last = last.Add(time.Second)
	numDays := int(last.Sub(first) / (24 * time.Hour))
	days := make([]*meteonook.Day, 0, numDays)
	for first.Before(last) {
		day, err := island.NewDay(first)
		if err != nil {
			return island, nil, loc, newError(http.StatusBadRequest, "error with weather engine", err)
		}
		days = append(days, day)
		first = first.AddDate(0, 0, 1)
	}

	return island, days, loc, nil
}

func truncate(date time.Time) time.Time {
	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		0,
		0,
		0,
		0,
		date.Location(),
	)
}
