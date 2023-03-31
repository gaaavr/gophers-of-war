package game

import (
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Unit struct {
	ID                  string  `json:"id"`
	X                   float64 `json:"x"`
	Y                   float64 `json:"y"`
	SpriteName          string  `json:"sprite_name"`
	Action              string  `json:"action"`
	Frame               int     `json:"frame"`
	HorizontalDirection int     `json:"direction"`
	IsShot              bool    `json:"is_shot"`
	IsDead              bool    `json:"is_dead"`
	Shot
}

type Units map[string]*Unit

type World struct {
	MyID     string `json:"-"`
	IsServer bool   `json:"-"`
	Units    `json:"units"`
}

type Event struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type EventConnect struct {
	Unit
}

type EventMove struct {
	UnitID    string `json:"unit_id"`
	Direction int    `json:"direction"`
	Shot      bool   `json:"shot"`
	IsDead    bool   `json:"is_dead"`
}

type Shot struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type EventIdle struct {
	UnitID string `json:"unit_id"`
	Shot   bool   `json:"shot"`
	IsDead bool   `json:"is_dead"`
}

type EventInit struct {
	PlayerID string `json:"player_id"`
	Units    Units  `json:"units"`
}

const EventTypeConnect = "connect"
const EventTypeMove = "move"
const EventTypeIdle = "idle"
const EventTypeInit = "init"

const ActionRun = "idle"
const ActionIdle = "idle"

const DirectionUp = 0
const DirectionDown = 1
const DirectionLeft = 2
const DirectionRight = 3

func (ev *EventMove) handleEvent(world *World) {
	unit := world.Units[ev.UnitID]
	unit.Action = ActionRun
	if ev.Shot {
		unit.IsShot = ev.Shot
	}
	if unit.IsShot && ev.Shot {
		unit.Shot.X = unit.X
		unit.Shot.Y = unit.Y
	}
	if unit.IsShot {
		for _, unit2 := range world.Units {
			if unit2.X >= unit.Shot.X && unit2.ID != unit.ID && !unit2.IsDead {
				world.Units[unit2.ID].IsDead = true
				unit.IsShot = false
				return
			}
		}
	}
	if unit.Shot.X > 10 {
		unit.Shot.X -= 3
	} else {
		unit.IsShot = false
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

func (ev *EventIdle) handleEvent(world *World) {
	unit := world.Units[ev.UnitID]
	unit.Action = ActionIdle
	if ev.Shot {
		unit.IsShot = ev.Shot
	}
	if unit.IsShot && ev.Shot {
		unit.Shot.X = unit.X
		unit.Shot.Y = unit.Y
	}
	if unit.IsShot {
		for _, unit2 := range world.Units {
			if unit2.X >= unit.Shot.X && unit2.ID != unit.ID && !unit2.IsDead {
				world.Units[unit2.ID].IsDead = true
				unit.IsShot = false
				return
			}
		}
	}
	if unit.Shot.X > 10 {
		unit.Shot.X -= 3
	} else {
		unit.IsShot = false
	}
}

func (ev *EventInit) handleEvent(world *World) {
	if !world.IsServer {
		world.MyID = ev.PlayerID
		world.Units = ev.Units
	}
}

func (ev *EventConnect) handleEvent(world *World) {
	world.Units[ev.ID] = &ev.Unit
}

// HandleEvent изменяет состояние мира в зависимости от произошедшего события
func (world *World) HandleEvent(event handler) {
	event.handleEvent(world)
}

func (world *World) AddPlayer() *Unit {
	skins := []string{
		"gopher",
	}
	id := uuid.NewV4().String()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	unit := &Unit{
		ID:         id,
		Action:     ActionIdle,
		X:          rnd.Float64() * 320,
		Y:          rnd.Float64() * 240,
		Frame:      rnd.Intn(4),
		SpriteName: skins[rnd.Intn(len(skins))],
	}
	world.Units[id] = unit
	world.MyID = id
	return unit
}
