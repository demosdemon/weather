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
	"time"

	"github.com/demosdemon/weather/pkg/meteonook/enums"
)

var (
	easterDays   = [61]int{23, 15, 31, 20, 11, 27, 16, 8, 23, 12, 4, 24, 8, 31, 20, 5, 27, 16, 1, 21, 12, 4, 17, 9, 31, 20, 5, 28, 16, 1, 21, 13, 28, 17, 9, 25, 13, 5, 25, 10, 1, 21, 6, 29, 17, 9, 25, 14, 5, 18, 10, 2, 21, 6, 29, 18, 2, 22, 14, 30, 18}
	easterMonths = [61]time.Month{4, 4, 3, 4, 4, 3, 4, 4, 3, 4, 4, 4, 4, 3, 4, 4, 3, 4, 4, 4, 4, 4, 4, 4, 3, 4, 4, 3, 4, 4, 4, 4, 3, 4, 4, 3, 4, 4, 4, 4, 4, 4, 4, 3, 4, 4, 3, 4, 4, 4, 4, 4, 4, 4, 3, 4, 4, 4, 4, 3, 4}

	fishConJan = [61]int{8, 13, 12, 11, 10, 8, 14, 13, 12, 10, 9, 8, 14, 12, 11, 10, 9, 14, 13, 12, 11, 9, 8, 14, 13, 11, 10, 9, 8, 13, 12, 11, 10, 8, 14, 13, 12, 10, 9, 8, 14, 12, 11, 10, 9, 14, 13, 12, 11, 9, 8, 14, 13, 11, 10, 9, 8, 13, 12, 11, 10}
	fishConApr = [61]int{8, 14, 13, 12, 10, 9, 8, 14, 12, 11, 10, 9, 14, 13, 12, 11, 9, 8, 14, 13, 11, 10, 9, 8, 13, 12, 11, 10, 8, 14, 13, 12, 10, 9, 8, 14, 12, 11, 10, 9, 14, 13, 12, 11, 9, 8, 14, 13, 11, 10, 9, 8, 13, 12, 11, 10, 8, 14, 13, 12, 10}
	fishConJul = [61]int{8, 14, 13, 12, 10, 9, 8, 14, 12, 11, 10, 9, 14, 13, 12, 11, 9, 8, 14, 13, 11, 10, 9, 8, 13, 12, 11, 10, 8, 14, 13, 12, 10, 9, 8, 14, 12, 11, 10, 9, 14, 13, 12, 11, 9, 8, 14, 13, 11, 10, 9, 8, 13, 12, 11, 10, 8, 14, 13, 12, 10}
	fishConOct = [61]int{14, 13, 12, 11, 9, 8, 14, 13, 11, 10, 9, 8, 13, 12, 11, 10, 8, 14, 13, 12, 10, 9, 8, 14, 12, 11, 10, 9, 14, 13, 12, 11, 9, 8, 14, 13, 11, 10, 9, 8, 13, 12, 11, 10, 8, 14, 13, 12, 10, 9, 8, 14, 12, 11, 10, 9, 14, 13, 12, 11, 9}

	bugConJun = [61]int{24, 23, 22, 28, 26, 25, 24, 23, 28, 27, 26, 25, 23, 22, 28, 27, 25, 24, 23, 22, 27, 26, 25, 24, 22, 28, 27, 26, 24, 23, 22, 28, 26, 25, 24, 23, 28, 27, 26, 25, 23, 22, 28, 27, 25, 24, 23, 22, 27, 26, 25, 24, 22, 28, 27, 26, 24, 23, 22, 28, 26}
	bugConJul = [61]int{22, 28, 27, 26, 24, 23, 22, 28, 26, 25, 24, 23, 28, 27, 26, 25, 23, 22, 28, 27, 25, 24, 23, 22, 27, 26, 25, 24, 22, 28, 27, 26, 24, 23, 22, 28, 26, 25, 24, 23, 28, 27, 26, 25, 23, 22, 28, 27, 25, 24, 23, 22, 27, 26, 25, 24, 22, 28, 27, 26, 24}
	bugConAug = [61]int{26, 25, 24, 23, 28, 27, 26, 25, 23, 22, 28, 27, 25, 24, 23, 22, 27, 26, 25, 24, 22, 28, 27, 26, 24, 23, 22, 28, 26, 25, 24, 23, 28, 27, 26, 25, 23, 22, 28, 27, 25, 24, 23, 22, 27, 26, 25, 24, 22, 28, 27, 26, 24, 23, 22, 28, 26, 25, 24, 23, 28}
	bugConSep = [61]int{23, 22, 28, 27, 25, 24, 23, 22, 27, 26, 25, 24, 22, 28, 27, 26, 24, 23, 22, 28, 26, 25, 24, 23, 28, 27, 26, 25, 23, 22, 28, 27, 25, 24, 23, 22, 27, 26, 25, 24, 22, 28, 27, 26, 24, 23, 22, 28, 26, 25, 24, 23, 28, 27, 26, 25, 23, 22, 28, 27, 25}

	bugConNov = [61]int{18, 17, 16, 15, 20, 19, 18, 17, 15, 21, 20, 19, 17, 16, 15, 21, 19, 18, 17, 16, 21, 20, 19, 18, 16, 15, 21, 20, 18, 17, 16, 15, 20, 19, 18, 17, 15, 21, 20, 19, 17, 16, 15, 21, 19, 18, 17, 16, 21, 20, 19, 18, 16, 15, 21, 20, 18, 17, 16, 15, 20}
	bugConDec = [61]int{16, 15, 21, 20, 18, 17, 16, 15, 20, 19, 18, 17, 15, 21, 20, 19, 17, 16, 15, 21, 19, 18, 17, 16, 21, 20, 19, 18, 16, 15, 21, 20, 18, 17, 16, 15, 20, 19, 18, 17, 15, 21, 20, 19, 17, 16, 15, 21, 19, 18, 17, 16, 21, 20, 19, 18, 16, 15, 21, 20, 18}
	bugConJan = [61]int{15, 20, 19, 18, 17, 15, 21, 20, 19, 17, 16, 15, 21, 19, 18, 17, 16, 21, 20, 19, 18, 16, 15, 21, 20, 18, 17, 16, 15, 20, 19, 18, 17, 15, 21, 20, 19, 17, 16, 15, 21, 19, 18, 17, 16, 21, 20, 19, 18, 16, 15, 21, 20, 18, 17, 16, 15, 20, 19, 18, 17}
	bugConFeb = [61]int{19, 17, 16, 15, 21, 19, 18, 17, 16, 21, 20, 19, 18, 16, 15, 21, 20, 18, 17, 16, 15, 20, 19, 18, 17, 15, 21, 20, 19, 17, 16, 15, 21, 19, 18, 17, 16, 21, 20, 19, 18, 16, 15, 21, 20, 18, 17, 16, 15, 20, 19, 18, 17, 15, 21, 20, 19, 17, 16, 15, 21}

	rateSets = [40][100]enums.Pattern{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 8, 9, 10, 10, 10, 11, 11, 11, 12, 12, 12, 13, 13, 13, 14, 14, 14, 15, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10, 10, 11, 11, 12, 12, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 26, 26, 26, 27, 27, 27, 28, 28, 28, 29, 29, 30, 30, 31, 31},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 10, 10, 10, 11, 11, 11, 12, 12, 12, 13, 13, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31},
		{0, 0, 2, 2, 2, 2, 2, 2, 2, 2, 4, 4, 4, 4, 4, 4, 4, 4, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 13, 13, 13, 14, 14, 14, 15, 15, 15, 16, 17, 18, 19, 20, 21, 26, 26, 26, 27, 27, 27, 28, 28, 28, 29, 29, 29, 30, 30, 30, 31, 31, 31},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 8, 9, 9, 10, 11, 12, 13, 14, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 22, 22, 22, 23, 23, 23, 23, 23, 24, 24, 24, 25, 25, 25},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 8, 9, 19, 19, 20, 20, 21, 21, 22, 22, 22, 23, 23, 23, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 10, 10, 11, 11, 12, 12, 13, 13, 13, 14, 14, 14, 15, 15, 15, 22, 22, 22, 22, 22, 22, 22, 22, 23, 23, 23, 23, 23, 23, 23, 23, 24, 24, 24, 24, 24, 24, 25, 25, 25, 25, 25, 25, 26, 27, 28, 29, 30, 31},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 8, 8, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 17, 18, 19, 19, 20, 20, 21, 21, 22, 22, 22, 22, 23, 23, 23, 23, 26, 27, 28, 29, 30, 31},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 14, 15, 16, 16, 17, 17, 17, 17, 18, 18, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 22, 23, 23, 23},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10, 10, 10, 11, 11, 11, 12, 12, 12, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31},
		{13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{0, 0, 0, 0, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 10, 10, 10, 11, 11, 11, 11, 11, 12, 12, 12, 12, 12, 13, 13, 13, 14, 14, 14, 15, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 26, 26, 26, 27, 27, 27, 28, 28, 28, 29, 29, 29, 30, 30, 30, 31, 31, 31},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6},

		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 8, 9, 9, 10, 11, 12, 13, 14, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 22, 22, 22, 23, 23, 23, 23, 23, 24, 24, 24, 25, 25, 25},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 19, 20, 20, 21, 21, 22, 22, 22, 22, 23, 23, 23, 23, 24, 24, 24, 24, 24, 25, 25, 25, 25, 25},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 13, 14, 14, 14, 15, 15, 15, 16, 17, 18, 19, 20, 21, 22, 22, 22, 22, 23, 23, 23, 23, 24, 24, 24, 24, 25, 25, 25, 25, 26, 27, 28, 29, 30, 31},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 8, 8, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 17, 18, 19, 19, 20, 20, 21, 21, 22, 22, 22, 22, 23, 23, 23, 23, 26, 27, 28, 29, 30, 31},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10, 10, 11, 11, 12, 12, 13, 14, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 22, 23, 23, 23, 26, 27, 29, 30},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 8, 8, 8, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10, 10, 10, 11, 11, 11, 12, 12, 12, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31},
		{13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{0, 0, 0, 0, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 10, 10, 10, 11, 11, 11, 11, 11, 12, 12, 12, 12, 12, 13, 13, 13, 14, 14, 14, 15, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 26, 26, 26, 27, 27, 27, 28, 28, 28, 29, 29, 29, 30, 30, 30, 31, 31, 31},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 8, 9, 10, 10, 10, 11, 11, 11, 12, 12, 12, 13, 13, 13, 14, 14, 14, 15, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10, 10, 11, 11, 12, 12, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 26, 26, 26, 27, 27, 27, 28, 28, 28, 29, 29, 30, 30, 31, 31},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 10, 10, 10, 11, 11, 11, 12, 12, 12, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 10, 10, 10, 11, 11, 11, 12, 12, 12, 13, 13, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 10, 10, 10, 11, 11, 11, 12, 12, 12, 13, 13, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31},
		{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6},
	}
	rateSetNorth = [12][31]int{
		{0, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 4, 4, 4, 4, 4, 4, 4},
		{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 12},
		{12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 13, 13, 13, 13, 13, 13, 13, 13, 13, 14, 15, 15, 15, 15, 15, 15},
		{15, 15, 15, 15, 15, 15, 15, 15, 15, 16, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 18, 18, 18, 18, 18, 18, 18, 19},
	}
	rateSetSouth = [12][31]int{
		{20, 21, 21, 21, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 24, 24},
		{24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25},
		{26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
		{26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 27, 27, 27, 27, 27, 27, 27, 27, 27, 28, 29, 29, 29, 29, 29, 29},
		{29, 29, 29, 29, 29, 29, 29, 29, 29, 30, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 32},
		{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
		{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 33, 34, 34, 34, 34, 34, 34, 34},
		{34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 35},
		{35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
		{36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
		{36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 37, 37, 37, 37, 37, 37, 37, 37, 38, 38, 38, 38, 38, 38, 38, 39},
	}

	patterns = map[enums.Pattern][24]enums.Weather{
		enums.Fine00:      {enums.Clear, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Clear, enums.Clear, enums.Sunny, enums.Clear},
		enums.Fine01:      {enums.Clear, enums.Clear, enums.Sunny, enums.Clear, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Cloudy, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Clear, enums.Sunny},
		enums.Fine02:      {enums.Clear, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Cloudy, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Clear},
		enums.Fine03:      {enums.Clear, enums.Sunny, enums.Clear, enums.Clear, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Cloudy, enums.Clear, enums.Clear, enums.Sunny},
		enums.Fine04:      {enums.Clear, enums.Clear, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear},
		enums.Fine05:      {enums.Cloudy, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny},
		enums.Fine06:      {enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Clear},
		enums.Cloud00:     {enums.Sunny, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Cloudy, enums.Sunny},
		enums.Cloud01:     {enums.Sunny, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.StormClouds, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Sunny},
		enums.Cloud02:     {enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Cloudy},
		enums.Rain00:      {enums.Sunny, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.StormClouds, enums.Cloudy, enums.Sunny, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy},
		enums.Rain01:      {enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.HeavyStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Sunny, enums.Sunny, enums.Cloudy},
		enums.Rain02:      {enums.StormClouds, enums.Cloudy, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.StormClouds, enums.LightStorm, enums.Sunny, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.StormClouds, enums.HeavyStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm},
		enums.Rain03:      {enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm, enums.LightStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.StormClouds},
		enums.Rain04:      {enums.LightStorm, enums.LightStorm, enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm, enums.LightStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm},
		enums.Rain05:      {enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.HeavyStorm, enums.LightStorm, enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm, enums.LightStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Sunny},
		enums.FineCloud00: {enums.StormClouds, enums.Cloudy, enums.StormClouds, enums.StormClouds, enums.Cloudy, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy},
		enums.FineCloud01: {enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Cloudy},
		enums.FineCloud02: {enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.StormClouds, enums.StormClouds, enums.LightStorm},
		enums.CloudFine00: {enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.LightStorm, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Clear},
		enums.CloudFine01: {enums.Sunny, enums.Clear, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny},
		enums.CloudFine02: {enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Clear, enums.Sunny, enums.Clear},
		enums.FineRain00:  {enums.Clear, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.LightStorm, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny},
		enums.FineRain01:  {enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.LightStorm, enums.Cloudy, enums.LightStorm, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny},
		enums.FineRain02:  {enums.Sunny, enums.Clear, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.HeavyStorm, enums.Cloudy, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear},
		enums.FineRain03:  {enums.Clear, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Clear, enums.Sunny, enums.Sunny, enums.HeavyStorm, enums.LightStorm, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny},
		enums.CloudRain00: {enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm},
		enums.CloudRain01: {enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.LightStorm, enums.LightStorm},
		enums.CloudRain02: {enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm, enums.HeavyStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.StormClouds},
		enums.RainCloud00: {enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy},
		enums.RainCloud01: {enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Sunny, enums.Sunny},
		enums.RainCloud02: {enums.Sunny, enums.Sunny, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.StormClouds, enums.HeavyStorm, enums.HeavyStorm, enums.LightStorm, enums.LightStorm, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.LightStorm, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.StormClouds, enums.Cloudy, enums.Cloudy},
		enums.Commun00:    {enums.Clear, enums.Clear, enums.Sunny, enums.Clear, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Cloudy, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Cloudy, enums.Sunny, enums.Clear, enums.Sunny},
		enums.EventDay00:  {enums.Clear, enums.Clear, enums.Sunny, enums.Clear, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny, enums.Clear, enums.Sunny},
	}

	winds = map[enums.Pattern][24]enums.WindType{
		enums.Fine00:      {enums.Land0, enums.Land1, enums.Land1, enums.Land2, enums.Land2, enums.Land0, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land1, enums.Land0},
		enums.Fine01:      {enums.Land1, enums.Land0, enums.Land1, enums.Land1, enums.Land1, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land0},
		enums.Fine02:      {enums.Land1, enums.Land0, enums.Land1, enums.Land2, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea1, enums.Sea0, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land2},
		enums.Fine03:      {enums.Land1, enums.Land2, enums.Land1, enums.Land1, enums.Land2, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land1, enums.Land0},
		enums.Fine04:      {enums.Land2, enums.Land1, enums.Land0, enums.Land1, enums.Land2, enums.Land1, enums.Land0, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land1, enums.Land1},
		enums.Fine05:      {enums.Land0, enums.Land1, enums.Land2, enums.Land1, enums.Land1, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land0, enums.Land1},
		enums.Fine06:      {enums.Land0, enums.Land1, enums.Land1, enums.Land2, enums.Land2, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land1, enums.Land0},
		enums.Cloud00:     {enums.Land1, enums.Land0, enums.Land1, enums.Land2, enums.Land2, enums.Land1, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land0, enums.Land0},
		enums.Cloud01:     {enums.Land1, enums.Land2, enums.Land2, enums.Land1, enums.Land2, enums.Land0, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land1},
		enums.Cloud02:     {enums.Land1, enums.Land1, enums.Land2, enums.Land2, enums.Land1, enums.Land1, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea1, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land0, enums.Land0},
		enums.Rain00:      {enums.Land2, enums.Land1, enums.Land0, enums.Land1, enums.Land2, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land1},
		enums.Rain01:      {enums.Land0, enums.Land1, enums.Land1, enums.Land2, enums.Land2, enums.Land1, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land1, enums.Land1, enums.Land0},
		enums.Rain02:      {enums.Land2, enums.Land1, enums.Land0, enums.Land1, enums.Land2, enums.Land2, enums.Land0, enums.Calm, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea2, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land1, enums.Land1, enums.Land0, enums.Land1},
		enums.Rain03:      {enums.Land1, enums.Land2, enums.Land1, enums.Land0, enums.Land1, enums.Land1, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea1, enums.Calm, enums.Land0, enums.Land1, enums.Land1, enums.Land0},
		enums.Rain04:      {enums.Land1, enums.Land0, enums.Land0, enums.Land1, enums.Land1, enums.Land2, enums.Land1, enums.Calm, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land1, enums.Land1, enums.Land2, enums.Land1, enums.Land2},
		enums.Rain05:      {enums.Land0, enums.Land0, enums.Land1, enums.Land2, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land1, enums.Land2, enums.Land1},
		enums.FineCloud00: {enums.Land0, enums.Land2, enums.Land1, enums.Land1, enums.Land2, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land1, enums.Land1},
		enums.FineCloud01: {enums.Land2, enums.Land1, enums.Land2, enums.Land1, enums.Land1, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land1, enums.Land1, enums.Land0},
		enums.FineCloud02: {enums.Land1, enums.Land0, enums.Land2, enums.Land1, enums.Land2, enums.Land0, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land1, enums.Land2},
		enums.CloudFine00: {enums.Land1, enums.Land2, enums.Land2, enums.Land1, enums.Land2, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land1},
		enums.CloudFine01: {enums.Land0, enums.Land0, enums.Land1, enums.Land1, enums.Land2, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land1, enums.Land0},
		enums.CloudFine02: {enums.Land1, enums.Land0, enums.Land1, enums.Land2, enums.Land1, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land1, enums.Land2},
		enums.FineRain00:  {enums.Land2, enums.Land1, enums.Land0, enums.Land1, enums.Land2, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land1},
		enums.FineRain01:  {enums.Land0, enums.Land1, enums.Land2, enums.Land0, enums.Land2, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea2, enums.Sea1, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land0, enums.Land0},
		enums.FineRain02:  {enums.Land1, enums.Land2, enums.Land1, enums.Land1, enums.Land2, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land1},
		enums.FineRain03:  {enums.Land0, enums.Land1, enums.Land2, enums.Land2, enums.Land2, enums.Land1, enums.Land0, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land0, enums.Land0},
		enums.CloudRain00: {enums.Land1, enums.Land0, enums.Land0, enums.Land1, enums.Land2, enums.Land1, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land1, enums.Land2},
		enums.CloudRain01: {enums.Land1, enums.Land2, enums.Land1, enums.Land2, enums.Land1, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land0, enums.Land1},
		enums.CloudRain02: {enums.Land2, enums.Land1, enums.Land0, enums.Land1, enums.Land2, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land1, enums.Land2},
		enums.RainCloud00: {enums.Land2, enums.Land1, enums.Land0, enums.Land1, enums.Land2, enums.Land1, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land1, enums.Land0, enums.Land1},
		enums.RainCloud01: {enums.Land1, enums.Land2, enums.Land1, enums.Land2, enums.Land2, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land1, enums.Land1},
		enums.RainCloud02: {enums.Land1, enums.Land0, enums.Land1, enums.Land0, enums.Land2, enums.Sea0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea2, enums.Sea1, enums.Sea1, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land1, enums.Land2},
		enums.Commun00:    {enums.Land1, enums.Land0, enums.Land1, enums.Land1, enums.Land1, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land0},
		enums.EventDay00:  {enums.Land1, enums.Land0, enums.Land1, enums.Land1, enums.Land1, enums.Land1, enums.Land0, enums.Calm, enums.Sea0, enums.Sea0, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea0, enums.Sea1, enums.Sea2, enums.Sea1, enums.Sea0, enums.Sea0, enums.Calm, enums.Land0, enums.Land0, enums.Land0, enums.Land0},
	}

	preNormalFogPatterns = map[enums.Pattern]bool{
		enums.Fine00:      true,
		enums.Fine01:      true,
		enums.Fine02:      true,
		enums.Fine03:      true,
		enums.Fine04:      true,
		enums.Fine05:      true,
		enums.Fine06:      true,
		enums.FineCloud00: true,
		enums.FineCloud01: true,
		enums.FineCloud02: true,
		enums.CloudFine00: true,
		enums.CloudFine01: true,
		enums.CloudFine02: true,
		enums.FineRain00:  true,
		enums.FineRain01:  true,
		enums.FineRain02:  true,
		enums.FineRain03:  true,
		enums.EventDay00:  true,
	}
	preWaterFogPatterns = map[enums.Pattern]bool{
		enums.Fine00:      true,
		enums.Fine01:      true,
		enums.Fine02:      true,
		enums.Fine03:      true,
		enums.Fine04:      true,
		enums.Fine05:      true,
		enums.Fine06:      true,
		enums.FineCloud00: true,
		enums.FineCloud01: true,
		enums.FineCloud02: true,
		enums.CloudFine00: true,
		enums.CloudFine01: true,
		enums.CloudFine02: true,
		enums.FineRain00:  true,
		enums.FineRain01:  true,
		enums.FineRain02:  true,
		enums.FineRain03:  true,
	}
	fogPatterns = map[enums.Pattern]bool{
		enums.Cloud00:     true,
		enums.Cloud01:     true,
		enums.Cloud02:     true,
		enums.Rain00:      true,
		enums.Rain01:      true,
		enums.Rain02:      true,
		enums.Rain03:      true,
		enums.Rain04:      true,
		enums.Rain05:      true,
		enums.FineCloud00: true,
		enums.FineCloud01: true,
		enums.FineCloud02: true,
		enums.CloudFine00: true,
		enums.CloudFine01: true,
		enums.CloudFine02: true,
		enums.CloudRain00: true,
		enums.CloudRain01: true,
		enums.CloudRain02: true,
		enums.RainCloud00: true,
		enums.RainCloud01: true,
		enums.RainCloud02: true,
	}
)
