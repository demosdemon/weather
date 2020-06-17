package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
)

func getFeedICS(box *packr.Box) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, _, err := getFeed(box, ctx); err == nil {
			panic("not implemented")
		}
	}
}
