// weather
// Copyright (C) 2020 Brandon LeBlanc <brandon@leblanc.codes>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package router

import (
	"fmt"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.ErrorLogger(),
		gin.Recovery(),
		gzip.Gzip(gzip.BestCompression),
	)

	r.GET("/feed.json", getFeedJSON)
	r.GET("/feed.ics", getFeedICS)
	r.GET("/today.json", getDateJSON)

	r.GET("/seed/:seed/date/:date/feed.json", getDateJSON)
	r.GET("/seed/:seed/feed.json", getFeedJSON)
	r.GET("/seed/:seed/feed.ics", getFeedICS)

	return r
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
