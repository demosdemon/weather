package meteonook

type DataType int32

const (
	NoData DataType = iota
	None
	MeteorShower
	Rainbow
	Aurora
)
