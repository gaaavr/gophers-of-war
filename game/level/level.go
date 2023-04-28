package level

import (
	"fmt"
	"math/rand"
	"time"
)

// Level represents a Game level.
type Level struct {
	W, H int

	Tiles    [][]*Tile // (Y,X) array of Tiles
	TileSize int
}

// Tile returns the tile at the provided coordinates, or nil.
func (l *Level) Tile(x, y int) *Tile {
	if x >= 0 && y >= 0 && x < l.W && y < l.H {
		return l.Tiles[y][x]
	}
	return nil
}

// Size returns the size of the Level.
func (l *Level) Size() (width, height int) {
	return l.W, l.H
}

// NewLevel returns a new randomly generated Level.
func NewLevel() (*Level, error) {
	// Create a 108x108 Level.
	l := &Level{
		W:        108,
		H:        108,
		TileSize: 64,
	}

	// Load embedded SpriteSheet.
	ss, err := LoadSpriteSheet(l.TileSize)
	if err != nil {
		return nil, fmt.Errorf("failed to load embedded spritesheet: %s", err)
	}

	// Generate a unique permutation each time.
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	// Fill each tile with one or more sprites randomly.
	l.Tiles = make([][]*Tile, l.H)
	for y := 0; y < l.H; y++ {
		l.Tiles[y] = make([]*Tile, l.W)
		for x := 0; x < l.W; x++ {
			t := &Tile{}
			isBorderSpace := x == 0 || y == 0 || x == l.W-1 || y == l.H-1
			val := r.Intn(1000)
			switch {
			case isBorderSpace || val < 275:
				t.AddSprite(ss.Wall)
			case val < 285:
				t.AddSprite(ss.Statue)
			case val < 288:
				t.AddSprite(ss.Crown)
			case val < 289:
				t.AddSprite(ss.Floor)
				t.AddSprite(ss.Tube)
			case val < 290:
				t.AddSprite(ss.Portal)
			default:
				t.AddSprite(ss.Floor)
			}
			l.Tiles[y][x] = t
		}
	}

	return l, nil
}
