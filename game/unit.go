package game

type Unit struct {
	ID                  string
	X                   float64
	Y                   float64
	SpriteName          string
	Action              string
	Frame               int
	HorizontalDirection int
	IsDead              bool
}

type Units map[string]*Unit
