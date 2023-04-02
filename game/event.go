package game

const EventTypeConnect = "connect"
const EventTypeMove = "move"
const EventTypeIdle = "idle"
const EventTypeInit = "init"

type EventMove struct {
	Direction int  `json:"direction"`
	Shot      bool `json:"shot"`
	IsDead    bool `json:"is_dead"`
}

func (ev *EventMove) handleEvent(world *World) {
	unit := world.Units[world.MyID]
	unit.Action = ActionRun
	if ev.Shot {
		shot := Shot{}
		shot.getShotOpts(unit.HorizontalDirection, unit.X, unit.Y)
		world.Shots = append(world.Shots, shot)
	}
	switch ev.Direction {
	case DirectionUp:
		unit.Y--
	case DirectionDown:
		unit.Y++
	case DirectionLeft:
		unit.X--
		unit.HorizontalDirection = ev.Direction
	case DirectionRight:
		unit.X++
		unit.HorizontalDirection = ev.Direction
	}
}

type EventIdle struct {
	Shot   bool `json:"shot"`
	IsDead bool `json:"is_dead"`
}
