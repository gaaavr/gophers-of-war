package game

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	world *World
	frame int
}

func NewGame() Game {
	return Game{
		world: &World{
			Units: Units{},
		},
	}
}

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
	unitList := make([]*Unit, 0, len(g.world.Units))
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
		if unit.HorizontalDirection == DirectionLeft {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(32, 0)
		}
		op.GeoM.Translate(unit.X, unit.Y)
		spriteIndex := (g.frame / 7) % 4
		img, _, err = ebitenutil.NewImageFromFile("pictures/" + unit.SpriteName +
			"_" + unit.Action + "_" + strconv.Itoa(spriteIndex) + ".png")
		if err != nil {
			fmt.Println(err)
			return
		}
		screen.DrawImage(img, op)
	}
	for _, sh := range g.world.Shots {
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(sh.X, sh.Y)
		img, _, err = ebitenutil.NewImageFromFile("fire.png")
		if err != nil {
			fmt.Println(err)
			return
		}
		screen.DrawImage(img, op2)
	}
	if g.world.Units[g.world.MyID].IsDead {
		return
	}
	shot := false
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		shot = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.world.HandleEvent(&EventMove{
			Direction: DirectionRight,
			Shot:      shot,
		})
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.world.HandleEvent(&EventMove{
			Direction: DirectionLeft,
			Shot:      shot,
		})
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.world.HandleEvent(&EventMove{
			Direction: DirectionUp,
			Shot:      shot,
		})
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.world.HandleEvent(&EventMove{
			Direction: DirectionDown,
			Shot:      shot,
		})
		return
	}

	g.world.HandleEvent(&EventIdle{
		Shot: shot,
	})

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) AddPlayer() {
	g.world.addPlayer()
}
