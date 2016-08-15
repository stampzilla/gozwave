package serialapi

type Event interface {
}

func (self *Connection) GetNextEvent() chan Event {
	c := make(chan Event)

	return c
}
