package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiChannelEncap(t *testing.T) {
	r := NewMultiChannelEncap([]byte{0x01, 0x02, 0x03,0x04}, 99)

	e := r.Encode() 

	assert.Equal(t, []byte{0x01, 0x02, 0x07, 0x60, 0x0d, 0x01, 0x63, 0x04}, e)
}
