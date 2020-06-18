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

var stringWeathers map[string]Weather

func init() {
	stringWeathers = make(map[string]Weather, len(weatherStrings))
	for k, v := range weatherStrings {
		stringWeathers[v] = k
	}
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

func (s *Weather) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	if v, ok := stringWeathers[str]; ok {
		*s = v
		return nil
	}

	return fmt.Errorf("invalid Weather: %s", str)
}
