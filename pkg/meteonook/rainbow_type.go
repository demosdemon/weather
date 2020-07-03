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

type RainbowType int

const (
	NoRainbow RainbowType = iota
	SingleRainbow
	DoubleRainbow
)

var rainbowStrings = map[RainbowType]string{
	NoRainbow:     "No Rainbow",
	SingleRainbow: "Single Rainbow",
	DoubleRainbow: "Double Rainbow",
}

var stringRainbows map[string]RainbowType

func init() {
	stringRainbows = make(map[string]RainbowType, len(rainbowStrings))
	for k, v := range rainbowStrings {
		stringRainbows[v] = k
	}
}

func (s RainbowType) String() string {
	if v, ok := rainbowStrings[s]; ok {
		return v
	}
	return fmt.Sprintf("RainbowType(%d)", s)
}

func (s RainbowType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *RainbowType) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	if v, ok := stringRainbows[str]; ok {
		*s = v
		return nil
	}

	return fmt.Errorf("invalid RainbowType: %s", str)
}
