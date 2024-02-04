package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	framesCount          int = 4
	backgroundTilesCount int = 6
)

type Assets struct {
	CharacterSprites [][][]*ebiten.Image
	TileSprites      []*ebiten.Image
}

func LoadAssets() (*Assets, error) {
	a := &Assets{}

	characterSprites, err := loadCharactersSprites()
	if err != nil {
		return nil, err
	}

	a.CharacterSprites = characterSprites

	// tile map
	a.TileSprites = make([]*ebiten.Image, backgroundTilesCount)
	for i := 0; i < backgroundTilesCount; i++ {
		s, _, err := ebitenutil.NewImageFromFile(
			fmt.Sprintf("assets/images/background/%d.png", i),
		)
		if err != nil {
			return nil, err
		}
		a.TileSprites[i] = s
	}

	return a, nil
}

func loadCharactersSprites() ([][][]*ebiten.Image, error) {
	charactersSprites := []string{"gopher"}
	animationTypes := []string{"idle", "run"}

	allSprites := make([][][]*ebiten.Image, 0, len(charactersSprites))

	for _, char := range charactersSprites {
		charAnimations := make([][]*ebiten.Image, 0, len(animationTypes))

		for _, animationType := range animationTypes {
			charFrames := make([]*ebiten.Image, 0, framesCount)

			for i := 0; i < framesCount; i++ {
				s, _, err := ebitenutil.NewImageFromFile(
					fmt.Sprintf("assets/images/heroes/%s/%s/%d.png", char, animationType, i),
				)
				if err != nil {
					return nil, err
				}

				charFrames = append(charFrames, s)
			}

			charAnimations = append(charAnimations, charFrames)
		}

		allSprites = append(allSprites, charAnimations)
	}

	return allSprites, nil
}
