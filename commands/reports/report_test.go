package reports

import (
	"testing"

	"github.com/stampzilla/gozwave/commands"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type row struct {
		data     []byte
		expected interface{}
	}

	m := map[commands.ZWaveCommand]map[byte][]row{
		commands.Alarm: map[byte][]row{
			0x05: []row{
				row{[]byte{0x00, 0x00}, &Alarm{}},
			},
		},
		commands.ManufacturerSpecific: map[byte][]row{
			0x05: []row{
				row{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, &ManufacturerSpecific{}},
			},
		},
		commands.MultiInstance: map[byte][]row{
			0x08: []row{
				row{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, &MultiChannelEndpoints{}},
			},
		},
		commands.SensorMultiLevel: map[byte][]row{
			0x05: []row{
				row{[]byte{0x00, 0x00}, &SensorMultiLevel{}},
			},
		},
		commands.SwitchBinary: map[byte][]row{
			0x03: []row{
				row{[]byte{0x00}, &SwitchBinaryV1{}},
				row{[]byte{0x00, 0x00, 0x00}, &SwitchBinaryV2{}},
			},
		},
		commands.SwitchMultilevel: map[byte][]row{
			0x03: []row{
				row{[]byte{0x00}, &SwitchMultilevelV1{}},
				row{[]byte{0x00, 0x00, 0x00}, &SwitchMultilevelV4{}},
			},
		},
		commands.WakeUp: map[byte][]row{
			0x00: []row{
				row{[]byte{}, &WakeUp{}},
			},
		},
	}
	for command, v := range m {
		for class, row := range v {
			for _, e := range row {
				r, err := New(command, class, []byte(e.data))

				assert.NoError(t, err)
				assert.IsType(t, e.expected, r)
			}
		}
	}
}

func TestNewNoData(t *testing.T) {
	r, err := New(commands.ZWaveCommand(0), byte(0), []byte{})

	assert.Error(t, err)
	assert.Nil(t, r)
}

func TestReport(t *testing.T) {
	r, err := New(commands.WakeUp, byte(0), []byte{})

	r.SetNode(5)

	assert.NoError(t, err)
	assert.NotNil(t, r)
}
