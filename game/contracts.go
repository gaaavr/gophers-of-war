package game

type Handler interface {
	handleEvent(world *World)
}
