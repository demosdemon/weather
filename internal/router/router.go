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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/demosdemon/weather/pkg/middleware"
)

func get(fn http.HandlerFunc) http.Handler {
	return handlers.MethodHandler{http.MethodGet: fn}
}

func NewRouter(w io.Writer) http.Handler {
	r := mux.NewRouter()
	r.Use(
		middleware.Log(w),
		handlers.CompressHandler,
		middleware.Recover(w),
	)

	s := r.PathPrefix("/v1/{seed:[0-9]+}/{hemisphere:[NS]}/{name}").Subrouter()
	s.Handle("/feed.json", get(getFeedJSON))
	s.Handle("/feed.ics", get(getFeedICS))
	s.Handle("/today.json", get(getDateJSON))
	s.Handle("/{date:[0-9]{4}-[0-9]{2}-[0-9]{2}}.json", get(getDateJSON))

	return r
}

func writeError(w http.ResponseWriter, err error) {
	if err, ok := err.(Error); ok {
		writeJSON(w, err.Code, err)
		return
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	var buf bytes.Buffer
	dec := json.NewEncoder(&buf)
	dec.SetIndent("", "  ")
	err := dec.Encode(v)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	w.WriteHeader(code)
	_, _ = io.Copy(w, &buf)
}

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"error"`
	Err  string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Msg, e.Err)
}

func newError(code int, msg string, err error) Error {
	if err == nil {
		return Error{Code: code, Msg: msg}
	}

	return Error{
		Code: code,
		Msg:  msg,
		Err:  err.Error(),
	}
}
