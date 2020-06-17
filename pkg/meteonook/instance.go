package meteonook

import (
	"os"
	"strconv"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

type Instance struct {
	//mu       sync.Mutex
	instance wasmer.Instance
	memory   *wasmer.Memory
}

func newMemory() (*wasmer.Memory, error) {
	const wasmPageSize = 1 << 16
	const defaultMemorySize = wasmPageSize * 256

	initialMemory := uint32(defaultMemorySize)

	if v, ok := os.LookupEnv("WASM_INITIAL_MEMORY"); ok {
		v, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return nil, err
		}
		initialMemory = uint32(v)
	}

	pages := initialMemory / wasmPageSize
	return wasmer.NewMemory(pages, pages)
}

func NewInstance(bytes []byte) (*Instance, error) {
	memory, err := newMemory()
	if err != nil {
		return nil, err
	}

	imports := wasmer.NewImports()
	_, err = imports.Namespace("a").AppendMemory("memory", memory)
	if err != nil {
		return nil, err
	}

	i, err := wasmer.NewInstanceWithImports(bytes, imports)
	if err != nil {
		return nil, err
	}
	v := Instance{instance: i, memory: memory}
	if err := v._WASMCallConstructors(); err != nil {
		i.Close()
		return nil, err
	}
	return &v, nil
}

func (i *Instance) Close() {
	i.instance.Close()
	i.memory.Close()
}

func (i *Instance) do(k string, args []interface{}) (wasmer.Value, error) {
	//i.mu.Lock()
	v, err := i.instance.Exports[k](args...)
	//i.mu.Unlock()
	return v, err
}

func (i *Instance) f32(k string, args ...interface{}) (float32, error) {
	v, err := i.do(k, args)
	if err != nil {
		return 0, err
	}
	return v.ToF32(), nil
}

func (i *Instance) i32(k string, args ...interface{}) (int32, error) {
	v, err := i.do(k, args)
	if err != nil {
		return 0, err
	}
	return v.ToI32(), nil
}

func (i *Instance) void(k string, args ...interface{}) error {
	_, err := i.do(k, args)
	return err
}

func (i *Instance) _WASMCallConstructors() error {
	return i.void("a")
}

func (i *Instance) IsSpecialDay(isSouthern, year, month, date int32) (int32, error) {
	return i.i32("b", isSouthern, year, month, date)
}

func (i *Instance) GetSnow(isSouthern, month, date int32) (int32, error) {
	return i.i32("c", isSouthern, month, date)
}

//! unused
func (i *Instance) GetCloud(isSouthern, month, date int32) (int32, error) {
	return i.i32("d", isSouthern, month, date)
}

//! used for guessing
func (i *Instance) GetSPWeather(isSouthern, month, date int32) (int32, error) {
	return i.i32("e", isSouthern, month, date)
}

func (i *Instance) GetFog(isSouthern, month, date int32) (int32, error) {
	return i.i32("f", isSouthern, month, date)
}

func (i *Instance) CheckWaterFog(seed, year, month, date int32) (int32, error) {
	return i.i32("g", seed, year, month, date)
}

//! unnecessary
func (i *Instance) IsHeavyShowerPattern(isSouthern, month, date, pattern int32) (int32, error) {
	return i.i32("h", isSouthern, month, date, pattern)
}

//! unnecessary
func (i *Instance) IsLightShowerPattern(isSouthern, month, date, pattern int32) (int32, error) {
	return i.i32("i", isSouthern, month, date, pattern)
}

func (i *Instance) IsRainbowPattern(isSouthern, seed, year, month, date, pattern int32) (int32, error) {
	return i.i32("j", isSouthern, seed, year, month, date, pattern)
}

func (i *Instance) IsAuroraPattern(isSouthern, seed, year, month, date, pattern int32) (int32, error) {
	return i.i32("k", isSouthern, seed, year, month, date, pattern)
}

func (i *Instance) GetPattern(seed, isSouthern, year, month, date int32) (int32, error) {
	return i.i32("l", seed, isSouthern, year, month, date)
}

//! used for guessing
func (i *Instance) IsPatternPossibleAtDate(isSouthern, month, date, pattern int32) (int32, error) {
	return i.i32("m", isSouthern, month, date, pattern)
}

func (i *Instance) GetWeather(pattern, hour int32) (int32, error) {
	return i.i32("n", pattern, hour)
}

func (i *Instance) GetWindPower(isSouthern, seed, year, month, date, hour, pattern int32) (int32, error) {
	return i.i32("o", isSouthern, seed, year, month, date, hour, pattern)
}

func (i *Instance) GetWindPowerRange(hour, pattern int32) (int32, error) {
	return i.i32("p", hour, pattern)
}

func (i *Instance) CanHaveShootingStars(pattern, hour int32) (int32, error) {
	return i.i32("q", pattern, hour)
}

//! mutates state
func (i *Instance) QueryStars(pattern, seed, year, month, date, hour, minute int32) (int32, error) {
	return i.i32("r", pattern, seed, year, month, date, hour, minute)
}

//! depends on QueryStars
func (i *Instance) GetStarAmount() (int32, error) {
	return i.i32("s")
}

//! depends on QueryStars and GetStarAmount
func (i *Instance) GetStarSecond(index int32) (int32, error) {
	return i.i32("t", index)
}

//! mutates state
//! used for guessing
func (i *Instance) GuessClear() error {
	return i.void("u")
}

//! mutates state
//! used for guessing
func (i *Instance) GuessAddType(year, month, date, hour, pattern int32) error {
	return i.void("v", year, month, date, hour, pattern)
}

//! mutates state
//! used for guessing
func (i *Instance) GuessAddPattern(year, month, date, hour, pattern int32) error {
	return i.void("w", year, month, date, hour, pattern)
}

//! mutates state
//! used for guessing
func (i *Instance) GuessAddMinute(year, month, date, hour, minute, value int32) (int32, error) {
	return i.i32("x", year, month, date, hour, minute, value)
}

//! mutates state
//! used for guessing
func (i *Instance) GuessAddSecond(year, month, date, hour, minute, second int32) (int32, error) {
	return i.i32("y", year, month, date, hour, minute, second)
}

//! mutates state
//! used for guessing
func (i *Instance) GuessAddDoubleRainbow(year, month, date, double int32) error {
	return i.void("z", year, month, date, double)
}

//! mutates state
//! used for guessing
func (i *Instance) SearchInit(isSouthern int32) error {
	return i.void("A", isSouthern)
}

//! used for guessing
func (i *Instance) SearchGetPercentage() (float32, error) {
	return i.f32("B")
}

//! used for guessing
func (i *Instance) SearchCompleted() (int32, error) {
	return i.i32("C")
}

//! used for guessing
func (i *Instance) SearchFailed() (int32, error) {
	return i.i32("D")
}

//! used for guessing
func (i *Instance) SearchGetMaxResultCount() (int32, error) {
	return i.i32("E")
}

//! used for guessing
func (i *Instance) SearchGetResultCount() (int32, error) {
	return i.i32("F")
}

//! used for guessing
func (i *Instance) SearchGetResult(index int32) (int32, error) {
	return i.i32("G", index)
}

//! mutates state
//! used for guessing
func (i *Instance) SearchStep() error {
	return i.void("H")
}
