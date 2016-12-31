package reports

import "fmt"

type WakeUp struct {
	*report
}

// NewWakeUp creates a new wakeUp report. Zwave nodes sends a wakeUp report when they wake up and listens for messages from the controller. This usualy happens something like once every hour and the node is typicaly awake for 1 second.
func NewWakeUp() (*WakeUp, error) {
	return &WakeUp{}, nil
}

func (wr *WakeUp) String() string {
	return fmt.Sprintf("WakeUp")
}
