package gozwave

import (
	"testing"

	"github.com/stampzilla/gozwave/serialapi"
	"github.com/stretchr/testify/assert"
)

func TestSendPackageMatch(t *testing.T) {
	s := newSendPackage([]byte{serialapi.SendData, 0x02, 0x03, 0x04, 0x05})

	assert.Equal(t, byte(serialapi.SendData), s.function, "Function")
	assert.Equal(t, byte(0x02), s.node, "Node")
	assert.Equal(t, byte(0x04), s.commandclass, "CommandClass")
	assert.Equal(t, byte(0x00), s.expectedReport, "ExpectedReport")

	assert.Equal(t, false, s.Match([]byte{0x00, 0x00, 0x00, 0xff, 0x00, 0x02, 0x04, 0x05, 0x06, 0x00}), "Function")
	assert.Equal(t, false, s.Match([]byte{0x00, 0x00, 0x00, byte(serialapi.ApplicationCommandHandler), 0x00, 0xff, 0x00, 0x04, 0x06, 0x00}), "Node")
	assert.Equal(t, false, s.Match([]byte{0x00, 0x00, 0x00, byte(serialapi.ApplicationCommandHandler), 0x00, 0x02, 0x00, 0xff, 0x06, 0x00}), "Command")
	assert.Equal(t, true, s.Match([]byte{0x00, 0x00, 0x00, byte(serialapi.ApplicationCommandHandler), 0x00, 0x02, 0x00, 0x04, 0x06}))

	s.expectedReport = byte(0x06)
	assert.Equal(t, true, s.Match([]byte{0x00, 0x00, 0x00, byte(serialapi.ApplicationCommandHandler), 0x00, 0x02, 0x00, 0x04, 0x06}))
	assert.Equal(t, false, s.Match([]byte{0x00, 0x00, 0x00, byte(serialapi.ApplicationCommandHandler), 0x00, 0x02, 0x00, 0x04, 0xff}), "Expected report")
}
