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

type SpecialDay int32

const (
	RegularDay SpecialDay = iota
	BunnyDay
	FishingTourney
	BugOff
	NewYearsEve
)

var specialDayStrings = map[SpecialDay]string{
	RegularDay:     "Regular Day",
	BunnyDay:       "Bunny Day",
	FishingTourney: "Fishing Tourney",
	BugOff:         "Bug-Off",
	NewYearsEve:    "New Year's Eve Countdown",
}

var stringSpecialDays map[string]SpecialDay

func init() {
	stringSpecialDays = make(map[string]SpecialDay, len(specialDayStrings))
	for k, v := range specialDayStrings {
		stringSpecialDays[v] = k
	}
}

func (s SpecialDay) String() string {
	if v, ok := specialDayStrings[s]; ok {
		return v
	}

	return fmt.Sprintf("SpecialDay(%d)", s)
}

func (s SpecialDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *SpecialDay) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	if v, ok := stringSpecialDays[str]; ok {
		*s = v
		return nil
	}

	return fmt.Errorf("invalid SpecialDay: %s", s)
}
