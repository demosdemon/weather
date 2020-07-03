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

type WeatherPattern int32

const (
	Fine00 WeatherPattern = iota
	Fine01
	Fine02
	Fine03
	Fine04
	Fine05
	Fine06
	Cloud00
	Cloud01
	Cloud02
	Rain00
	Rain01
	Rain02
	Rain03
	Rain04
	Rain05
	FineCloud00
	FineCloud01
	FineCloud02
	CloudFine00
	CloudFine01
	CloudFine02
	FineRain00
	FineRain01
	FineRain02
	FineRain03
	CloudRain00
	CloudRain01
	CloudRain02
	RainCloud00
	RainCloud01
	RainCloud02
	Common00
	EventDay00
)

var weatherPatternStrings = map[WeatherPattern]string{
	Fine00:      "Fine00",
	Fine01:      "Fine01",
	Fine02:      "Fine02",
	Fine03:      "Fine03",
	Fine04:      "Fine04",
	Fine05:      "Fine05",
	Fine06:      "Fine06",
	Cloud00:     "Cloud00",
	Cloud01:     "Cloud01",
	Cloud02:     "Cloud02",
	Rain00:      "Rain00",
	Rain01:      "Rain01",
	Rain02:      "Rain02",
	Rain03:      "Rain03",
	Rain04:      "Rain04",
	Rain05:      "Rain05",
	FineCloud00: "FineCloud00",
	FineCloud01: "FineCloud01",
	FineCloud02: "FineCloud02",
	CloudFine00: "CloudFine00",
	CloudFine01: "CloudFine01",
	CloudFine02: "CloudFine02",
	FineRain00:  "FineRain00",
	FineRain01:  "FineRain01",
	FineRain02:  "FineRain02",
	FineRain03:  "FineRain03",
	CloudRain00: "CloudRain00",
	CloudRain01: "CloudRain01",
	CloudRain02: "CloudRain02",
	RainCloud00: "RainCloud00",
	RainCloud01: "RainCloud01",
	RainCloud02: "RainCloud02",
	Common00:    "Common00",
	EventDay00:  "EventDay00",
}

var stringWeatherPatterns map[string]WeatherPattern

func init() {
	stringWeatherPatterns = make(map[string]WeatherPattern, len(weatherPatternStrings))
	for k, v := range weatherPatternStrings {
		stringWeatherPatterns[v] = k
	}
}

func (wp WeatherPattern) String() string {
	if v, ok := weatherPatternStrings[wp]; ok {
		return v
	}

	return fmt.Sprintf("WeatherPattern(%d)", wp)
}

func (wp WeatherPattern) MarshalJSON() ([]byte, error) {
	s := wp.String()
	return json.Marshal(s)
}

func (wp *WeatherPattern) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	if v, ok := stringWeatherPatterns[s]; ok {
		*wp = v
		return nil
	}

	return fmt.Errorf("invalid WeatherPattern: %s", s)
}

func (wp WeatherPattern) ShowerType() ShowerType {
	switch wp {
	case Fine00:
		return Heavy
	case Fine02, Fine04, Fine06:
		return Light
	default:
		return NoShower
	}
}
