package game

type World struct {
	MyID string
	Units
}

const ActionRun = "run"
const ActionIdle = "idle"

const (
	DirectionUp = iota + 1
	DirectionDown
	DirectionLeft
	DirectionRight
)

// HandleEvent изменяет состояние мира в зависимости от произошедшего события
func (world *World) HandleEvent(event Handler) {
	event.handleEvent(world)
}
