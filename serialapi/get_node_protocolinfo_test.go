package serialapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGetNodeProtocolInfo(t *testing.T) {
	m, err := NewGetNodeProtocolInfo([]byte{0x01, 0x00, 0xff, 0x04, 0x05, 0x06})
	assert.NoError(t, err)
	assert.Equal(t, false, m.Listening, "Listening")
	assert.Equal(t, false, m.Routing, "Routing")
	assert.Equal(t, false, m.Beaming, "Beaming")
	assert.Equal(t, byte(0x02), m.Version, "Version")
	assert.Equal(t, false, m.Flirs, "Flirs")
	assert.Equal(t, false, m.Security, "Security")
	assert.Equal(t, 9600, m.MaxBaud, "MaxBaud")
	assert.Equal(t, byte(0x04), m.Basic, "Basic")
	assert.Equal(t, byte(0x05), m.Generic, "Generic")
	assert.Equal(t, byte(0x06), m.Specific, "Specific")

	m, err = NewGetNodeProtocolInfo([]byte{0x80, 0x01, 0xff, 0xff, 0xff, 0xff})
	assert.NoError(t, err)
	assert.Equal(t, true, m.Listening, "Listening")
	assert.Equal(t, false, m.Routing, "Routing")
	assert.Equal(t, false, m.Beaming, "Beaming")
	assert.Equal(t, false, m.Flirs, "Flirs")
	assert.Equal(t, true, m.Security, "Security")
	assert.Equal(t, 9600, m.MaxBaud, "MaxBaud")

	m, err = NewGetNodeProtocolInfo([]byte{0x40, 0x10, 0xff, 0xff, 0xff, 0xff})
	assert.NoError(t, err)
	assert.Equal(t, false, m.Listening, "Listening")
	assert.Equal(t, true, m.Routing, "Routing")
	assert.Equal(t, true, m.Beaming, "Beaming")
	assert.Equal(t, false, m.Flirs, "Flirs")
	assert.Equal(t, false, m.Security, "Security")
	assert.Equal(t, 9600, m.MaxBaud, "MaxBaud")

	m, err = NewGetNodeProtocolInfo([]byte{0x10, 0x60, 0xff, 0xff, 0xff, 0xff})
	assert.NoError(t, err)
	assert.Equal(t, false, m.Listening, "Listening")
	assert.Equal(t, false, m.Routing, "Routing")
	assert.Equal(t, false, m.Beaming, "Beaming")
	assert.Equal(t, true, m.Flirs, "Flirs")
	assert.Equal(t, false, m.Security, "Security")
	assert.Equal(t, 40000, m.MaxBaud, "MaxBaud")

	m, err = NewGetNodeProtocolInfo([]byte{0x01, 0x02, 0x03, 0x04, 0x05})

	assert.Error(t, err)
	assert.Nil(t, m)
}
