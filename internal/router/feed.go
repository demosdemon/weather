package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"

	"github.com/demosdemon/weather/pkg/meteonook"
)

func getFeed(box *packr.Box, c *gin.Context) ([]*meteonook.Day, *time.Location, error) {
	var query FeedQuery
	if err := c.BindQuery(&query); err != nil {
		// already aborted
		return nil, nil, err
	}

	loc, err := time.LoadLocation(query.Timezone)
	if err != nil {
		return nil, nil, c.AbortWithError(http.StatusBadRequest, newError("invalid timezone", err))
	}

	const oneDay = time.Hour * 24
	today := time.Now().In(loc).Truncate(oneDay)

	first := query.FirstDate
	if first.IsZero() {
		first = today.AddDate(0, -3, 0)
	}

	last := query.LastDate
	if last.IsZero() {
		last = first.AddDate(1, 0, 0)
	}
	last = last.Add(oneDay)

	if last.Before(first) {
		return nil, loc, c.AbortWithError(http.StatusBadRequest, newError("last is before first", nil))
	}

	var hemisphere meteonook.Hemisphere
	switch query.Hemisphere {
	case "N", "n":
		hemisphere = meteonook.Northern
	case "S", "s":
		hemisphere = meteonook.Southern
	default:
		return nil, loc, c.AbortWithError(http.StatusBadRequest, newError("invalid hemisphere", nil))
	}

	bytes, err := box.Find("weather.wasm")
	if err != nil {
		return nil, loc, c.AbortWithError(http.StatusInternalServerError, newError("unable to load weather engine", err))
	}

	island := meteonook.Island{
		Name:       query.IslandName,
		Hemisphere: hemisphere,
		Seed:       query.Seed,
		Timezone:   meteonook.Timezone{Location: loc},
	}

	instance, err := meteonook.NewInstance(bytes)
	if err != nil {
		return nil, loc, c.AbortWithError(http.StatusInternalServerError, newError("unable to instantiate weather engine", err))
	}

	days := make([]*meteonook.Day, 0, int(last.Sub(first)/oneDay))
	for first.Before(last) {
		day, err := island.NewDay(instance, first)
		if err != nil {
			return nil, loc, c.AbortWithError(http.StatusInternalServerError, newError("error with weather engine", err))
		}
		days = append(days, day)
		first = first.Add(oneDay)
	}
	return days, loc, nil
}
