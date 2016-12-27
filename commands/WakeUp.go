package commands

import "fmt"

type WakeUpReport struct {
	*report
}

func NewWakeUpReport() *WakeUpReport {
	return &WakeUpReport{}
}

func (self *WakeUpReport) String() string {
	return fmt.Sprintf("WakeUp")
}
