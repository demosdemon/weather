package router

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/memcachier/mc"
)

func NewRouter(box *packr.Box) *gin.Engine {
	c := mc.NewMC(os.Getenv("MEMCACHIER_SERVERS"), os.Getenv("MEMCACHIER_USERNAME"), os.Getenv("MEMCACHIER_PASSWORD"))

	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	r.GET("/feed.json", getFeedJSON(c, box))
	r.GET("/feed.ics", getFeedICS(c, box))

	return r
}

type FeedQuery struct {
	IslandName string `form:"island_name"`
	Hemisphere string `form:"hemisphere" binding:"required"`
	Seed       int32  `form:"seed" binding:"required"`
	Timezone   string `form:"timezone"`
	FirstDate  string `form:"first_date"`
	LastDate   string `form:"last_date"`
}

func (f FeedQuery) first() (time.Time, error) {
	if f.FirstDate == "" {
		return time.Time{}, nil
	}

	return time.Parse("2006-01-02", f.FirstDate)
}

func (f FeedQuery) last() (time.Time, error) {
	if f.LastDate == "" {
		return time.Time{}, nil
	}

	return time.Parse("2006-01-02", f.LastDate)
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
