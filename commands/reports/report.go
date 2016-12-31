package reports

import (
	"fmt"

	"github.com/stampzilla/gozwave/commands"
)

type Report interface {
	SetNode(byte)
}

type report struct {
	node byte
}

func (r *report) SetNode(n byte) {
	if r == nil {
		r = &report{}
		//r = nr
	}

	r.node = n
}

// New creates a new report based on command, command class and data
func New(c commands.ZWaveCommand, class byte, data []byte) (Report, error) {
	switch c {
	case commands.Alarm:
		switch class {
		case 0x05: // Report
			return NewAlarm(data)
		}
	case commands.ManufacturerSpecific:
		switch class {
		case 0x05: // Report
			return NewManufacturerSpecific(data)
		}
	case commands.MultiInstance:
		switch class {
		case 0x08: // MultiChannelCmd_EndPointReport
			return NewMultiChannelEndpoints(data)
		}
	case commands.SensorMultiLevel:
		switch class {
		case 0x05: // Report
			return NewSensorMultiLevel(data)
		}
	case commands.SwitchBinary:
		switch class {
		case 0x03: // Report
			return NewSwitchBinary(data)
		}
	case commands.SwitchMultilevel:
		switch class {
		case 0x03: // Report
			return NewSwitchMultilevel(data)
		}
	case commands.WakeUp:
		return NewWakeUp()
	}

	return nil, fmt.Errorf("Unknown command (%x - %s) / command class (%x)", byte(c), c, class)
}
