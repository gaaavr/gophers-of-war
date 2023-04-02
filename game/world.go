package game

import (
	uuid "github.com/satori/go.uuid"
)

type World struct {
	MyID         string
	counterSkins int
	Units
	Shots
}

const ActionRun = "idle"
const ActionIdle = "idle"

const DirectionUp = 0
const DirectionDown = 1
const DirectionLeft = 2
const DirectionRight = 3

func (ev *EventIdle) handleEvent(world *World) {
	unit := world.Units[world.MyID]
	unit.Action = ActionIdle
	if ev.Shot {
		shot := Shot{}
		shot.getShotOpts(unit.HorizontalDirection, unit.X, unit.Y)
		world.Shots = append(world.Shots, shot)
	}
}

// HandleEvent изменяет состояние мира в зависимости от произошедшего события
func (world *World) HandleEvent(event handler) {
	world.Shots = world.Shots.resolveShots()
	event.handleEvent(world)
}

func (world *World) addPlayer() *Unit {
	skins := []string{
		"gopher_1",
	}
	id := uuid.NewV4().String()
	unit := &Unit{
		ID:         id,
		Action:     ActionIdle,
		X:          150,
		Y:          110,
		Frame:      1,
		SpriteName: skins[world.counterSkins],
	}
	world.Units[id] = unit
	world.MyID = id
	world.counterSkins++
	return unit
}
