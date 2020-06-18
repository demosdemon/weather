package router

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/memcachier/mc"

	"github.com/demosdemon/weather/pkg/meteonook"
)

func cacheKey(q string) string {
	p, err := url.ParseQuery(q)
	if err == nil {
		return p.Encode()
	}
	return q
}

func getFeed(c *mc.Client, box *packr.Box, ctx *gin.Context) ([]*meteonook.Day, *time.Location, error) {
	var query FeedQuery
	if err := ctx.BindQuery(&query); err != nil {
		// already aborted
		return nil, nil, err
	}

	loc, err := time.LoadLocation(query.Timezone)
	if err != nil {
		return nil, nil, ctx.AbortWithError(http.StatusBadRequest, newError("invalid timezone", err))
	}

	key := cacheKey(ctx.Request.URL.RawQuery)
	v, _, _, err := c.Get(key)
	if err != nil {
		log.Printf("error fetching from cache: %v", err)
	}

	if v == "" {
		log.Printf("cache miss for %s", key)
	} else {
		var days []*meteonook.Day
		err := json.Unmarshal([]byte(v), &days)
		return days, loc, err
	}

	const oneDay = time.Hour * 24
	today := time.Now().In(loc).Truncate(oneDay)

	first, err := query.first()
	if err != nil {
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("invalid first_date", err))
	}
	if first.IsZero() {
		first = today.AddDate(0, -3, 0)
	}

	last, err := query.last()
	if err != nil {
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("invalid last_date", err))
	}
	if last.IsZero() {
		last = first.AddDate(1, 0, 0)
	}
	last = last.Add(oneDay)

	if last.Before(first) {
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("last is before first", nil))
	}

	var hemisphere meteonook.Hemisphere
	switch query.Hemisphere {
	case "N", "n":
		hemisphere = meteonook.Northern
	case "S", "s":
		hemisphere = meteonook.Southern
	default:
		return nil, loc, ctx.AbortWithError(http.StatusBadRequest, newError("invalid hemisphere", nil))
	}

	bytes, err := box.Find("weather.wasm")
	if err != nil {
		return nil, loc, ctx.AbortWithError(http.StatusInternalServerError, newError("unable to load weather engine", err))
	}

	island := meteonook.Island{
		Name:       query.IslandName,
		Hemisphere: hemisphere,
		Seed:       query.Seed,
		Timezone:   meteonook.Timezone{Location: loc},
	}

	instance, err := meteonook.NewInstance(bytes)
	if err != nil {
		return nil, loc, ctx.AbortWithError(http.StatusInternalServerError, newError("unable to instantiate weather engine", err))
	}
	defer instance.Close()

	days := make([]*meteonook.Day, 0, int(last.Sub(first)/oneDay))
	for first.Before(last) {
		day, err := island.NewDay(instance, first)
		if err != nil {
			return nil, loc, ctx.AbortWithError(http.StatusInternalServerError, newError("error with weather engine", err))
		}
		days = append(days, day)
		first = first.Add(oneDay)
	}

	cache, _ := json.Marshal(days)
	if _, err := c.Set(key, string(cache), 0, 0, 0); err != nil {
		log.Printf("error saving %s to cache: %v", key, err)
	}

	return days, loc, nil
}
