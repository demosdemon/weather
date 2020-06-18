package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/memcachier/mc"

	"github.com/demosdemon/weather/pkg/meteonook"
)

func getFeedJSON(c *mc.Client, box *packr.Box) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if days, _, err := getFeed(c, box, ctx); err == nil {
			var data = make(map[string]*meteonook.Day, len(days))
			for _, day := range days {
				data[fmt.Sprintf("%04d-%02d-%02d", day.Year, day.Month, day.Date)] = day
			}
			ctx.JSON(http.StatusOK, data)
		}
	}
}
