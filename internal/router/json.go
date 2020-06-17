package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
)

func getFeedJSON(box *packr.Box) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if days, _, err := getFeed(box, ctx); err == nil {
			ctx.JSON(http.StatusOK, days)
		}
	}
}
