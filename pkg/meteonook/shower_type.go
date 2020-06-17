package meteonook

import (
	"encoding/json"
	"fmt"
)

type ShowerType int32

const (
	NotSure ShowerType = iota
	NoShower
	Light
	Heavy
)

var showerStrings = map[ShowerType]string{
	NotSure:  "Not Sure",
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
