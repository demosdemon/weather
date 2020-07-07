# weather

Calculate Animal Crossing: New Horizons weather patterns in go.

## Usage

Visit [MeteoNook](https://wuffs.org/acnh/weather) to calculate your game's `seed`.

All GET endpoints accept an optional `timezone` query parameter. 

    https://serene-garden-64785.herokuapp.com/v1/1160352294/N/%23island/today.json?timezone=America/Chicago

* `seed` is an unsigned 32-bit integer.
* `hemisphere` is one of "N" or "S"
* `name` is purely cosmetic

### Today

*GET `/v1/{seed:\d+}/{hemisphere:N|S}/{name}/today.json`* 

```
$ curl 'https://serene-garden-64785.herokuapp.com/v1/1160352294/N/%23island/today.json'
{
  "island": {
    "name": "#island",
    "hemisphere": "Northern",
    "seed": 1160352294,
    "timezone": "UTC"
  },
  "year": 2020,
  "month": 7,
  "date": 7,
  "weekday": 2,
  "constellation": "Cancer",
  "pattern": "Fine05",
  "hours": [
    {
      "hour": "05:00",
      "weather": "Sunny",
      "wind_power": 2
    },
    {
      "hour": "06:00",
      "weather": "Sunny",
      "wind_power": 1
    },
    {
      "hour": "07:00",
      "weather": "Sunny"
    },
    {
      "hour": "08:00",
      "weather": "Clear/Fine",
      "wind_power": 1
    },
    {
      "hour": "09:00",
      "weather": "Sunny",
      "wind_power": 1
    },
    {
      "hour": "10:00",
      "weather": "Sunny",
      "wind_power": 2
    },
    {
      "hour": "11:00",
      "weather": "Clear/Fine",
      "wind_power": 4
    },
    {
      "hour": "12:00",
      "weather": "Sunny",
      "wind_power": 2
    },
    {
      "hour": "13:00",
      "weather": "Sunny",
      "wind_power": 1
    },
    {
      "hour": "14:00",
      "weather": "Cloudy",
      "wind_power": 2
    },
    {
      "hour": "15:00",
      "weather": "Sunny",
      "wind_power": 4
    },
    {
      "hour": "16:00",
      "weather": "Sunny",
      "wind_power": 2
    },
    {
      "hour": "17:00",
      "weather": "Clear/Fine",
      "wind_power": 1
    },
    {
      "hour": "18:00",
      "weather": "Sunny"
    },
    {
      "hour": "19:00",
      "weather": "Cloudy",
      "wind_power": 1
    },
    {
      "hour": "20:00",
      "weather": "Sunny",
      "wind_power": 1
    },
    {
      "hour": "21:00",
      "weather": "Sunny",
      "wind_power": 1
    },
    {
      "hour": "22:00",
      "weather": "Clear/Fine",
      "wind_power": 1
    },
    {
      "hour": "23:00",
      "weather": "Sunny",
      "wind_power": 2
    },
    {
      "hour": "00:00",
      "weather": "Cloudy",
      "wind_power": 2
    },
    {
      "hour": "01:00",
      "weather": "Sunny",
      "wind_power": 4
    },
    {
      "hour": "02:00",
      "weather": "Sunny",
      "wind_power": 5
    },
    {
      "hour": "03:00",
      "weather": "Clear/Fine",
      "wind_power": 4
    },
    {
      "hour": "04:00",
      "weather": "Sunny",
      "wind_power": 4
    }
  ]
}
```

### Specific Date

*GET `/v1/{seed:\d+}/{hemisphere:N|S}/{name}/{date:YYYY-MM-DD}.json`*

```
$ curl 'https://serene-garden-64785.herokuapp.com/v1/1160352294/N/%23island/2020-07-27.json'
{
  "island": {
    "name": "#island",
    "hemisphere": "Northern",
    "seed": 1160352294,
    "timezone": "UTC"
  },
  "year": 2020,
  "month": 7,
  "date": 27,
  "weekday": 1,
  "constellation": "Leo",
  "cloud_level": "Cumulonimbus",
  "pattern": "Fine06",
  "shower_type": "LightShower",
  "hours": [
    {
      "hour": "05:00",
      "weather": "Sunny",
      "wind_power": 3
    },
    {
      "hour": "06:00",
      "weather": "Cloudy",
      "wind_power": 1
    },
    {
      "hour": "07:00",
      "weather": "Sunny"
    },
    {
      "hour": "08:00",
      "weather": "Sunny",
      "wind_power": 1
    },
    {
      "hour": "09:00",
      "weather": "Cloudy",
      "wind_power": 1
    },
    {
      "hour": "10:00",
      "weather": "Clear/Fine",
      "wind_power": 3
    },
    {
      "hour": "11:00",
      "weather": "Sunny",
      "wind_power": 1
    },
    {
      "hour": "12:00",
      "weather": "Clear/Fine",
      "wind_power": 1
    },
    {
      "hour": "13:00",
      "weather": "Sunny",
      "wind_power": 3
    },
    {
      "hour": "14:00",
      "weather": "Sunny",
      "wind_power": 4
    },
    {
      "hour": "15:00",
      "weather": "Clear/Fine",
      "wind_power": 4
    },
    {
      "hour": "16:00",
      "weather": "Sunny",
      "wind_power": 3
    },
    {
      "hour": "17:00",
      "weather": "Cloudy",
      "wind_power": 1
    },
    {
      "hour": "18:00",
      "weather": "Cloudy"
    },
    {
      "hour": "19:00",
      "weather": "Sunny",
      "wind_power": 1,
      "shower_level": "LightShower",
      "shooting_stars": [
        "2020-07-27T19:23:00Z",
        "2020-07-27T19:23:17Z",
        "2020-07-27T19:23:29Z",
        "2020-07-27T19:23:38Z",
        "2020-07-27T19:23:58Z",
        "2020-07-27T19:31:03Z",
        "2020-07-27T19:31:11Z",
        "2020-07-27T19:31:27Z",
        "2020-07-27T19:31:41Z",
        "2020-07-27T19:31:57Z",
        "2020-07-27T19:41:03Z",
        "2020-07-27T19:41:19Z",
        "2020-07-27T19:41:34Z",
        "2020-07-27T19:41:40Z",
        "2020-07-27T19:41:55Z",
        "2020-07-27T19:42:05Z",
        "2020-07-27T19:42:09Z",
        "2020-07-27T19:42:15Z",
        "2020-07-27T19:42:23Z",
        "2020-07-27T19:42:28Z"
      ]
    },
    {
      "hour": "20:00",
      "weather": "Sunny",
      "wind_power": 1,
      "shower_level": "LightShower",
      "shooting_stars": [
        "2020-07-27T20:03:04Z",
        "2020-07-27T20:03:12Z",
        "2020-07-27T20:03:31Z",
        "2020-07-27T20:03:36Z",
        "2020-07-27T20:03:49Z",
        "2020-07-27T20:47:14Z",
        "2020-07-27T20:47:23Z",
        "2020-07-27T20:47:33Z",
        "2020-07-27T20:47:52Z",
        "2020-07-27T20:47:57Z"
      ]
    },
    {
      "hour": "21:00",
      "weather": "Sunny",
      "wind_power": 3,
      "shower_level": "LightShower",
      "shooting_stars": [
        "2020-07-27T21:17:15Z",
        "2020-07-27T21:17:30Z",
        "2020-07-27T21:17:35Z",
        "2020-07-27T21:17:37Z",
        "2020-07-27T21:17:49Z",
        "2020-07-27T21:45:01Z",
        "2020-07-27T21:45:14Z",
        "2020-07-27T21:45:17Z",
        "2020-07-27T21:45:41Z",
        "2020-07-27T21:45:58Z"
      ]
    },
    {
      "hour": "22:00",
      "weather": "Clear/Fine",
      "wind_power": 3,
      "shower_level": "LightShower",
      "shooting_stars": [
        "2020-07-27T22:23:01Z",
        "2020-07-27T22:23:13Z",
        "2020-07-27T22:23:21Z",
        "2020-07-27T22:23:29Z",
        "2020-07-27T22:23:49Z",
        "2020-07-27T22:44:05Z",
        "2020-07-27T22:44:07Z",
        "2020-07-27T22:44:44Z",
        "2020-07-27T22:44:49Z",
        "2020-07-27T22:44:59Z",
        "2020-07-27T22:55:13Z",
        "2020-07-27T22:55:24Z",
        "2020-07-27T22:55:29Z",
        "2020-07-27T22:55:52Z",
        "2020-07-27T22:55:56Z"
      ]
    },
    {
      "hour": "23:00",
      "weather": "Clear/Fine",
      "wind_power": 1,
      "shower_level": "LightShower",
      "shooting_stars": [
        "2020-07-27T23:25:09Z",
        "2020-07-27T23:25:47Z",
        "2020-07-27T23:25:48Z",
        "2020-07-27T23:25:54Z",
        "2020-07-27T23:25:55Z",
        "2020-07-27T23:39:10Z",
        "2020-07-27T23:39:30Z",
        "2020-07-27T23:39:47Z",
        "2020-07-27T23:39:49Z",
        "2020-07-27T23:39:53Z"
      ]
    },
    {
      "hour": "00:00",
      "weather": "Clear/Fine",
      "wind_power": 2,
      "shower_level": "LightShower",
      "shooting_stars": [
        "2020-07-28T00:12:10Z",
        "2020-07-28T00:12:11Z",
        "2020-07-28T00:12:15Z",
        "2020-07-28T00:12:43Z",
        "2020-07-28T00:12:46Z",
        "2020-07-28T00:20:14Z",
        "2020-07-28T00:20:15Z",
        "2020-07-28T00:20:18Z",
        "2020-07-28T00:20:25Z",
        "2020-07-28T00:20:42Z",
        "2020-07-28T00:37:14Z",
        "2020-07-28T00:37:25Z",
        "2020-07-28T00:37:32Z",
        "2020-07-28T00:37:45Z",
        "2020-07-28T00:37:54Z"
      ]
    },
    {
      "hour": "01:00",
      "weather": "Sunny",
      "wind_power": 4,
      "shower_level": "LightShower",
      "shooting_stars": [
        "2020-07-28T01:58:32Z",
        "2020-07-28T01:58:35Z",
        "2020-07-28T01:58:43Z",
        "2020-07-28T01:58:58Z",
        "2020-07-28T01:58:59Z"
      ]
    },
    {
      "hour": "02:00",
      "weather": "Clear/Fine",
      "wind_power": 4,
      "shower_level": "LightShower",
      "shooting_stars": [
        "2020-07-28T02:10:02Z",
        "2020-07-28T02:10:09Z",
        "2020-07-28T02:10:19Z",
        "2020-07-28T02:10:30Z",
        "2020-07-28T02:10:42Z",
        "2020-07-28T02:14:14Z",
        "2020-07-28T02:14:20Z",
        "2020-07-28T02:14:23Z",
        "2020-07-28T02:14:27Z",
        "2020-07-28T02:14:47Z",
        "2020-07-28T02:37:10Z",
        "2020-07-28T02:37:20Z",
        "2020-07-28T02:37:35Z",
        "2020-07-28T02:37:45Z",
        "2020-07-28T02:37:49Z",
        "2020-07-28T02:42:04Z",
        "2020-07-28T02:42:24Z",
        "2020-07-28T02:42:33Z",
        "2020-07-28T02:42:53Z",
        "2020-07-28T02:42:54Z",
        "2020-07-28T02:45:04Z",
        "2020-07-28T02:45:09Z",
        "2020-07-28T02:45:23Z",
        "2020-07-28T02:45:34Z",
        "2020-07-28T02:45:54Z"
      ]
    },
    {
      "hour": "03:00",
      "weather": "Sunny",
      "wind_power": 5,
      "shower_level": "LightShower",
      "shooting_stars": [
        "2020-07-28T03:03:04Z",
        "2020-07-28T03:03:12Z",
        "2020-07-28T03:03:26Z",
        "2020-07-28T03:03:37Z",
        "2020-07-28T03:03:47Z",
        "2020-07-28T03:15:08Z",
        "2020-07-28T03:15:26Z",
        "2020-07-28T03:15:30Z",
        "2020-07-28T03:15:43Z",
        "2020-07-28T03:15:49Z",
        "2020-07-28T03:24:10Z",
        "2020-07-28T03:24:18Z",
        "2020-07-28T03:24:33Z",
        "2020-07-28T03:24:54Z",
        "2020-07-28T03:24:57Z"
      ]
    },
    {
      "hour": "04:00",
      "weather": "Sunny",
      "wind_power": 5
    }
  ]
}
```

### Feed

*GET `/v1/{seed:\d+}/{hemisphere:N|S}/{name}/feed.json?first_date=YYYY-MM-DD&last_date=YYYY-MM-DD`*

`first_date` and `last_date` are optional. If not provided, `first_date` is `today - 3 months` and `last_date` is `first_date + 1 year - 1 day`. Dates are inclusive.

```
$ curl 'https://serene-garden-64785.herokuapp.com/v1/1160352294/N/%23island/feed.json?first_date=2020-07-01&last_date=2020-07-02'
{
  "island": {
    "name": "#island",
    "hemisphere": "Northern",
    "seed": 1160352294,
    "timezone": "UTC"
  },
  "days": [
    {
      "island": {
        "name": "#island",
        "hemisphere": "Northern",
        "seed": 1160352294,
        "timezone": "UTC"
      },
      "year": 2020,
      "month": 7,
      "date": 1,
      "weekday": 3,
      "constellation": "Cancer",
      "pattern": "CloudFine02",
      "rainbow_info": {
        "hour": "12:00",
        "type": "Double Rainbow"
      },
      "hours": [
        {
          "hour": "05:00",
          "weather": "Clear/Fine",
          "wind_power": 4
        },
        {
          "hour": "06:00",
          "weather": "Cloudy",
          "wind_power": 2
        },
        {
          "hour": "07:00",
          "weather": "Storm Clouds"
        },
        {
          "hour": "08:00",
          "weather": "Light Storm",
          "wind_power": 2
        },
        {
          "hour": "09:00",
          "weather": "Cloudy",
          "wind_power": 2
        },
        {
          "hour": "10:00",
          "weather": "Storm Clouds",
          "wind_power": 4
        },
        {
          "hour": "11:00",
          "weather": "Light Storm",
          "wind_power": 2,
          "rainbow_type": "Double Rainbow"
        },
        {
          "hour": "12:00",
          "weather": "Sunny",
          "wind_power": 4,
          "rainbow_type": "Double Rainbow"
        },
        {
          "hour": "13:00",
          "weather": "Sunny",
          "wind_power": 5
        },
        {
          "hour": "14:00",
          "weather": "Sunny",
          "wind_power": 4
        },
        {
          "hour": "15:00",
          "weather": "Sunny",
          "wind_power": 5
        },
        {
          "hour": "16:00",
          "weather": "Clear/Fine",
          "wind_power": 4
        },
        {
          "hour": "17:00",
          "weather": "Sunny",
          "wind_power": 2
        },
        {
          "hour": "18:00",
          "weather": "Clear/Fine"
        },
        {
          "hour": "19:00",
          "weather": "Sunny",
          "wind_power": 2
        },
        {
          "hour": "20:00",
          "weather": "Clear/Fine",
          "wind_power": 2
        },
        {
          "hour": "21:00",
          "weather": "Clear/Fine",
          "wind_power": 2
        },
        {
          "hour": "22:00",
          "weather": "Sunny",
          "wind_power": 4
        },
        {
          "hour": "23:00",
          "weather": "Clear/Fine",
          "wind_power": 5
        },
        {
          "hour": "00:00",
          "weather": "Sunny",
          "wind_power": 2
        },
        {
          "hour": "01:00",
          "weather": "Sunny",
          "wind_power": 1
        },
        {
          "hour": "02:00",
          "weather": "Sunny",
          "wind_power": 2
        },
        {
          "hour": "03:00",
          "weather": "Sunny",
          "wind_power": 4
        },
        {
          "hour": "04:00",
          "weather": "Sunny",
          "wind_power": 2
        }
      ]
    },
    {
      "island": {
        "name": "#island",
        "hemisphere": "Northern",
        "seed": 1160352294,
        "timezone": "UTC"
      },
      "year": 2020,
      "month": 7,
      "date": 2,
      "weekday": 4,
      "constellation": "Cancer",
      "pattern": "Rain01",
      "hours": [
        {
          "hour": "05:00",
          "weather": "Light Storm",
          "wind_power": 2
        },
        {
          "hour": "06:00",
          "weather": "Light Storm"
        },
        {
          "hour": "07:00",
          "weather": "Light Storm",
          "wind_power": 1
        },
        {
          "hour": "08:00",
          "weather": "Light Storm",
          "wind_power": 1
        },
        {
          "hour": "09:00",
          "weather": "Cloudy",
          "wind_power": 2
        },
        {
          "hour": "10:00",
          "weather": "Storm Clouds",
          "wind_power": 2
        },
        {
          "hour": "11:00",
          "weather": "Light Storm",
          "wind_power": 4
        },
        {
          "hour": "12:00",
          "weather": "Light Storm",
          "wind_power": 2
        },
        {
          "hour": "13:00",
          "weather": "Cloudy",
          "wind_power": 4
        },
        {
          "hour": "14:00",
          "weather": "Cloudy",
          "wind_power": 2
        },
        {
          "hour": "15:00",
          "weather": "Storm Clouds",
          "wind_power": 1
        },
        {
          "hour": "16:00",
          "weather": "Heavy Storm",
          "wind_power": 2
        },
        {
          "hour": "17:00",
          "weather": "Light Storm",
          "wind_power": 1
        },
        {
          "hour": "18:00",
          "weather": "Light Storm",
          "wind_power": 1
        },
        {
          "hour": "19:00",
          "weather": "Light Storm"
        },
        {
          "hour": "20:00",
          "weather": "Light Storm",
          "wind_power": 1
        },
        {
          "hour": "21:00",
          "weather": "Sunny",
          "wind_power": 2
        },
        {
          "hour": "22:00",
          "weather": "Sunny",
          "wind_power": 2
        },
        {
          "hour": "23:00",
          "weather": "Cloudy",
          "wind_power": 1
        },
        {
          "hour": "00:00",
          "weather": "Storm Clouds",
          "wind_power": 2
        },
        {
          "hour": "01:00",
          "weather": "Light Storm",
          "wind_power": 4
        },
        {
          "hour": "02:00",
          "weather": "Light Storm",
          "wind_power": 4
        },
        {
          "hour": "03:00",
          "weather": "Light Storm",
          "wind_power": 5
        },
        {
          "hour": "04:00",
          "weather": "Light Storm",
          "wind_power": 5
        }
      ]
    }
  ]
}
```

### In Code

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/demosdemon/weather/pkg/meteonook"
	"github.com/demosdemon/weather/pkg/meteonook/enums"
)

func main() {
	island := meteonook.Island{
		Name:       "#island",
		Hemisphere: enums.Northern,
		Seed:       1160352294,
		Timezone:   meteonook.Timezone{Location: time.Local},
	}

	today := time.Now().In(time.Local)
	day, err := island.NewDay(today)

	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v\n", day)
}
```

## Thanks

Thanks to [@_Ninji](https://twitter.com/_Ninji) for his work on [MeteoNook](https://wuffs.org/acnh/weather)

## License

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
