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
	"log"
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
	Seed       uint32           `json:"seed"`
	Timezone   Timezone         `json:"timezone"`
}

type Day struct {
	Island         *Island             `json:"island,omitempty"`
	Year           int32               `json:"year"`
	Month          time.Month          `json:"month"`
	Date           int32               `json:"date"`
	Weekday        time.Weekday        `json:"weekday"`
	Constellation  enums.Constellation `json:"constellation"`
	SpecialDay     enums.SpecialDay    `json:"special_day,omitempty"`
	CloudLevel     enums.CloudLevel    `json:"cloud_level,omitempty"`
	Pattern        enums.Pattern       `json:"pattern"`
	ShowerLevel    enums.ShowerLevel   `json:"shower_type,omitempty"`
	FogType        enums.FogType       `json:"fog_type,omitempty"`
	RainbowInfo    *RainbowInfo        `json:"rainbow_info,omitempty"`
	AuroraPossible bool                `json:"aurora_possible,omitempty"`
	SnowLevel      enums.SnowLevel     `json:"snow_level,omitempty"`
	Hours          [24]Hour            `json:"hours"`
}

type Hour struct {
	Hour          LinearHour        `json:"hour"`
	Weather       enums.Weather     `json:"weather"`
	WindPower     uint8             `json:"wind_power,omitempty"`
	ShowerLevel   enums.ShowerLevel `json:"shower_level,omitempty"`
	FogType       enums.FogType     `json:"fog_type,omitempty"`
	RainbowType   RainbowType       `json:"rainbow_type,omitempty"`
	AuroraVisible bool              `json:"aurora_visible,omitempty"`
	SnowLevel     enums.SnowLevel   `json:"snow_level,omitempty"`
	ShootingStars []time.Time       `json:"shooting_stars,omitempty"`
}

type RainbowInfo struct {
	Hour LinearHour  `json:"hour"`
	Type RainbowType `json:"type,omitempty"`
}

func (island *Island) NewDay(ts time.Time) (*Day, error) {
	ts = ts.Add(5 * time.Hour)
	log.Printf("ts=%v;", ts)
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

	yts := ts.AddDate(0, 0, -1)

	yesterday := GetPattern(yts, island.Hemisphere, island.Seed)
	day.Constellation = GetConstellation(ts)
	day.SpecialDay = GetSpecialDay(ts, island.Hemisphere)
	day.CloudLevel = GetCloudLevel(ts, island.Hemisphere)
	day.Pattern = GetPattern(ts, island.Hemisphere, island.Seed)
	day.ShowerLevel = day.Pattern.ShowerLevel()
	day.FogType = getFogType(island.Hemisphere, island.Seed, ts, yesterday, day.Pattern)
	day.RainbowInfo = GetRainbowInfo(ts, island.Hemisphere, day.Pattern, island.Seed)
	day.AuroraPossible = IsAuroraPattern(ts, island.Hemisphere, day.Pattern)
	day.SnowLevel = GetSnowLevel(ts, island.Hemisphere)

	for idx := range day.Hours {
		ts := ts.Add(time.Duration(idx) * time.Hour)
		hour := LinearHour(idx)
		weather := GetWeather(hour, day.Pattern)

		showerLevel := day.ShowerLevel
		shootingStarsPossible := IsShootingStarsPossible(hour, day.Pattern)
		if !shootingStarsPossible {
			showerLevel = enums.NoShower
		}

		fogType := enums.NoFog
		if SixAM < hour && hour < TenAM {
			fogType = day.FogType
		}

		rainbowType := NoRainbow
		if day.RainbowInfo != nil && (day.RainbowInfo.Hour == hour || day.RainbowInfo.Hour == hour+1) {
			rainbowType = day.RainbowInfo.Type
		}

		snow := day.SnowLevel
		if weather <= enums.StormClouds {
			snow = enums.NoSnow
		}

		day.Hours[idx] = Hour{
			Hour:          hour,
			Weather:       weather,
			WindPower:     GetWindPower(island.Seed, ts, day.Pattern),
			ShowerLevel:   showerLevel,
			FogType:       fogType,
			RainbowType:   rainbowType,
			AuroraVisible: day.AuroraPossible && day.Pattern == enums.Fine00 && SixPM < hour && hour < ThreeAM,
			SnowLevel:     snow,
			ShootingStars: GetShootingStars(island.Seed, ts, day.Pattern),
		}
	}

	return day, nil
}

func getFogType(
	hemisphere enums.Hemisphere,
	seed uint32,
	date time.Time,
	today enums.Pattern,
	yesterday enums.Pattern,
) enums.FogType {
	getWindPower := func(hour int) bool {
		date := date.Add(time.Duration(hour) * time.Hour)
		v := GetWindPower(seed, date, yesterday)
		return 3 > v
	}

	switch GetFogType(date, hemisphere) {
	case enums.HeavyFog:
		if preNormalFogPatterns[yesterday] &&
			fogPatterns[today] &&
			getWindPower(5) &&
			getWindPower(6) &&
			getWindPower(7) &&
			getWindPower(8) {
			return enums.HeavyFog
		}
		fallthrough
	case enums.WaterFog:
		if preWaterFogPatterns[yesterday] && fogPatterns[today] && CheckWaterFog(date, seed) {
			return enums.WaterFog
		}
	}

	return enums.NoFog
}
