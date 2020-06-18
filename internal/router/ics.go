package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/memcachier/mc"
)

func getFeedICS(c *mc.Client, box *packr.Box) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, _, err := getFeed(c, box, ctx); err == nil {
			panic("not implemented")
		}
	}
}
