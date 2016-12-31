package reports

import (
	"testing"

	"github.com/stampzilla/gozwave/commands"
	"github.com/stretchr/testify/assert"
)

type row struct {
	data     []byte
	expected interface{}
}

func TestNewWakeUp(t *testing.T) {
	m := map[commands.ZWaveCommand]map[byte]row{
		commands.Alarm: map[byte]row{
			0x05: row{[]byte{0x00, 0x00}, &Alarm{}},
		},
		commands.ManufacturerSpecific: map[byte]row{
			0x05: row{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, &ManufacturerSpecific{}},
		},
		commands.MultiInstance: map[byte]row{
			0x08: row{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, &MultiChannelEndpoints{}},
		},
		commands.SensorMultiLevel: map[byte]row{
			0x05: row{[]byte{0x00, 0x00}, &SensorMultiLevel{}},
		},
		commands.SwitchBinary: map[byte]row{
			0x03: row{[]byte{0x00, 0x00}, &SwitchBinary{}},
		},
		commands.SwitchMultilevel: map[byte]row{
			0x03: row{[]byte{0x00, 0x00}, &SwitchMultilevel{}},
		},
		commands.WakeUp: map[byte]row{
			0x00: row{[]byte{}, &WakeUp{}},
		},
	}
	for command, v := range m {
		for class, e := range v {
			r, err := New(command, class, []byte(e.data))

			assert.NoError(t, err)
			assert.IsType(t, e.expected, r)
		}
	}
}
