package serialapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRaw(t *testing.T) {
	r := NewRaw([]byte{0x01, 0x02, 0x03, 0x04})

	e := r.Encode()

	assert.Equal(t, []byte{0x01, 0x02, 0x03, 0x04}, e)
}
