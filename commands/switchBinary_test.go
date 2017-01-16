package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchBinaryTrue(t *testing.T) {
	r := NewSwitchBinary()
	r.SetValue(true)
	r.SetNode(99)

	e := r.Encode() 

	assert.Equal(t, []byte{0x13, 0x63, 0x03, 0x25, 0x01, 0xff, 0x25}, e)
}

func TestSwitchBinaryFalse(t *testing.T) {
	r := NewSwitchBinary()
	r.SetValue(false)
	r.SetNode(99)

	e := r.Encode() 

	assert.Equal(t, []byte{0x13, 0x63, 0x03, 0x25, 0x01, 0x00, 0x25}, e)
}
