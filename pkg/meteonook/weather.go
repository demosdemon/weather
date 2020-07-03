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
