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

func (h *Hemisphere) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	switch s {
	case "Northern":
		*h = Northern
	case "Southern":
		*h = Southern
	default:
		return fmt.Errorf("invalid Hemisphere: %s", s)
	}

	return nil
}
