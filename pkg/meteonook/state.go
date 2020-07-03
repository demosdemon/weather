// weather
// Copyright (C) 2020 Brandon LeBlanc <brandon@leblanc.codes>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package meteonook

import (
	"errors"
	"time"

	"github.com/demosdemon/weather/pkg/meteonook/enums"
)

const (
	minYear = 2000
	maxYear = 2060
)

// Pattern is EventDay00 when SpecialDay > RegularDay
// Fog is visible when FogType > NoFog && SixAM < LinearHour < TenAM
// Aurora is visible when AuroraPossible is true && Pattern == Fine00 && SixPM < LinearHour < ThreeAM
// Show is visible when SnowPossible is true && Weather > StormClouds

type Island struct {
	Name       string           `json:"name,omitempty"`
	Hemisphere enums.Hemisphere `json:"hemisphere"`
	Seed       int32            `json:"seed"`
	Timezone   Timezone         `json:"timezone"`
}

type Day struct {
	Island         *Island           `json:"island"`
	Year           int32             `json:"year"`
	Month          time.Month        `json:"month"`
	Date           int32             `json:"date"`
	Weekday        time.Weekday      `json:"weekday"`
	SpecialDay     enums.SpecialDay  `json:"special_day,omitempty"`
	Pattern        enums.Pattern     `json:"pattern"`
	ShowerLevel    enums.ShowerLevel `json:"shower_type,omitempty"`
	FogType        enums.FogType     `json:"fog_type,omitempty"`
	RainbowType    RainbowType       `json:"rainbow_type,omitempty"`
	AuroraPossible bool              `json:"aurora_possible,omitempty"`
	SnowPossible   bool              `json:"snow_possible,omitempty"`
	Hours          [24]Hour          `json:"hours"`
}

type Hour struct {
	Hour          LinearHour        `json:"hour"`
	Weather       enums.Weather     `json:"weather"`
	WindPower     int32             `json:"wind_power,omitempty"`
	ShowerType    enums.ShowerLevel `json:"shower_type,omitempty"`
	FogType       enums.FogType     `json:"fog_type,omitempty"`
	RainbowType   RainbowType       `json:"rainbow_type,omitempty"`
	AuroraVisible bool              `json:"aurora_visible,omitempty"`
	SnowVisible   bool              `json:"snow_visible,omitempty"`
	ShootingStars []time.Time       `json:"shooting_stars,omitempty"`
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
	if day.Pattern, err = getWeatherPattern(instance, island.Seed, island.Hemisphere, year, month, date); err != nil {
		return nil, err
	}
	day.ShowerLevel = day.Pattern.ShowerLevel()
	if day.FogType, err = getFogType(instance, island.Seed, island.Hemisphere, year, month, date, day.Pattern, yesterday); err != nil {
		return nil, err
	}
	var rainbowHour LinearHour
	if day.RainbowType, rainbowHour, err = getRainbowType(instance, island.Seed, island.Hemisphere, year, month, date, day.Pattern); err != nil {
		return nil, err
	}
	if day.AuroraPossible, err = getAuroraPossible(instance, island.Seed, island.Hemisphere, year, month, date, day.Pattern); err != nil {
		return nil, err
	}
	if day.SnowPossible, err = getSnowPossible(instance, island.Hemisphere, month, date); err != nil {
		return nil, err
	}

	for idx := range day.Hours {
		hour := LinearHour(idx)
		weather, err := getWeather(instance, day.Pattern, hour)
		if err != nil {
			return nil, err
		}
		windPower, err := getWindPower(instance, island.Seed, island.Hemisphere, year, month, date, hour, day.Pattern)
		if err != nil {
			return nil, err
		}
		shootingStarsPossible, err := getShootingStarsPossible(instance, day.Pattern, hour)
		showerType := day.ShowerLevel
		if !shootingStarsPossible {
			showerType = enums.NoShower
		}
		fogType := enums.NoFog
		if SixAM < hour && hour < TenAM {
			fogType = day.FogType
		}
		rainbowType := NoRainbow
		if rainbowHour == hour || rainbowHour == hour+1 {
			rainbowType = day.RainbowType
		}
		aurora := day.AuroraPossible && day.Pattern == enums.Fine00 && SixPM < hour && hour < ThreeAM
		snow := day.SnowPossible && weather > enums.StormClouds
		var shootingStars []time.Time
		if showerType > enums.NoShower {
			if shootingStars, err = getShootingStars(instance, island.Seed, year, month, date, hour, island.Timezone, day.Pattern); err != nil {
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
	hemisphere enums.Hemisphere,
	year int,
	month time.Month,
	date int,
) (enums.SpecialDay, error) {
	v, err := instance.IsSpecialDay(int32(hemisphere), int32(year), int32(month), int32(date))
	return enums.SpecialDay(v), err
}

func getWeatherPattern(
	instance *Instance,
	seed int32,
	hemisphere enums.Hemisphere,
	year int,
	month time.Month,
	date int,
) (enums.Pattern, error) {
	v, err := instance.GetPattern(seed, int32(hemisphere), int32(year), int32(month), int32(date))
	return enums.Pattern(v), err
}

func getFogType(
	instance *Instance,
	seed int32,
	hemisphere enums.Hemisphere,
	year int,
	month time.Month,
	date int,
	today enums.Pattern,
	yesterday enums.Pattern,
) (enums.FogType, error) {
	v, err := instance.GetFog(int32(hemisphere), int32(month), int32(date))

	if err != nil {
		return 0, err
	}

	if v == 0 {
		return enums.NoFog, nil
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
		return enums.HeavyFog, nil
	}

	waterFog := preWaterFogPatterns[yesterday]
	waterFog = waterFog && fogPatterns[today]
	if waterFog {
		v, err = instance.CheckWaterFog(seed, int32(year), int32(month), int32(date))
		if err != nil {
			return 0, err
		}
		if v > 0 {
			return enums.WaterFog, nil
		}
	}

	return enums.NoFog, nil
}

func getRainbowType(
	instance *Instance,
	seed int32,
	hemisphere enums.Hemisphere,
	year int,
	month time.Month,
	date int,
	pattern enums.Pattern,
) (RainbowType, LinearHour, error) {
	v, err := instance.IsRainbowPattern(int32(hemisphere), seed, int32(year), int32(month), int32(date), int32(pattern))
	if err != nil {
		return 0, 0, err
	}

	return RainbowType(v >> 8), NewLinearHour(v&0xFF) + 1, nil
}

func getAuroraPossible(
	instance *Instance,
	seed int32,
	hemisphere enums.Hemisphere,
	year int,
	month time.Month,
	date int,
	pattern enums.Pattern,
) (bool, error) {
	v, err := instance.IsAuroraPattern(int32(hemisphere), seed, int32(year), int32(month), int32(date), int32(pattern))
	return v > 0, err
}

func getSnowPossible(
	instance *Instance,
	hemisphere enums.Hemisphere,
	month time.Month,
	date int,
) (bool, error) {
	v, err := instance.GetSnow(int32(hemisphere), int32(month), int32(date))
	return v > 0, err
}

func getWeather(
	instance *Instance,
	pattern enums.Pattern,
	hour LinearHour,
) (enums.Weather, error) {
	v, err := instance.GetWeather(int32(pattern), hour.RegularHour())
	return enums.Weather(v), err
}

func getWindPower(
	instance *Instance,
	seed int32,
	hemisphere enums.Hemisphere,
	year int,
	month time.Month,
	date int,
	hour LinearHour,
	pattern enums.Pattern,
) (int32, error) {
	return instance.GetWindPower(int32(hemisphere), seed, int32(year), int32(month), int32(date), hour.RegularHour(), int32(pattern))
}

func getShootingStarsPossible(
	instance *Instance,
	pattern enums.Pattern,
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
	pattern enums.Pattern,
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
