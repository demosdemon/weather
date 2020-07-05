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
	"net/http"

	"github.com/demosdemon/weather/pkg/meteonook"
)

type FeedResponse struct {
	Island *meteonook.Island `json:"island"`
	Days   []*meteonook.Day  `json:"days"`
}

func getFeedJSON(w http.ResponseWriter, r *http.Request) {
	island, days, _, err := getFeed(r)
	if err != nil {
		writeError(w, err)
		return
	}

	res := FeedResponse{
		Island: island,
		Days:   days,
	}

	writeJSON(w, http.StatusOK, res)
}

func getDateJSON(w http.ResponseWriter, r *http.Request) {
	_, day, _, err := getDate(r)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, day)
}
