package game

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TileSize = 16 * 3
)

type Tile struct {
	rect       *Rect
	tileType   int
	tileSprite *ebiten.Image
	isObstacle bool
}

type Rect struct {
	x, y          float64
	width, height float64
}

type Level struct {
	assets *Assets
	tiles  [][]*Tile
}

func (t *Tile) XY() (float64, float64) {
	return t.rect.XY()
}

func (r *Rect) XY() (float64, float64) {
	return r.x, r.y
}

func NewLevel(name string) (*Level, error) {
	assets, err := LoadAssets()
	if err != nil {
		return nil, fmt.Errorf("failed to load assets: %w", err)
	}

	levelData, err := loadLevelData(name)
	if err != nil {
		return nil, fmt.Errorf("failed to load Level data: %w", err)
	}

	if len(levelData) == 0 {
		return nil, errors.New("failed to load Level: Level data is empty")
	}

	tiles := make([][]*Tile, len(levelData))
	for y := range tiles {
		tiles[y] = make([]*Tile, len(levelData[0]))
	}

	l := &Level{
		assets: assets,
		tiles:  tiles,
	}

	for y, row := range levelData {
		for x, col := range row {
			l.processTiles(x, y, col)
		}
	}

	return &Level{
		assets: assets,
		tiles:  tiles,
	}, nil
}

// load Level csv schema
func loadLevelData(levelName string) ([][]int, error) {
	file, err := os.Open(fmt.Sprintf("levels/%s.csv", levelName))
	if err != nil {
		return nil, fmt.Errorf("failed to open csv file: %w", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	result := make([][]int, len(data))

	for y, row := range data {
		result[y] = make([]int, 0, len(row))

		for _, col := range row {
			tileType, err := strconv.Atoi(col)
			if err != nil {
				return nil, fmt.Errorf("failed to convert Level schema value: %w", err)
			}

			result[y] = append(result[y], tileType)
		}
	}

	return result, nil
}

func (l *Level) processTiles(x, y, tileType int) {
	// empty tile
	if tileType == -1 {
		l.tiles[y][x] = &Tile{tileType: -1}
		return
	}

	// not empty tile
	image := l.assets.TileSprites[tileType]
	tile := &Tile{
		tileType:   tileType,
		tileSprite: image,
		isObstacle: false,
		rect: &Rect{
			x:      float64(x * TileSize),
			y:      float64(y * TileSize),
			width:  TileSize,
			height: TileSize,
		},
	}

	l.tiles[y][x] = tile
}
