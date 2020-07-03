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

type ShowerType int32

const (
	NoShower ShowerType = iota
	Light
	Heavy
)

var showerStrings = map[ShowerType]string{
	NoShower: "None",
	Light:    "Light",
	Heavy:    "Heavy",
}

var stringShowers map[string]ShowerType

func init() {
	stringShowers = make(map[string]ShowerType, len(showerStrings))
	for k, v := range showerStrings {
		stringShowers[v] = k
	}
}

func (s ShowerType) String() string {
	if v, ok := showerStrings[s]; ok {
		return v
	}
	return fmt.Sprintf("ShowerType(%d)", s)
}

func (s ShowerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *ShowerType) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	if v, ok := stringShowers[str]; ok {
		*s = v
		return nil
	}

	return fmt.Errorf("invalid ShowerType: %v", str)
}
