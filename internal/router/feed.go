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
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/demosdemon/weather/pkg/meteonook"
	"github.com/demosdemon/weather/pkg/meteonook/enums"
)

const oneDay = time.Hour * 24

type FeedQuery struct {
	IslandName string `form:"island_name"`
	Hemisphere string `form:"hemisphere" binding:"required"`
	Seed       uint32 `form:"seed"`
	Timezone   string `form:"timezone"`
	Date       string `form:"date"`
	FirstDate  string `form:"first_date"`
	LastDate   string `form:"last_date"`
}

func (f FeedQuery) tz() (*time.Location, error) {
	return time.LoadLocation(f.Timezone)
}

func (f FeedQuery) parse(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}

	loc, err := f.tz()
	if err != nil {
		return time.Time{}, err
	}

	return time.ParseInLocation("2006-01-02", s, loc)
}

func (f FeedQuery) date() (time.Time, error) {
	date, err := f.parse(f.Date)
	if err != nil {
		return date, err
	}
	if date.IsZero() {
		tz, _ := f.tz()
		date = time.Now().In(tz).Truncate(oneDay)
	}
	return date, nil
}

func (f FeedQuery) first() (time.Time, error) {
	return f.parse(f.FirstDate)
}

func (f FeedQuery) last() (time.Time, error) {
	return f.parse(f.LastDate)
}

func getQuery(ctx *gin.Context) (*FeedQuery, *time.Location, *meteonook.Island, error) {
	var query FeedQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		return nil, nil, nil, ctx.AbortWithError(http.StatusBadRequest, newError("invalid query", err)).SetType(gin.ErrorTypePublic)
	}

	s := ctx.Param("seed")
	if s == "" {
		s = ctx.Query("seed")
	}
	if s == "" {
		return nil, nil, nil, ctx.AbortWithError(http.StatusBadRequest, newError("missing seed", nil)).SetType(gin.ErrorTypePublic)
	}

	seed, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return nil, nil, nil, ctx.AbortWithError(http.StatusBadRequest, newError("invalid seed", err)).SetType(gin.ErrorTypePublic)
	}

	query.Seed = uint32(seed)

	loc, err := time.LoadLocation(query.Timezone)
	if err != nil {
		return nil, nil, nil, ctx.AbortWithError(http.StatusBadRequest, newError("invalid timezone", err)).SetType(gin.ErrorTypePublic)
	}

	var hemisphere enums.Hemisphere
	switch query.Hemisphere {
	case "N", "n":
		hemisphere = enums.Northern
	case "S", "s":
		hemisphere = enums.Southern
	default:
		return nil, loc, nil, ctx.AbortWithError(http.StatusBadRequest, newError("invalid hemisphere", nil)).SetType(gin.ErrorTypePublic)
	}

	island := meteonook.Island{
		Name:       query.IslandName,
		Hemisphere: hemisphere,
		Seed:       query.Seed,
		Timezone:   meteonook.Timezone{Location: loc},
	}

	return &query, loc, &island, nil
}

func getDate(ctx *gin.Context) (*meteonook.Day, *time.Location, error) {
	query, loc, island, err := getQuery(ctx)
	if err != nil {
		return nil, loc, err
	}

	today, err := query.date()
	if err != nil {
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("invalid date", err)).SetType(gin.ErrorTypePublic)
	}

	day, err := island.NewDay(today)
	if err != nil {
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("error with weather engine", err)).SetType(gin.ErrorTypePublic)
	}

	return day, loc, nil
}

func getFeed(ctx *gin.Context) ([]*meteonook.Day, *time.Location, error) {
	query, loc, island, err := getQuery(ctx)
	if err != nil {
		return nil, loc, err
	}

	today, err := query.date()
	if err != nil {
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("invalid date", err)).SetType(gin.ErrorTypePublic)
	}

	first, err := query.first()
	if err != nil {
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("invalid first_date", err)).SetType(gin.ErrorTypePublic)
	}
	if first.IsZero() {
		first = today.AddDate(0, -3, 0)
	}

	last, err := query.last()
	if err != nil {
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("invalid last_date", err)).SetType(gin.ErrorTypePublic)
	}
	if last.IsZero() {
		last = first.AddDate(1, 0, 0)
	}
	last = last.Add(oneDay)
	if last.Before(first) {
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("last is before first", nil)).SetType(gin.ErrorTypePublic)
	}

	numDays := int(last.Sub(first) / oneDay)
	days := make([]*meteonook.Day, 0, numDays)
	for first.Before(last) {
		day, err := island.NewDay(first)
		if err != nil {
			return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("error with weather engine", err)).SetType(gin.ErrorTypePublic)
		}
		days = append(days, day)
		first = first.Add(oneDay)
	}

	return days, loc, nil
}
