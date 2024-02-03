package game

import (
	"math/rand"
	"time"
)

var counter int

type EventMove struct {
	DirectionX int  `json:"direction"`
	DirectionY int  `json:"direction"`
	Shot       bool `json:"shot"`
}

func (ev *EventMove) handleEvent(world *World) {
	unit := world.Units[world.MyID]
	unit.Action = ActionRun
	if ev.Shot {
		counter++
		shot := Shot{}
		shot.getShotOpts(unit.HorizontalDirection, unit.X, unit.Y)
		world.Shots = append(world.Shots, shot)
		if counter%5 == 0 {
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			world.mobs = append(world.mobs, mob{X: rnd.Float64()*300 + 10, Y: rnd.Float64()*220 + 10, spriteName: "monster_1"})
		}
	}

	pan1, pan2 := 0.0, 0.0
	switch ev.DirectionX {
	case DirectionLeft:
		pan1 = -1
		unit.HorizontalDirection = ev.DirectionX
	case DirectionRight:
		unit.HorizontalDirection = ev.DirectionX
		pan1 = 1
	}

	switch ev.DirectionY {
	case DirectionUp:
		pan2 = -1
	case DirectionDown:
		pan2 = 1
	}

	if ev.DirectionX != 0 {
		unit.X += pan1
	}

	if ev.DirectionY != 0 {
		unit.Y += pan2
	}

}

type EventIdle struct {
	Shot bool `json:"shot"`
}

func (ev *EventIdle) handleEvent(world *World) {
	unit := world.Units[world.MyID]
	unit.Action = ActionIdle
	if ev.Shot {
		counter++
		shot := Shot{}
		shot.getShotOpts(unit.HorizontalDirection, unit.X, unit.Y)
		world.Shots = append(world.Shots, shot)
		if counter%5 == 0 {
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			world.mobs = append(world.mobs, mob{X: rnd.Float64()*300 + 10, Y: rnd.Float64()*220 + 10, spriteName: "monster_1"})
		}
	}
}
