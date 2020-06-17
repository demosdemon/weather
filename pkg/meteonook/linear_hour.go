package meteonook

import (
	"encoding/json"
	"fmt"
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
