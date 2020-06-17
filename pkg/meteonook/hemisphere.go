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
