package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gobuffalo/packr/v2"

	"github.com/demosdemon/weather/pkg/meteonook"
)

func main() {
	box := packr.New("Data", "./data")

	bytes, err := box.Find("weather.wasm")
	if err != nil {
		log.Fatal(err)
	}

	instance, err := meteonook.NewInstance(bytes)
	if err != nil {
		log.Fatal(err)
	}

	island := meteonook.Island{
		Name:       "#island",
		Hemisphere: meteonook.Northern,
		Seed:       1160352294,
		Timezone:   meteonook.Timezone{Location: time.Local},
	}

	day, err := island.NewDay(instance, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(day); err != nil {
		log.Fatal(err)
	}
}
