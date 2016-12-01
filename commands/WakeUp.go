package commands

import "fmt"

type CmdWakeUp struct {
	data []byte
}

func NewWakeUp(data []byte) *CmdWakeUp {
	return &CmdWakeUp{data: data}
}

func (self *CmdWakeUp) String() string {
	return fmt.Sprintf("WakeUp data: %x", self.data)
}
