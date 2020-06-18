package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/memcachier/mc"
)

func getFeedJSON(c *mc.Client, box *packr.Box) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if days, _, err := getFeed(c, box, ctx); err == nil {
			ctx.JSON(http.StatusOK, days)
		}
	}
}
