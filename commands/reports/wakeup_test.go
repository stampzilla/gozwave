package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWakeUp(t *testing.T) {
	w, err := NewWakeUp()

	assert.NoError(t, err)
	assert.IsType(t, &WakeUp{}, w)
}
