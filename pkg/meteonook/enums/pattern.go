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

//go:generate go run github.com/alvaroloes/enumer -type=Pattern -json -yaml -text

package enums

type Pattern int

const (
	Fine00 Pattern = iota
	Fine01
	Fine02
	Fine03
	Fine04
	Fine05
	Fine06
	Cloud00
	Cloud01
	Cloud02
	Rain00
	Rain01
	Rain02
	Rain03
	Rain04
	Rain05
	FineCloud00
	FineCloud01
	FineCloud02
	CloudFine00
	CloudFine01
	CloudFine02
	FineRain00
	FineRain01
	FineRain02
	FineRain03
	CloudRain00
	CloudRain01
	CloudRain02
	RainCloud00
	RainCloud01
	RainCloud02
	Commun00
	EventDay00
)

func (i Pattern) Kind() PatternKind {
	switch i {
	case Fine00, Fine01, Fine02, Fine03, Fine04, Fine05, Fine06:
		return Fine
	case Cloud00, Cloud01, Cloud02:
		return Cloud
	case Rain00, Rain01, Rain02, Rain03, Rain04, Rain05:
		return Rain
	case FineCloud00, FineCloud01, FineCloud02:
		return FineCloud
	case CloudFine00, CloudFine01, CloudFine02:
		return CloudFine
	case FineRain00, FineRain01, FineRain02, FineRain03:
		return FineRain
	case CloudRain00, CloudRain01, CloudRain02:
		return CloudRain
	case RainCloud00, RainCloud01, RainCloud02:
		return RainCloud
	case Commun00:
		return Commun
	case EventDay00:
		return EventDay
	default:
		return -1
	}
}

func (i Pattern) ShowerLevel() ShowerLevel {
	switch i {
	case Fine00:
		return HeavyShower
	case Fine02, Fine04, Fine06:
		return LightShower
	default:
		return NoShower
	}
}
