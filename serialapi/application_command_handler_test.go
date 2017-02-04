package serialapi

import (
	"testing"

	"github.com/stampzilla/gozwave/commands"
	"github.com/stretchr/testify/assert"
)

func TestNewApplicationCommandHandler(t *testing.T) {
	m, err := NewApplicationCommandHandler([]byte{0x00, 0x00, 0x00})

	assert.Nil(t, m)
	assert.Error(t, err)

	m, err = NewApplicationCommandHandler([]byte{commands.WakeUp, 0x00, 0x00})

	assert.NotNil(t, m)
	assert.NoError(t, err)
}
