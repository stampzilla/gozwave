package reports

import (
	"fmt"

	"github.com/stampzilla/gozwave/commands"
)

type Report interface {
	SetNode(byte)
	String() string
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
			switch len(data) {
			case 1:
				return NewSwitchBinaryV1(data)
			case 3:
				return NewSwitchBinaryV2(data)
			}
		}
	case commands.SwitchMultilevel:
		switch class {
		case 0x03: // Report
			switch len(data) {
			case 1:
				return NewSwitchMultilevelV1(data)
			// V2 reports is same as V1
			// V3 reports is depreciated and is not recomended to implement by sigma designs
			case 3:
				return NewSwitchMultilevelV4(data)
			}
		}
	case commands.WakeUp:
		return NewWakeUp()
	}

	return nil, fmt.Errorf("Unknown command (%x - %s) / command class (%x)", byte(c), c, class)
}
