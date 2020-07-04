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

package middleware

import (
	"fmt"
	"io"
	"net/http"

	"github.com/demosdemon/cpanic"
)

type recoveryHandler struct {
	h http.Handler
	w io.Writer
}

func (h recoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer cpanic.Recover(func(p *cpanic.Panic) {
		msg := p.String()
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", fmt.Sprint(len(msg)))
		w.WriteHeader(http.StatusInternalServerError)

		wr := io.MultiWriter(h.w, w)
		_, _ = fmt.Fprint(wr, msg)
	})

	h.h.ServeHTTP(w, r)
}

func Recover(w io.Writer) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return recoveryHandler{h, w}
	}
}
