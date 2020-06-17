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

func (s RainbowType) String() string {
	if v, ok := rainbowStrings[s]; ok {
		return v
	}
	return fmt.Sprintf("RainbowType(%d)", s)
}

func (s RainbowType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
