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
)

type Random [4]uint32

func NewRandom(seed uint32) Random {
	var r Random
	r.Init(seed)
	return r
}

func (r *Random) Init(seed uint32) {
	const multi = 0x6c078965
	r[0] = (seed^(seed>>30))*multi + 1
	r[1] = (r[0]^(r[0]>>30))*multi + 2
	r[2] = (r[1]^(r[1]>>30))*multi + 3
	r[3] = (r[2]^(r[2]>>30))*multi + 4
}

func (r *Random) Roll() uint32 {
	n := r[0] ^ (r[0] << 11)
	r[0] = r[1]
	r[1] = r[2]
	r[2] = r[3]
	r[3] = n ^ (n >> 8) ^ r[3] ^ (r[3] >> 19)
	return r[3]
}

func (r *Random) RollMax(limit uint32) uint32 {
	value := uint64(r.Roll()) * uint64(limit)
	return uint32(value >> 32)
}

func (r *Random) RollMax8(limit uint8) uint8 {
	value := uint64(r.Roll()) * uint64(limit)
	return uint8(value >> 32)
}

const mask = 0x80000000

func ComputeSeed(base, yearMulti, monthMulti, dayMulti uint32, date time.Time) uint32 {
	year, month, day := date.Date()
	y := uint32(year) * yearMulti
	m := uint32(month) * monthMulti
	d := uint32(day) * dayMulti
	return (base | mask) + y + m + d
}

func ComputeSeedHour(base, yearMulti, monthMulti, dayMulti, hourMulti uint32, date time.Time) uint32 {
	seed := ComputeSeed(base, yearMulti, monthMulti, dayMulti, date)
	h := uint32(date.Hour()) * hourMulti
	return seed + h
}
