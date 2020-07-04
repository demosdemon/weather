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
	"time"

	"github.com/gin-gonic/gin"

	"github.com/demosdemon/weather/pkg/meteonook"
	"github.com/demosdemon/weather/pkg/meteonook/enums"
)

func getFeed(ctx *gin.Context) ([]*meteonook.Day, *time.Location, error) {
	var query FeedQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		return nil, nil, ctx.AbortWithError(http.StatusBadRequest, newError("invalid query", err)).SetType(gin.ErrorTypePublic)
	}

	loc, err := time.LoadLocation(query.Timezone)
	if err != nil {
		return nil, nil, ctx.AbortWithError(http.StatusBadRequest, newError("invalid timezone", err)).SetType(gin.ErrorTypePublic)
	}

	const oneDay = time.Hour * 24
	today := time.Now().In(loc).Truncate(oneDay)

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

	var hemisphere enums.Hemisphere
	switch query.Hemisphere {
	case "N", "n":
		hemisphere = enums.Northern
	case "S", "s":
		hemisphere = enums.Southern
	default:
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("invalid hemisphere", nil)).SetType(gin.ErrorTypePublic)
	}

	island := meteonook.Island{
		Name:       query.IslandName,
		Hemisphere: hemisphere,
		Seed:       query.Seed,
		Timezone:   meteonook.Timezone{Location: loc},
	}

	numDays := int(last.Sub(first) / oneDay)
	days := make([]*meteonook.Day, 0, numDays)
	for first.Before(last) {
		day, err := island.NewDay(first)
		if err != nil {
			return nil, loc, ctx.AbortWithError(http.StatusInternalServerError, newError("error with weather engine", err)).SetType(gin.ErrorTypePublic)
		}
		days = append(days, day)
		first = first.Add(oneDay)
	}

	return days, loc, nil
}
