package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/gaaavr/gophers-of-war/game"
	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var store = map[float64]struct{}{}
var x float64

type Game struct {
	world *game.World
	conn  *websocket.Conn
	frame int
}

const url = "ws://127.0.0.1:3001/ws"

func (g *Game) Update() error {
	g.frame++
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	img, _, err := ebitenutil.NewImageFromFile("pictures/background.png")
	if err != nil {
		log.Fatal(err)
	}
	screen.DrawImage(img, nil)
	// Сделаем слайс для сортировки юнитов, чтобы тот кто по Y находится
	// дальше, был ниже по слою
	unitList := make([]*game.Unit, 0, len(g.world.Units))
	for _, unit := range g.world.Units {
		if unit.IsDead {
			continue
		}
		unitList = append(unitList, unit)
	}
	sort.Slice(unitList, func(i, j int) bool {
		return unitList[i].Y < unitList[j].Y
	})
	// Отрисовываем юнитов по возрастанию, кто ближе по Y, тот будет выше слоем
	for _, unit := range unitList {
		op := &ebiten.DrawImageOptions{}
		op2 := &ebiten.DrawImageOptions{}
		if unit.HorizontalDirection == game.DirectionLeft {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(16, 0)
		}
		op.GeoM.Translate(unit.X, unit.Y)
		spriteIndex := (g.frame/7 + unit.Frame) % 4
		img, _, err = ebitenutil.NewImageFromFile("pictures/" + unit.SpriteName +
			"_" + unit.Action + "_" + strconv.Itoa(spriteIndex) + ".png")
		if err != nil {
			fmt.Println(err)
			return
		}
		screen.DrawImage(img, op)
		if unit.IsShot {
			op2.GeoM.Translate(unit.Shot.X, unit.Shot.Y)
			img, _, _ = ebitenutil.NewImageFromFile("fire.png")
			screen.DrawImage(img, op2)
		}
	}
	if g.world.Units[g.world.MyID].IsDead {
		return
	}
	shot := false
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		shot = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.conn.WriteJSON(game.Event{
			Type: game.EventTypeMove,
			Data: game.EventMove{
				UnitID:    g.world.MyID,
				Direction: game.DirectionRight,
				Shot:      shot,
			},
		})
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.conn.WriteJSON(game.Event{
			Type: game.EventTypeMove,
			Data: game.EventMove{
				UnitID:    g.world.MyID,
				Direction: game.DirectionLeft,
				Shot:      shot,
			},
		})
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.conn.WriteJSON(game.Event{
			Type: game.EventTypeMove,
			Data: game.EventMove{
				UnitID:    g.world.MyID,
				Direction: game.DirectionUp,
				Shot:      shot,
			},
		})
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.conn.WriteJSON(game.Event{
			Type: game.EventTypeMove,
			Data: game.EventMove{
				UnitID:    g.world.MyID,
				Direction: game.DirectionDown,
				Shot:      shot,
			},
		})
		return
	}
	if g.world.Units[g.world.MyID].Action == game.ActionRun || g.world.Units[g.world.MyID].Action == game.EventTypeIdle {
		g.conn.WriteJSON(game.Event{
			Type: game.EventTypeIdle,
			Data: game.EventMove{
				UnitID: g.world.MyID,
				Shot:   shot,
			},
		})
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatalf("connection error: %s", err.Error())
	}
	g := Game{
		world: &game.World{
			IsServer: false,
			Units:    game.Units{},
		},
		conn: conn,
	}
	go func(conn *websocket.Conn) {
		defer conn.Close()
		for {
			var event game.Event
			conn.ReadJSON(&event)
			g.world.HandleEvent(&event)
			log.Println(event)
		}

	}(conn)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gophers of war")
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
