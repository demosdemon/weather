package meteonook

import (
	"encoding/json"
	"fmt"
)

type FogType int

const (
	NoFog FogType = iota
	HeavyFog
	WaterFog
)

var fogTypeStrings = map[FogType]string{
	NoFog:    "No Fog",
	HeavyFog: "Heavy Fog",
	WaterFog: "Water Fog",
}

func (t FogType) String() string {
	if v, ok := fogTypeStrings[t]; ok {
		return v
	}

	return fmt.Sprintf("FogType(%d)", t)
}

func (t FogType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

var preNormalFogPatterns = map[WeatherPattern]bool{
	Fine00:      true,
	Fine01:      true,
	Fine02:      true,
	Fine03:      true,
	Fine04:      true,
	Fine05:      true,
	Fine06:      true,
	FineCloud00: true,
	FineCloud01: true,
	FineCloud02: true,
	CloudFine00: true,
	CloudFine01: true,
	CloudFine02: true,
	FineRain00:  true,
	FineRain01:  true,
	FineRain02:  true,
	FineRain03:  true,
	EventDay00:  true,
}

var preWaterFogPatterns = map[WeatherPattern]bool{
	Fine00:      true,
	Fine01:      true,
	Fine02:      true,
	Fine03:      true,
	Fine04:      true,
	Fine05:      true,
	Fine06:      true,
	FineCloud00: true,
	FineCloud01: true,
	FineCloud02: true,
	CloudFine00: true,
	CloudFine01: true,
	CloudFine02: true,
	FineRain00:  true,
	FineRain01:  true,
	FineRain02:  true,
	FineRain03:  true,
}

var fogPatterns = map[WeatherPattern]bool{
	Cloud00:     true,
	Cloud01:     true,
	Cloud02:     true,
	Rain00:      true,
	Rain01:      true,
	Rain02:      true,
	Rain03:      true,
	Rain04:      true,
	Rain05:      true,
	FineCloud00: true,
	FineCloud01: true,
	FineCloud02: true,
	CloudFine00: true,
	CloudFine01: true,
	CloudFine02: true,
	CloudRain00: true,
	CloudRain01: true,
	CloudRain02: true,
	RainCloud00: true,
	RainCloud01: true,
	RainCloud02: true,
}
