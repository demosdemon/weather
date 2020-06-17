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
