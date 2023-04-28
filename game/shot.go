package game

type Shot struct {
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Direction int
}

type Shots []Shot

func (s *Shot) getShotOpts(direction int, x, y float64) {
	if direction == DirectionLeft {
		s.X = x - 16
		s.Direction = DirectionLeft
	} else {
		s.X = x + 16
		s.Direction = DirectionRight
	}
	s.Y = y + 2
}

func (world *World) resolveShots() {
	newShots := make(Shots, 0, len(world.Shots))
	for _, shot := range world.Shots {
		var isShot bool
		if shot.X <= 10 || shot.X >= 300 {
			continue
		}
		for id, mobe := range world.mobs {
			if mobe.isDead {
				continue
			}
			if mobe.isKilled(shot.X, shot.Y) {
				world.mobs[id].isDead = true
				isShot = true
				break
			}
		}
		if isShot {
			continue
		}
		switch shot.Direction {
		case DirectionLeft:
			shot.X -= 1
		case DirectionRight:
			shot.X += 1
		}
		newShots = append(newShots, shot)
	}
	world.Shots = newShots

}
