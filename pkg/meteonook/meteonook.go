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
	"sort"
	"time"

	"github.com/demosdemon/weather/pkg/meteonook/enums"
)

func GetConstellation(date time.Time) enums.Constellation {
	_, month, day := date.Date()
	switch {
	case month == time.December && day > 21, month == time.January && day <= 19:
		return enums.Capricorn
	case month == time.January && day > 19, month == time.February && day <= 18:
		return enums.Aquarius
	case month == time.February && day > 18, month == time.March && day <= 20:
		return enums.Pisces
	case month == time.March && day > 20, month == time.April && day <= 19:
		return enums.Aries
	case month == time.April && day > 19, month == time.May && day <= 20:
		return enums.Taurus
	case month == time.May && day > 20, month == time.June && day <= 21:
		return enums.Gemini
	case month == time.June && day > 21, month == time.July && day <= 22:
		return enums.Cancer
	case month == time.July && day > 22, month == time.August && day <= 22:
		return enums.Leo
	case month == time.August && day > 22, month == time.September && day <= 22:
		return enums.Virgo
	case month == time.September && day > 22, month == time.October && day <= 23:
		return enums.Libra
	case month == time.October && day > 23, month == time.November && day <= 22:
		return enums.Scorpio
	case month == time.November && day > 22, month == time.December && day <= 21:
		return enums.Sagittarius
	}

	panic("impossible")
}

func GetSpecialDay(date time.Time, hemisphere enums.Hemisphere) enums.SpecialDay {
	year, month, day := date.Date()
	y := year - 2000

	if y < 0 || 61 < y {
		return 0
	}

	switch {
	// TODO: why is this 2020 only?
	case year == 2020 && month == easterMonths[y] && day == easterDays[y]:
		return enums.BunnyDay
	case
		month == time.January && day == fishConJan[y],
		month == time.April && day == fishConApr[y],
		month == time.July && day == fishConJul[y],
		month == time.October && day == fishConOct[y]:
		return enums.FishContest
	case
		hemisphere == enums.Northern && month == time.June && day == bugConJun[y],
		hemisphere == enums.Northern && month == time.July && day == bugConJul[y],
		hemisphere == enums.Northern && month == time.August && day == bugConAug[y],
		hemisphere == enums.Northern && month == time.September && day == bugConSep[y],
		hemisphere == enums.Southern && month == time.December && day == bugConDec[y],
		hemisphere == enums.Southern && month == time.January && day == bugConJan[y],
		hemisphere == enums.Southern && month == time.February && day == bugConFeb[y],
		hemisphere == enums.Southern && month == time.March && day == bugConMar[y]:
		return enums.InsectContest
	case month == time.December && day == 31:
		return enums.Countdown
	default:
		return enums.RegularDay
	}
}

func GetSnowLevel(date time.Time, hemisphere enums.Hemisphere) enums.SnowLevel {
	_, month, day := date.Date()
	switch {
	case
		hemisphere == enums.Northern && month == time.November && day > 25,
		hemisphere == enums.Northern && month == time.December && day < 11,
		hemisphere == enums.Southern && month == time.May && day > 25,
		hemisphere == enums.Southern && month == time.June && day < 11:
		return enums.LowSnow
	case
		hemisphere == enums.Northern && month == time.December && day >= 11,
		hemisphere == enums.Northern && month == time.January,
		hemisphere == enums.Northern && month == time.February && day < 25,
		hemisphere == enums.Southern && month == time.June && day >= 11,
		hemisphere == enums.Southern && month == time.July,
		hemisphere == enums.Southern && month == time.August && day < 25:
		return enums.FullSnow
	default:
		return enums.NoSnow
	}
}

func GetCloudLevel(date time.Time, hemisphere enums.Hemisphere) enums.CloudLevel {
	_, month, day := date.Date()
	switch {
	case
		hemisphere == enums.Northern && month == time.July && day > 20,
		hemisphere == enums.Northern && month == time.August,
		hemisphere == enums.Northern && month == time.September && day <= 15,
		hemisphere == enums.Southern && month == time.January && day > 20,
		hemisphere == enums.Southern && month == time.February,
		hemisphere == enums.Southern && month == time.March && day <= 15:
		return enums.Cumulonimbus
	case
		hemisphere == enums.Northern && month == time.September && day > 15,
		hemisphere == enums.Northern && month == time.October,
		hemisphere == enums.Northern && month == time.November,
		hemisphere == enums.Southern && month == time.March && day > 15,
		hemisphere == enums.Southern && month == time.April,
		hemisphere == enums.Southern && month == time.May:
		return enums.Cirrus
	case
		hemisphere == enums.Northern && month == time.December,
		hemisphere == enums.Northern && month == time.January,
		hemisphere == enums.Northern && month == time.February,
		hemisphere == enums.Southern && month == time.June,
		hemisphere == enums.Southern && month == time.July,
		hemisphere == enums.Southern && month == time.August:
		return enums.Billow
	case
		hemisphere == enums.Northern && month == time.March,
		hemisphere == enums.Northern && month == time.April,
		hemisphere == enums.Northern && month == time.May,
		hemisphere == enums.Southern && month == time.September,
		hemisphere == enums.Southern && month == time.October,
		hemisphere == enums.Southern && month == time.November:
		return enums.Thin
	default:
		return enums.NoClouds
	}
}

func GetSpecialWeatherLevel(date time.Time, hemisphere enums.Hemisphere) enums.SpecialWeatherLevel {
	_, month, day := date.Date()
	switch {
	case
		hemisphere == enums.Northern && month == time.December && day > 10,
		hemisphere == enums.Northern && month == time.January,
		hemisphere == enums.Northern && month == time.February && day < 25,
		hemisphere == enums.Southern && month == time.June && day > 10,
		hemisphere == enums.Southern && month == time.July,
		hemisphere == enums.Southern && month == time.August && day < 25:
		return enums.Aurora
	case
		hemisphere == enums.Northern && month == time.February && day >= 25,
		hemisphere == enums.Northern && month >= time.March && month <= time.October,
		hemisphere == enums.Northern && month == time.November && day <= 25,
		hemisphere == enums.Southern && month == time.August && day >= 25,
		hemisphere == enums.Southern && (month >= time.September || month <= time.April),
		hemisphere == enums.Southern && month == time.May && day <= 25:
		return enums.Rainbow
	default:
		return enums.NoSpecialWeather
	}
}

func IsAuroraPattern(date time.Time, hemisphere enums.Hemisphere, pattern enums.Pattern) bool {
	return GetSpecialWeatherLevel(date, hemisphere) == enums.Aurora &&
		(pattern == enums.Fine01 || pattern == enums.Fine03 || pattern == enums.Fine05)
}

func GetFogType(date time.Time, hemisphere enums.Hemisphere) enums.FogType {
	_, month, day := date.Date()
	switch {
	case
		hemisphere == enums.Northern && month == time.September && day > 20,
		hemisphere == enums.Northern && month >= time.October,
		hemisphere == enums.Northern && month == time.January,
		hemisphere == enums.Northern && month == time.February && day < 25,
		hemisphere == enums.Southern && month == time.March && day > 20,
		hemisphere == enums.Southern && month >= time.April && month <= time.July,
		hemisphere == enums.Northern && month == time.August && day < 25:
		return enums.HeavyFog
	case
		hemisphere == enums.Northern && month == time.February && day >= 25,
		hemisphere == enums.Northern && month == time.March,
		hemisphere == enums.Southern && month == time.August && day >= 25,
		hemisphere == enums.Southern && month == time.September:
		return enums.WaterFog
	default:
		return enums.NoFog
	}
}

func CheckWaterFog(date time.Time, seed uint32) bool {
	year, month, day := date.Date()
	rng := Random{
		uint32(year) << 8,
		uint32(month) << 8,
		uint32(day) << 8,
		seed | mask,
	}
	rng.Roll()
	rng.Roll()
	return (rng.Roll() & 1) == 1
}

func GetRainbowInfo(date time.Time, hemisphere enums.Hemisphere, pattern enums.Pattern, seed uint32) *RainbowInfo {
	var info RainbowInfo

	switch GetSpecialWeatherLevel(date, hemisphere) {
	case enums.Rainbow:
		switch pattern.Kind() {
		case enums.CloudFine, enums.FineRain:
			const (
				year  = 0x1000000
				month = 0x40000
				day   = 0x1000
			)
			rng := NewRandom(ComputeSeed(seed, year, month, day, date))
			rng.Roll()
			rng.Roll()
			info.Type = RainbowType(rng.Roll()&1) + 1
			for hour := 7; hour <= 17; hour++ {
				a := patterns[pattern][hour]
				b := patterns[pattern][hour+1]
				if a >= enums.LightStorm && b <= enums.Sunny {
					info.Hour = NewLinearHour(hour + 1)
					return &info
				}
			}
		}
	}

	return nil
}

func GetPattern(date time.Time, hemisphere enums.Hemisphere, seed uint32) enums.Pattern {
	if GetSpecialDay(date, hemisphere) > enums.RegularDay {
		return enums.EventDay00
	}

	const (
		year  = 0x2000000
		month = 0x200000
		day   = 0x10000
	)
	rng := NewRandom(ComputeSeed(seed, year, month, day, date))
	rng.Roll()
	rng.Roll()

	rateSet := getRateSet(hemisphere, date)
	rate := rng.RollMax(100)
	return rateSets[rateSet][rate]
}

func GetWeather(hour LinearHour, pattern enums.Pattern) enums.Weather {
	return patterns[pattern][hour.RegularHour()]
}

func GetWindPower(seed uint32, date time.Time, pattern enums.Pattern) uint8 {
	const (
		year  = 0x2000000
		month = 0x200000
		day   = 0x10000
	)
	rng := NewRandom(ComputeSeed(seed, year, month, day, date))
	rng.Roll()
	rng.Roll()
	switch winds[pattern][date.Hour()] {
	case enums.Land0, enums.Sea0:
		return rng.RollMax8(3)
	case enums.Land1, enums.Sea1:
		return rng.RollMax8(4) + 1
	case enums.Land2, enums.Sea2:
		return rng.RollMax8(3) + 3
	default:
		return 0
	}
}

func IsShootingStarsPossible(hour LinearHour, pattern enums.Pattern) bool {
	return (hour >= SevenPM && hour < FourAM) && pattern.ShowerLevel() > enums.NoShower
}

func GetShootingStars(seed uint32, date time.Time, pattern enums.Pattern) []time.Time {
	if !IsShootingStarsPossible(NewLinearHour(date.Hour()), pattern) {
		return nil
	}

	const (
		year  = 0x20000
		month = 0x2000
		day   = 0x100
		hour  = 0x10000
	)
	date = date.Truncate(time.Minute)
	seed = ComputeSeedHour(seed, year, month, day, hour, date)

	var rv []time.Time
	for idx := 0; idx < 60; idx++ {
		date := date.Add(time.Duration(idx) * time.Minute)
		rv = append(rv, getShootingStars(seed, date, pattern)...)
	}

	return rv
}

func getShootingStars(seed uint32, date time.Time, pattern enums.Pattern) []time.Time {
	rng := NewRandom(seed + uint32(date.Minute()*0x100))
	count := 0
	switch pattern {
	case enums.Fine00:
		if rng.RollMax(100) < 50 {
			if rng.RollMax(100) < 50 {
				count = 8
			} else {
				count = 5
			}
		}
	case enums.Fine02, enums.Fine04, enums.Fine06:
		var chance uint32 = 4
		if date.Minute()&1 == 0 {
			chance = 2
		}
		if rng.RollMax(60) < chance {
			count = 5
		}
	}

	if count == 0 {
		return nil
	}

	secs := make(map[uint32]bool, count)
	rv := make([]time.Time, 0, count)
	for len(rv) < cap(rv) {
		sec := rng.RollMax(60)
		if !secs[sec] {
			secs[sec] = true
			rv = append(rv, date.Add(time.Duration(sec)*time.Second))
		}
	}

	sort.Slice(rv, func(i, j int) bool { return rv[i].Before(rv[j]) })

	return rv
}

func getRateSet(hemisphere enums.Hemisphere, date time.Time) int {
	_, month, day := date.Date()
	switch hemisphere {
	case enums.Northern:
		return rateSetNorth[month-1][day-1]
	case enums.Southern:
		return rateSetSouth[month-1][day-1]
	}
	return 0
}
