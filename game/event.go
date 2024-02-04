package game

type EventMove struct {
	DirectionX int
	DirectionY int
}

func (ev *EventMove) handleEvent(world *World) {
	unit := world.Units[world.MyID]
	unit.Action = ActionRun

	pan1, pan2 := 0.0, 0.0
	switch ev.DirectionX {
	case DirectionLeft:
		pan1 = -3
		unit.HorizontalDirection = ev.DirectionX
	case DirectionRight:
		unit.HorizontalDirection = ev.DirectionX
		pan1 = 3
	}

	switch ev.DirectionY {
	case DirectionUp:
		pan2 = -3
	case DirectionDown:
		pan2 = 3
	}

	if ev.DirectionX != 0 {
		unit.X += pan1
	}

	if ev.DirectionY != 0 {
		unit.Y += pan2
	}

}

type EventIdle struct {
}

func (ev *EventIdle) handleEvent(world *World) {
	unit := world.Units[world.MyID]
	unit.Action = ActionIdle
}
