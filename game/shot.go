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
	s.Y = y
}

func (s Shots) resolveShots() Shots {
	newShots := make(Shots, 0, len(s))
	for idx, shot := range s {
		if shot.X <= 10 || shot.X >= 300 {
			continue
		}
		switch shot.Direction {
		case DirectionLeft:
			s[idx].X--
		case DirectionRight:
			s[idx].X++
		}
		newShots = append(newShots, s[idx])
	}
	return newShots

}
