package meteonook

import (
	"errors"
	"time"
)

const (
	minYear = 2000
	maxYear = 2060
)

// WeatherPattern is EventDay00 when SpecialDay > RegularDay
// Fog is visible when FogType > NoFog && SixAM < LinearHour < TenAM
// Aurora is visible when AuroraPossible is true && WeatherPattern == Fine00 && SixPM < LinearHour < ThreeAM
// Show is visible when SnowPossible is true && Weather > StormClouds

type Island struct {
	Name       string     `json:"name,omitempty"`
	Hemisphere Hemisphere `json:"hemisphere"`
	Seed       int32      `json:"seed"`
	Timezone   Timezone   `json:"timezone"`
}

type Day struct {
	Island         *Island        `json:"island"`
	Year           int32          `json:"year"`
	Month          time.Month     `json:"month"`
	Date           int32          `json:"date"`
	Weekday        time.Weekday   `json:"weekday"`
	SpecialDay     SpecialDay     `json:"special_day,omitempty"`
	WeatherPattern WeatherPattern `json:"weather_pattern"`
	ShowerType     ShowerType     `json:"shower_type,omitempty"`
	FogType        FogType        `json:"fog_type,omitempty"`
	RainbowType    RainbowType    `json:"rainbow_type,omitempty"`
	AuroraPossible bool           `json:"aurora_possible,omitempty"`
	SnowPossible   bool           `json:"snow_possible,omitempty"`
	Hours          [24]Hour       `json:"hours"`
}

type Hour struct {
	Hour          LinearHour  `json:"hour"`
	Weather       Weather     `json:"weather"`
	WindPower     int32       `json:"wind_power,omitempty"`
	ShowerType    ShowerType  `json:"shower_type,omitempty"`
	FogType       FogType     `json:"fog_type,omitempty"`
	RainbowType   RainbowType `json:"rainbow_type,omitempty"`
	AuroraVisible bool        `json:"aurora_visible,omitempty"`
	SnowVisible   bool        `json:"snow_visible,omitempty"`
	ShootingStars []time.Time `json:"shooting_stars,omitempty"`
}

func (island *Island) NewDay(instance *Instance, ts time.Time) (*Day, error) {
	const oneDay = time.Hour * 24
	year, month, date := ts.Date()
	if year < minYear || maxYear < year {
		return nil, errors.New("the provided time is outside of the valid range [2000, 2060]")
	}

	day := &Day{
		Island:  island,
		Year:    int32(year),
		Month:   month,
		Date:    int32(date),
		Weekday: ts.Weekday(),
	}

	ts = ts.Truncate(oneDay)
	yts := ts.Add(-oneDay)

	yesterday, err := getWeatherPattern(instance, island.Seed, island.Hemisphere, yts.Year(), yts.Month(), yts.Day())
	if err != nil {
		return nil, err
	}

	if day.SpecialDay, err = getSpecialDay(instance, island.Hemisphere, year, month, date); err != nil {
		return nil, err
	}
	if day.WeatherPattern, err = getWeatherPattern(instance, island.Seed, island.Hemisphere, year, month, date); err != nil {
		return nil, err
	}
	day.ShowerType = day.WeatherPattern.ShowerType()
	if day.FogType, err = getFogType(instance, island.Seed, island.Hemisphere, year, month, date, day.WeatherPattern, yesterday); err != nil {
		return nil, err
	}
	var rainbowHour LinearHour
	if day.RainbowType, rainbowHour, err = getRainbowType(instance, island.Seed, island.Hemisphere, year, month, date, day.WeatherPattern); err != nil {
		return nil, err
	}
	if day.AuroraPossible, err = getAuroraPossible(instance, island.Seed, island.Hemisphere, year, month, date, day.WeatherPattern); err != nil {
		return nil, err
	}
	if day.SnowPossible, err = getSnowPossible(instance, island.Hemisphere, month, date); err != nil {
		return nil, err
	}

	for idx := range day.Hours {
		hour := LinearHour(idx)
		weather, err := getWeather(instance, day.WeatherPattern, hour)
		if err != nil {
			return nil, err
		}
		windPower, err := getWindPower(instance, island.Seed, island.Hemisphere, year, month, date, hour, day.WeatherPattern)
		if err != nil {
			return nil, err
		}
		shootingStarsPossible, err := getShootingStarsPossible(instance, day.WeatherPattern, hour)
		showerType := day.ShowerType
		if !shootingStarsPossible {
			showerType = NoShower
		}
		fogType := NoFog
		if SixAM < hour && hour < TenAM {
			fogType = day.FogType
		}
		rainbowType := NoRainbow
		if rainbowHour == hour || rainbowHour == hour+1 {
			rainbowType = day.RainbowType
		}
		aurora := day.AuroraPossible && day.WeatherPattern == Fine00 && SixPM < hour && hour < ThreeAM
		snow := day.SnowPossible && weather > StormClouds
		var shootingStars []time.Time
		if showerType > NoShower {
			if shootingStars, err = getShootingStars(instance, island.Seed, year, month, date, hour, island.Timezone, day.WeatherPattern); err != nil {
				return nil, err
			}
		}

		day.Hours[idx] = Hour{
			Hour:          hour,
			Weather:       weather,
			WindPower:     windPower,
			ShowerType:    showerType,
			FogType:       fogType,
			RainbowType:   rainbowType,
			AuroraVisible: aurora,
			SnowVisible:   snow,
			ShootingStars: shootingStars,
		}
	}

	return day, nil
}

func getSpecialDay(
	instance *Instance,
	hemisphere Hemisphere,
	year int,
	month time.Month,
	date int,
) (SpecialDay, error) {
	v, err := instance.IsSpecialDay(int32(hemisphere), int32(year), int32(month), int32(date))
	return SpecialDay(v), err
}

func getWeatherPattern(
	instance *Instance,
	seed int32,
	hemisphere Hemisphere,
	year int,
	month time.Month,
	date int,
) (WeatherPattern, error) {
	v, err := instance.GetPattern(seed, int32(hemisphere), int32(year), int32(month), int32(date))
	return WeatherPattern(v), err
}

func getFogType(
	instance *Instance,
	seed int32,
	hemisphere Hemisphere,
	year int,
	month time.Month,
	date int,
	today WeatherPattern,
	yesterday WeatherPattern,
) (FogType, error) {
	v, err := instance.GetFog(int32(hemisphere), int32(month), int32(date))

	if err != nil {
		return 0, err
	}

	if v == 0 {
		return NoFog, nil
	}

	getWindPower := func(prev bool, hour int32) (bool, error) {
		if !prev {
			return prev, nil
		}

		v, err := instance.GetWindPower(int32(hemisphere), seed, int32(year), int32(month), int32(date), hour, int32(yesterday))
		return 3 > v, err
	}

	normalFog := preNormalFogPatterns[yesterday]
	normalFog = normalFog && fogPatterns[today]
	for _, hour := range []int32{5, 6, 7, 8} {
		if normalFog, err = getWindPower(normalFog, hour); err != nil {
			return 0, err
		}
	}

	if normalFog {
		return HeavyFog, nil
	}

	waterFog := preWaterFogPatterns[yesterday]
	waterFog = waterFog && fogPatterns[today]
	if waterFog {
		v, err = instance.CheckWaterFog(seed, int32(year), int32(month), int32(date))
		if err != nil {
			return 0, err
		}
		if v > 0 {
			return WaterFog, nil
		}
	}

	return NoFog, nil
}

func getRainbowType(
	instance *Instance,
	seed int32,
	hemisphere Hemisphere,
	year int,
	month time.Month,
	date int,
	pattern WeatherPattern,
) (RainbowType, LinearHour, error) {
	v, err := instance.IsRainbowPattern(int32(hemisphere), seed, int32(year), int32(month), int32(date), int32(pattern))
	if err != nil {
		return 0, 0, err
	}

	return RainbowType(v >> 8), NewLinearHour(v & 0xFF) + 1, nil
}

func getAuroraPossible(
	instance *Instance,
	seed int32,
	hemisphere Hemisphere,
	year int,
	month time.Month,
	date int,
	pattern WeatherPattern,
) (bool, error) {
	v, err := instance.IsAuroraPattern(int32(hemisphere), seed, int32(year), int32(month), int32(date), int32(pattern))
	return v > 0, err
}

func getSnowPossible(
	instance *Instance,
	hemisphere Hemisphere,
	month time.Month,
	date int,
) (bool, error) {
	v, err := instance.GetSnow(int32(hemisphere), int32(month), int32(date))
	return v > 0, err
}

func getWeather(
	instance *Instance,
	pattern WeatherPattern,
	hour LinearHour,
) (Weather, error) {
	v, err := instance.GetWeather(int32(pattern), hour.RegularHour())
	return Weather(v), err
}

func getWindPower(
	instance *Instance,
	seed int32,
	hemisphere Hemisphere,
	year int,
	month time.Month,
	date int,
	hour LinearHour,
	pattern WeatherPattern,
) (int32, error) {
	return instance.GetWindPower(int32(hemisphere), seed, int32(year), int32(month), int32(date), hour.RegularHour(), int32(pattern))
}

func getShootingStarsPossible(
	instance *Instance,
	pattern WeatherPattern,
	hour LinearHour,
) (bool, error) {
	v, err := instance.CanHaveShootingStars(int32(pattern), hour.RegularHour())
	return v > 0, err
}

func getShootingStars(
	instance *Instance,
	seed int32,
	year int,
	month time.Month,
	date int,
	hour LinearHour,
	tz Timezone,
	pattern WeatherPattern,
) ([]time.Time, error) {
	const nsec = 0
	loc := tz.Location
	if loc == nil {
		loc = time.UTC
	}

	var rv []time.Time
	for minute := int32(0); minute < 60; minute++ {
		v, err := instance.QueryStars(int32(pattern), seed, int32(year), int32(month), int32(date), hour.RegularHour(), minute)
		if err != nil {
			return nil, err
		}

		if v <= 0 {
			continue
		}

		v, err = instance.GetStarAmount()
		if err != nil {
			return nil, err
		}

		for idx := int32(0); idx < v; idx++ {
			sec, err := instance.GetStarSecond(idx)
			if err != nil {
				return nil, err
			}

			ts := time.Date(year, month, date, int(hour.RegularHour()), int(minute), int(sec), nsec, loc)
			rv = append(rv, ts)
		}
	}

	cp := make([]time.Time, len(rv))
	copy(cp, rv)
	return cp, nil
}
