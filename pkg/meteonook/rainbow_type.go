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
