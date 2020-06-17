package meteonook

import (
	"encoding/json"
	"fmt"
)

type Weather int32

const (
	Clear Weather = iota
	Sunny
	Cloudy
	StormClouds
	LightStorm
	HeavyStorm
)

var weatherStrings = map[Weather]string{
	Clear:       "Clear/Fine",
	Sunny:       "Sunny",
	Cloudy:      "Cloudy",
	StormClouds: "Storm Clouds",
	LightStorm:  "Light Storm",
	HeavyStorm:  "Heavy Storm",
}

func (s Weather) String() string {
	if v, ok := weatherStrings[s]; ok {
		return v
	}
	return fmt.Sprintf("Weather(%d)", s)
}

func (s Weather) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
