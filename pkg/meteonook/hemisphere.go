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

type Hemisphere int32

const (
	Northern Hemisphere = iota
	Southern
)

func (h Hemisphere) String() string {
	switch h {
	case Northern:
		return "Northern"
	case Southern:
		return "Southern"
	default:
		return fmt.Sprintf("Hemisphere(%d)", h)
	}
}

func (h Hemisphere) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}

func (h *Hemisphere) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	switch s {
	case "Northern":
		*h = Northern
	case "Southern":
		*h = Southern
	default:
		return fmt.Errorf("invalid Hemisphere: %s", s)
	}

	return nil
}
