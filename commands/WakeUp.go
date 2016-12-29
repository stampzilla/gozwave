package commands

import "fmt"

type WakeUpReport struct {
	*report
}

func NewWakeUpReport() *WakeUpReport {
	return &WakeUpReport{}
}

func (wr *WakeUpReport) String() string {
	return fmt.Sprintf("WakeUp")
}
