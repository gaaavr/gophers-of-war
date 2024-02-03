package game

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	uuid "github.com/satori/go.uuid"
)

const (
	// tile
	TileSize = 16 * 3
)

type Tile struct {
	*Rect

	tileType   int
	tileSprite *ebiten.Image
	isObstacle bool
}

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) XY() (float64, float64) {
	return r.x, r.y
}

type Game struct {
	hero  [][]*ebiten.Image
	tiles [][]*Tile
	w, h  int

	camX, camY float64
	camScale   float64
	camScaleTo float64

	mousePanX, mousePanY int
	world                *World
	offscreen            *ebiten.Image
	frame                int
}

func NewGame() (*Game, error) {
	assets, err := LoadAssets()
	if err != nil {
		return nil, fmt.Errorf("failed to load assets: %w", err)
	}

	name := "level_1"
	levelData, err := loadLevelData(name)
	if err != nil {
		return nil, fmt.Errorf("failed to load level data: %w", err)
	}

	tiles := make([][]*Tile, len(levelData))
	for y := range tiles {
		tiles[y] = make([]*Tile, len(levelData[0]))
	}

	for y, row := range levelData {
		for x, col := range row {
			//  empty tile
			if col == -1 {
				tiles[y][x] = &Tile{tileType: -1}
				continue
			}

			// not empty tile
			image := assets.TileSprites[col]
			tile := &Tile{
				tileType:   col,
				tileSprite: image,
				isObstacle: false,
				Rect: &Rect{
					x:      float64(x * TileSize),
					y:      float64(y * TileSize),
					width:  TileSize,
					height: TileSize,
				},
			}

			tiles[y][x] = tile
		}
	}

	return &Game{
		hero:       assets.CharacterSprites[0],
		tiles:      tiles,
		camScale:   1,
		camScaleTo: 1,
		mousePanX:  math.MinInt32,
		mousePanY:  math.MinInt32,
		world: &World{
			Units: Units{},
		},
	}, nil

}

func (g *Game) AddPlayer() *Unit {
	skins := []string{"gopher"}
	id := uuid.NewV4().String()
	unit := &Unit{
		ID:         id,
		Action:     ActionIdle,
		X:          50,
		Y:          100,
		Frame:      1,
		SpriteName: skins[0],
	}
	g.world.Units[id] = unit
	g.world.MyID = id
	return unit
}

func (g *Game) Update() error {
	g.frame++

	var event Handler

	shot := false
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		shot = true
	}

	event = &EventIdle{
		Shot: shot,
	}

	var x, y int
	// Pan camera via keyboard.
	pan := 1.0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		x = DirectionLeft
		g.camX -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		x = DirectionRight
		g.camX += pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		y = DirectionDown
		g.camY -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		y = DirectionUp
		g.camY += pan
	}

	if x != 0 || y != 0 {
		event = &EventMove{
			DirectionX: x,
			DirectionY: y,
			Shot:       shot,
		}
	}

	g.world.HandleEvent(event)

	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render level.
	g.renderLevel(screen)
	// Print game info.
	ebitenutil.DebugPrint(screen, fmt.Sprintf("KEYS WASD \nFPS  %0.0f\nTPS  %0.0f\nSCA  %0.2f\nPOS  "+
		"%0.0f,%0.0f\nCOORD  %0.0f, %0.0f", ebiten.ActualFPS(), ebiten.ActualTPS(),
		g.camScale, g.camX, g.camY, g.world.Units[g.world.MyID].X, g.world.Units[g.world.MyID].Y))
}

// Layout is called when the Game's layout changes.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.w, g.h = outsideWidth, outsideHeight
	return g.w, g.h
}

// renderLevel draws the current Level on the screen.
func (g *Game) renderLevel(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	img, _, err := ebitenutil.NewImageFromFile("assets/images/background.png")
	if err != nil {
		log.Fatal(err)
	}

	screen.DrawImage(img, op)
	screen.Fill(color.Black)

	// Draw tiles
	for _, row := range g.tiles {
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

		img := g.hero[i][spriteIndex]

		screen.DrawImage(img, op)
	}
}
