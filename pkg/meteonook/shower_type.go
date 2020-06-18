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
