package meteonook

import (
	"encoding/json"
	"time"
)

type Timezone struct {
	*time.Location
}

func (tz Timezone) MarshalJSON() ([]byte, error) {
	return json.Marshal(tz.Location.String())
}

func (tz *Timezone) UnmarshalJSON(data []byte) error {
	var name string
	err := json.Unmarshal(data, &name)
	if err != nil {
		return err
	}
	tz.Location, err = time.LoadLocation(name)
	return err
}
