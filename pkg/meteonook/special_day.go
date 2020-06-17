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

func (s SpecialDay) String() string {
	if v, ok := specialDayStrings[s]; ok {
		return v
	}

	return fmt.Sprintf("SpecialDay(%d)", s)
}

func (s SpecialDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
