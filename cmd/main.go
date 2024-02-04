package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/gaaavr/gophers-of-war/game"
)

const (
	ScreenWidth  = 1200
	ScreenHeight = 912
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	g.AddPlayer()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Gophers of war")
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
