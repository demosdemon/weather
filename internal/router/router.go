package router

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
)

func NewRouter(box *packr.Box) *gin.Engine {
	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	r.GET("/feed.json", getFeedJSON(box))
	r.GET("/feed.ics", getFeedICS(box))

	return r
}

type FeedQuery struct {
	IslandName string    `form:"island_name"`
	Hemisphere string    `form:"hemisphere" binding:"required"`
	Seed       int32     `form:"seed" binding:"required"`
	Timezone   string    `form:"timezone"`
	FirstDate  time.Time `form:"first_date"`
	LastDate   time.Time `form:"last_date"`
}

type Error struct {
	Msg string `json:"error"`
	Err string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Msg, e.Err)
}

func newError(msg string, err error) Error {
	if err == nil {
		return Error{Msg: msg}
	}

	return Error{
		Msg: msg,
		Err: err.Error(),
	}
}
