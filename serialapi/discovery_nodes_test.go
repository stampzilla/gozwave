package serialapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDiscoverdNodes(t *testing.T) {
	m, err := NewDiscoverdNodes([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06})

	assert.NoError(t, err)
	assert.Equal(t, [232]bool{true, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, true, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}, m.ActiveNodes)

	m, err = NewDiscoverdNodes([]byte{0x01, 0x02, 0x03, 0x04, 0x05})

	assert.Error(t, err)
	assert.Nil(t, m)
}
