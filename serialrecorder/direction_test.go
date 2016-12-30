package serialrecorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectionRead(t *testing.T) {
	var direction = DirectionRead
	assert.Equal(t, true, direction.IsRead())
	assert.Equal(t, false, bool(direction))
	assert.Equal(t, "read", direction.String())
}
func TestDirectionWrite(t *testing.T) {
	var direction = DirectionWrite
	assert.Equal(t, true, direction.IsWrite())
	assert.Equal(t, true, bool(direction))
	assert.Equal(t, "write", direction.String())
}
