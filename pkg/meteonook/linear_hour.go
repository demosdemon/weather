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
	"time"
)

type LinearHour int32

const (
	FiveAM LinearHour = iota
	SixAM
	SevenAM
	EightAM
	NineAM
	TenAM
	ElevenAM
	TwelvePM
	OnePM
	TwoPM
	ThreePM
	FourPM
	FivePM
	SixPM
	SevenPM
	EightPM
	NinePM
	TenPM
	ElevenPM
	TwelveAM
	OneAM
	TwoAM
	ThreeAM
	FourAM
)

func (hour LinearHour) String() string {
	x := hour.RegularHour()
	return fmt.Sprintf("%02d:00", x)
}

func (hour LinearHour) TwelveHourString() string {
	x := hour.RegularHour()
	v := x
	if x == 0 || x == 12 {
		v = 12
	} else {
		v %= 12
	}

	ampm := "AM"
	if x >= 12 {
		ampm = "PM"
	}

	return fmt.Sprintf("%02d:00 %s", v, ampm)
}

func NewLinearHour(hour int32) LinearHour {
	if 5 > hour {
		return LinearHour(19 + hour)
	}

	return LinearHour(hour - 5)
}

func (hour LinearHour) RegularHour() int32 {
	if 19 <= hour {
		return int32(hour - 19)
	}

	return int32(hour + 5)
}

func (hour LinearHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(hour.String())
}

func (hour *LinearHour) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	t, err := time.Parse("15:04", s)
	if err != nil {
		return err
	}

	*hour = NewLinearHour(int32(t.Hour()))
	return nil
}
