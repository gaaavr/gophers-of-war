package main

import (
	"log"

	"github.com/gaaavr/gophers-of-war/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()
	g.AddPlayer()
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Gophers of war")
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
