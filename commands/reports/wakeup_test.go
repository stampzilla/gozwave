package reports

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	w, err := NewWakeUp()

	assert.NoError(t, err)
	assert.IsType(t, &WakeUp{}, w)
	assert.Implements(t, (*fmt.Stringer)(nil), w)
	assert.Equal(t, w.String(), "wakeup")
}
