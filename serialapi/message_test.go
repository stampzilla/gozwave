package serialapi

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	raw := []string{
		"01 11 00 04 00 02 0b 71 05 07 00 00 ff 07 00 01 08 00 00", // Alarm report
		"01 11 00 04 00 02 0b 71 05 07 ff 00 ff 07 08 01 08 00 00", // Alarm report
		"01 11 00 04 00 02 0b 71 05 07 00 00 ff 07 00 01 08 00 00", // Alarm report

		"01 0c 00 04 00 03 06 31 05 03 0a 00 c2 00", // Sensor multilevel
		"01 0c 00 04 00 03 06 31 05 03 0a 00 1d 00", // Sensor multilevel

		"01 08 00 04 04 03 02 84 07 00", // Wakeup

		"01 10 00 02 00 00 06 01 02 03 00 00 00", // DiscoveryNodes
		"01 10 00 41 00 00 00 00 00 00 ",         // GetNodeProtocolInfo
	}

	for _, val := range raw {
		data := []byte{}
		bytes := strings.Split(val, " ")
		for _, b := range bytes {
			dhex, _ := hex.DecodeString(b)
			data = append(data, dhex...)
		}

		msg, err := NewMessage(data)
		assert.NoError(t, err, "data: %s", val)

		if msg != nil && testing.Verbose() {
			fmt.Printf("%+v\n\n", msg.Data)
		}
	}
}

func TestDecode(t *testing.T) {
	raw := []string{
		"06", // Ack
		"15", // Nak
		"18", // Can

		"01 11 00 04 00 02 0b 71 05 07 00 00 ff 07 00 01 08 00 61", // Alarm report
		"01 11 00 04 00 02 0b 71 05 07 ff 00 ff 07 08 01 08 00 96", // Alarm report
		"01 11 00 04 00 02 0b 71 05 07 00 00 ff 07 00 01 08 00 61", // Alarm report

		"01 0c 00 04 00 03 06 31 05 03 0a 00 c2 0d", // Sensor multilevel
		"01 0c 00 04 00 03 06 31 05 03 0a 00 1d d2", // Sensor multilevel

		"01 08 00 04 04 03 02 84 07 75", // Wakeup
	}

	for _, val := range raw {
		data := []byte{}
		bytes := strings.Split(val, " ")
		for _, b := range bytes {
			dhex, _ := hex.DecodeString(b)
			data = append(data, dhex...)
		}

		l, msg, err := Decode(data)
		assert.NotNil(t, msg)
		assert.NotEqual(t, -1, l)
		assert.NoError(t, err)

		if msg != nil && testing.Verbose() {
			fmt.Printf("%+v\n\n", msg.Data)
		}
	}

	// Without length
	l, msg, err := Decode([]byte{0x01})
	assert.Equal(t, 2, l)
	assert.Nil(t, msg)
	assert.NoError(t, err)

	// To short
	l, msg, err = Decode([]byte{0x01, 0x02})
	assert.Equal(t, 4, l)
	assert.Nil(t, msg)
	assert.NoError(t, err)

	// Invalid checksum
	l, msg, err = Decode([]byte{0x01, 0x02, 0x00, 0x00})
	assert.Equal(t, -1, l)
	assert.Nil(t, msg)
	assert.Error(t, err)

	// Invalid function
	l, msg, err = Decode([]byte{0x01, 0x02, 0x00, 0xfd})
	assert.Equal(t, 4, l)
	assert.Nil(t, msg)
	assert.Error(t, err)

	// Invalid start char
	l, msg, err = Decode([]byte{0x99})
	assert.Equal(t, -1, l)
	assert.Nil(t, msg)
	assert.NoError(t, err)

}

func TestMessageIs(t *testing.T) {
	m := Message{}

	assert.Equal(t, false, m.IsACK())
	assert.Equal(t, false, m.IsNAK())
	assert.Equal(t, false, m.IsCAN())

	m.startByte = 0x06
	assert.Equal(t, true, m.IsACK())
	assert.Equal(t, false, m.IsNAK())
	assert.Equal(t, false, m.IsCAN())

	m.startByte = 0x15
	assert.Equal(t, false, m.IsACK())
	assert.Equal(t, true, m.IsNAK())
	assert.Equal(t, false, m.IsCAN())

	m.startByte = 0x18
	assert.Equal(t, false, m.IsACK())
	assert.Equal(t, false, m.IsNAK())
	assert.Equal(t, true, m.IsCAN())
}

func TestCompileMessage(t *testing.T) {
	m := CompileMessage([]byte{0x01, 0x02})

	assert.Equal(t, []byte{0x1, 0x4, 0x0, 0x1, 0x2, 0xf8}, m)
}
