package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchMultilevelNoData(t *testing.T) {
	w, err := NewSwitchMultilevel([]byte{})

	assert.NoError(t, err)
	assert.IsType(t, &SwitchMultilevel{}, w)
	assert.Equal(t, w.Level, byte(0))
}

func TestSwitchMultilevel0(t *testing.T) {
	w, err := NewSwitchMultilevel([]byte{0x00})

	assert.NoError(t, err)
	assert.IsType(t, &SwitchMultilevel{}, w)
	assert.Equal(t, w.Level, byte(0))
}

func TestSwitchMultilevel100(t *testing.T) {
	w, err := NewSwitchMultilevel([]byte{100})

	assert.NoError(t, err)
	assert.IsType(t, &SwitchMultilevel{}, w)
	assert.Equal(t, w.Level, byte(100))
}
