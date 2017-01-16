package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchMultilevel0(t *testing.T) {
	r := NewSwitchMultilevel()
	r.SetValue(0)
	r.SetNode(99)

	assert.Equal(t, []byte{0x13, 0x63, 0x03, 0x26, 0x01, 0x00, 0x25}, r.Encode())
}

func TestSwitchMultilevel1(t *testing.T) {
	r := NewSwitchMultilevel()
	r.SetValue(1)
	r.SetNode(99)

	assert.Equal(t, []byte{0x13, 0x63, 0x03, 0x26, 0x01, 0x01, 0x25}, r.Encode())
}

func TestSwitchMultilevel99(t *testing.T) {
	r := NewSwitchMultilevel()
	r.SetValue(99)
	r.SetNode(99)

	assert.Equal(t, []byte{0x13, 0x63, 0x03, 0x26, 0x01, 0x63, 0x25}, r.Encode())
}

func TestSwitchMultilevel100(t *testing.T) {
	r := NewSwitchMultilevel()
	r.SetValue(100)
	r.SetNode(99)

	assert.Equal(t, []byte{0x13, 0x63, 0x03, 0x26, 0x01, 0x63, 0x25}, r.Encode())
}

func TestSwitchMultilevel255(t *testing.T) {
	r := NewSwitchMultilevel()
	r.SetValue(255)
	r.SetNode(99)

	assert.Equal(t, []byte{0x13, 0x63, 0x03, 0x26, 0x01, 0x63, 0x25}, r.Encode())
}
