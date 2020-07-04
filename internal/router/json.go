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
	"net/http"

	"github.com/demosdemon/weather/pkg/meteonook"
)

func getFeedJSON(w http.ResponseWriter, r *http.Request) {
	days, _, err := getFeed(r)
	if err != nil {
		writeError(w, err)
		return
	}

	data := make(map[string]*meteonook.Day, len(days))
	for _, day := range days {
		data[fmt.Sprintf("%04d-%02d-%02d", day.Year, day.Month, day.Date)] = day
	}

	writeJSON(w, http.StatusOK, data)
}

func getDateJSON(w http.ResponseWriter, r *http.Request) {
	day, _, err := getDate(r)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, day)
}
