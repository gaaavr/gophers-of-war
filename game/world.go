package game

import (
	uuid "github.com/satori/go.uuid"
)

type World struct {
	MyID               string
	counterPlayerSkins int
	counterMobSkins    int
	Units
	Shots
	mobs
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
	world.resolveShots()
	world.resolveMobs(world.Units[world.MyID].X, world.Units[world.MyID].Y)
	event.handleEvent(world)
}

func (world *World) addPlayer() *Unit {
	skins := []string{"gopher"}
	id := uuid.NewV4().String()
	unit := &Unit{
		ID:         id,
		Action:     ActionIdle,
		X:          50,
		Y:          100,
		Frame:      1,
		SpriteName: skins[world.counterPlayerSkins],
	}
	world.Units[id] = unit
	world.MyID = id
	world.counterPlayerSkins++
	return unit
}

func (world *World) addMob() *Unit {
	skins := []string{"monster_1"}
	id := uuid.NewV4().String()
	unit := &Unit{
		ID:         id,
		Action:     ActionIdle,
		X:          180,
		Y:          90,
		Frame:      1,
		SpriteName: skins[world.counterMobSkins],
	}
	world.Units[id] = unit
	world.counterMobSkins++
	return unit
}
