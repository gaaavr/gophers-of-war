package game

type handler interface {
	handleEvent(world *World)
}
