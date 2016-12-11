package gozwave

// Controllable interface used by nodes and endpoints to controll them
type Controllable interface {
	On()
	Off()
	Level(float64)
}
