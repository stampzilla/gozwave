package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchBinaryNoData(t *testing.T) {
	w, err := NewSwitchBinary([]byte{})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchBinary{})
	assert.Equal(t, w.Status, false)
	assert.Equal(t, w.String(), "status: false")
}

func TestSwitchBinaryFalse(t *testing.T) {
	w, err := NewSwitchBinary([]byte{0x00})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchBinary{})
	assert.Equal(t, w.Status, false)
	assert.Equal(t, w.String(), "status: false")
}

func TestSwitchBinaryTrue(t *testing.T) {
	w, err := NewSwitchBinary([]byte{0xFF})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchBinary{})
	assert.Equal(t, w.Status, true)
	assert.Equal(t, w.String(), "status: true")
}
