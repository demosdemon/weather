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
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/demosdemon/weather/internal/router"
)

func main() {
	port := "3000"
	if v, ok := os.LookupEnv("PORT"); ok {
		port = v
	}

	srv := http.Server{
		Addr:    ":" + port,
		Handler: router.NewRouter(os.Stdout),
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-ch
		log.Printf("signal: %v", sig)
		_ = srv.Close()
	}()

	log.Printf("listening on port %s", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	log.Println("fin")
}
