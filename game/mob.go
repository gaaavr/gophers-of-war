package game

type mob struct {
	spriteName string
	X          float64
	Y          float64
	isDead     bool
}

type mobs []mob

func (world *World) resolveMobs(playerX, playerY float64) {
	newMobs := make(mobs, 0, len(world.mobs))
	for _, mob := range world.mobs {
		if mob.isDead {
			continue
		}

		if mob.X-playerX == 0 && mob.Y-playerY == 0 {
			mob.isDead = true
		}

		if mob.X-playerX > 0 {
			mob.X -= 0.5
		} else {
			mob.X += 0.5
		}

		if mob.Y-playerY > 0 {
			mob.Y -= 0.5
		} else {
			mob.Y += 0.5
		}

		newMobs = append(newMobs, mob)
	}
	world.mobs = newMobs
}

func (m mob) isKilled(shotX, shotY float64) bool {
	if (shotX <= m.X+12 && shotX >= m.X-12) && (shotY <= m.Y+12 && shotY >= m.Y-12) {
		return true
	}
	return false
}
