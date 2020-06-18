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

func (s ShowerType) String() string {
	if v, ok := showerStrings[s]; ok {
		return v
	}
	return fmt.Sprintf("ShowerType(%d)", s)
}

func (s ShowerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
