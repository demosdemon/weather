package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"

	"github.com/demosdemon/weather/pkg/meteonook"
)

func getFeed(box *packr.Box, ctx *gin.Context) ([]*meteonook.Day, *time.Location, error) {
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

	var hemisphere meteonook.Hemisphere
	switch query.Hemisphere {
	case "N", "n":
		hemisphere = meteonook.Northern
	case "S", "s":
		hemisphere = meteonook.Southern
	default:
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("invalid hemisphere", nil)).SetType(gin.ErrorTypePublic)
	}

	bytes, err := box.Find("weather.wasm")
	if err != nil {
		return nil, loc, ctx.AbortWithError(http.StatusInternalServerError, newError("unable to load weather engine", err)).SetType(gin.ErrorTypePublic)
	}

	island := meteonook.Island{
		Name:       query.IslandName,
		Hemisphere: hemisphere,
		Seed:       query.Seed,
		Timezone:   meteonook.Timezone{Location: loc},
	}

	instance, err := meteonook.NewInstance(bytes)
	if err != nil {
		return nil, loc, ctx.AbortWithError(http.StatusInternalServerError, newError("unable to instantiate weather engine", err)).SetType(gin.ErrorTypePublic)
	}
	defer instance.Close()

	numDays := int(last.Sub(first) / oneDay)
	days := make([]*meteonook.Day, 0, numDays)
	for first.Before(last) {
		day, err := island.NewDay(instance, first)
		if err != nil {
			return nil, loc, ctx.AbortWithError(http.StatusInternalServerError, newError("error with weather engine", err)).SetType(gin.ErrorTypePublic)
		}
		days = append(days, day)
		first = first.Add(oneDay)
	}

	return days, loc, nil
}
