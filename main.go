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

package main

import (
	"log"
	"os"

	"github.com/gobuffalo/packr/v2"

	"github.com/demosdemon/weather/internal/router"
)

func main() {
	box := packr.New("Data", "./data")
	r := router.NewRouter(box)

	port := "3000"
	if v, ok := os.LookupEnv("PORT"); ok {
		port = v
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
