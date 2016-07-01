package functions

import (
	"fmt"

	"github.com/stampzilla/gozwave/commands"
)

type FuncApplicationCommandHandler struct {
	Command commands.ZWaveCommand
	Class   byte // Command Class

	Data interface{}
}

func NewApplicationCommandHandler(data []byte) *FuncApplicationCommandHandler {
	f := &FuncApplicationCommandHandler{}
	f.Decode(data)

	return f
}

func (self *FuncApplicationCommandHandler) Decode(data []byte) {
	self.Command = commands.ZWaveCommand(data[0])
	self.Class = data[1]

	switch self.Command {
	case commands.Basic:
	case commands.Alarm:
		switch self.Class {
		case 0x05: // Report
			self.Data = commands.NewCmdAlarm(data[2:])
			fmt.Printf("%+v\n", self.Data)
		}
	case commands.SensorMultiLevel:
		switch self.Class {
		case 0x05: // Report
			self.Data = commands.NewCmdSensorMultiLevel(data[2:])
			fmt.Printf("%+v\n", self.Data)
		}
	case commands.WakeUp:
	}
}
