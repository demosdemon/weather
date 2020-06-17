# Weather Code Breakdown

## WASM Functions

| code | name | signature |
| ---- | ---- | --------- |
| `a` | `___wasm_call_ctors` | `() -> void` |
| `b` | `_isSpecialDay` | `(isSouthern: u32, year: u32, month: u32, date: u32) -> u32` |
| `c` | `_getSnow` | `(isSouthern: bool, month: int, date: int) -> bool` |
| `d` | `_getCloud` | unknown |
| `e` | `_getSPWeather` | `(isSouthern: bool, month: int, date: int) -> (1 | 2)` |
| `f` | `_getFog` | |
| `g` | `_checkWaterFog` | |
| `h` | `_isHeavyShowerPattern` | |
| `i` | `_isLightShowerPattern` | |
| `j` | `_isRainbowPattern` | `(isSouthern: u32, seed: u32, year: u32, month: u32, date: u32, weather: u32) -> u32` |
| `k` | `_isAuroraPattern` | |
| `l` | `_getPattern` | |
| `m` | `_isPatternPossibleAtDate` | |
| `n` | `_getWeather` | |
| `o` | `_getWindPower` | |
| `p` | `_getWindPowerRange` | |
| `q` | `_canHaveShootingStars` | |
| `r` | `_queryStars` | |
| `s` | `_getStarAmount` | |
| `t` | `_getStarSecond` | |
| `u` | `_guessClear` | |
| `v` | `_guessAddType` | |
| `w` | `_guessAddPattern` | |
| `x` | `_guessAddMinute` | |
| `y` | `_guessAddSecond` | |
| `z` | `_guessAddRainbowDouble` | |
| `A` | `_searchInit` | |
| `B` | `_searchGetPercentage` | |
| `C` | `_searchCompleted` | |
| `D` | `_searchFailed` | |
| `E` | `_searchGetMaxResultCount` | |
| `F` | `_searchGetResultCount` | |
| `G` | `_searchGetResult` | |
| `H` | `_searchStep` | |

## Special Days

| code | name |
| ---- | ---- |
| 1 | Bunny Day |
| 2 | Fishing Tourney |
| 3 | Bug-Off |
| 4 | New Year's Eve Countdown |

## SP Weather

| code | name |
| ---- | ---- |
| 1 | Rainbow |
| 2 | Aurora |
