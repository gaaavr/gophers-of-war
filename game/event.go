package game

import (
	"math/rand"
	"time"
)

var counter int

type EventMove struct {
	Direction int  `json:"direction"`
	Shot      bool `json:"shot"`
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
