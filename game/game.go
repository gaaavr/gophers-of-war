package game

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	uuid "github.com/satori/go.uuid"
)

type Game struct {
	level *Level
	w, h  int

	mousePanX, mousePanY int
	world                *World
	offscreen            *ebiten.Image
	frame                int
}

func NewGame() (*Game, error) {
	name := "level_1"

	lvl, err := NewLevel(name)
	if err != nil {
		return nil, fmt.Errorf("failed to load Level: %w", err)
	}

	return &Game{
		level:     lvl,
		mousePanX: math.MinInt32,
		mousePanY: math.MinInt32,
		world: &World{
			Units: Units{},
		},
	}, nil

}

func (g *Game) AddPlayer() *Unit {
	id := uuid.NewV4().String()
	unit := &Unit{
		ID:     id,
		Action: ActionIdle,
		X:      50,
		Y:      100,
		Frame:  1,
	}
	g.world.Units[id] = unit
	g.world.MyID = id

	return unit
}

func (g *Game) Update() error {
	g.frame++

	var event Handler

	event = &EventIdle{}

	var x, y int
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		x = DirectionLeft
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		x = DirectionRight
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		y = DirectionDown
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		y = DirectionUp
	}

	if x != 0 || y != 0 {
		event = &EventMove{
			DirectionX: x,
			DirectionY: y,
		}
	}

	g.world.HandleEvent(event)

	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render Level.
	g.renderLevel(screen)
	// Print game info.
	ebitenutil.DebugPrint(screen, fmt.Sprintf("KEYS WASD \nFPS  %0.0f\nTPS  %0.0f\n"+
		"\nCOORD  %0.0f, %0.0f", ebiten.ActualFPS(), ebiten.ActualTPS(),
		g.world.Units[g.world.MyID].X, g.world.Units[g.world.MyID].Y))
}

// Layout is called when the Game's layout changes.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.w, g.h = outsideWidth, outsideHeight
	return g.w, g.h
}

// renderLevel draws the current Level on the screen.
func (g *Game) renderLevel(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.Fill(color.Black)

	// Draw tiles
	for _, row := range g.level.tiles {
		for _, t := range row {
			if t.tileType == -1 {
				continue
			}

			op.GeoM.Reset()
			op.GeoM.Scale(3, 3)
			op.GeoM.Scale(1, 1)
			op.GeoM.Translate(t.XY())

			screen.DrawImage(t.tileSprite, op)
		}
	}

	for _, unit := range g.world.Units {
		op.GeoM.Reset()
		if unit.HorizontalDirection == DirectionLeft {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(32, 0)
		}
		op.GeoM.Translate(unit.X, unit.Y)
		spriteIndex := (g.frame / 7) % 4

		var i int
		if unit.Action == "run" {
			i = 1
		}

		img := g.level.assets.CharacterSprites[0][i][spriteIndex]

		screen.DrawImage(img, op)
	}
}
