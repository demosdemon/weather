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
