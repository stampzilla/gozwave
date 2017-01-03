package commands

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration0(t *testing.T) {
	a := Duration(0)

	d, err := a.Duration()

	assert.NoError(t, err)
	assert.Equal(t, d, time.Duration(0))
}

func TestDuration1s(t *testing.T) {
	a := Duration(0x01)

	d, err := a.Duration()

	assert.NoError(t, err)
	assert.Equal(t, d, 1*time.Second)
}
func TestDuration127s(t *testing.T) {
	a := Duration(0x7F)

	d, err := a.Duration()

	assert.NoError(t, err)
	assert.Equal(t, d, 127*time.Second)
}

func TestDuration1m(t *testing.T) {
	a := Duration(0x80)

	d, err := a.Duration()

	assert.NoError(t, err)
	assert.Equal(t, d, 1*time.Minute)
}

func TestDuration126m(t *testing.T) {
	a := Duration(0xFD)

	d, err := a.Duration()

	assert.NoError(t, err)
	assert.Equal(t, d, 126*time.Minute)
}

func TestDurationUnknown(t *testing.T) {
	a := Duration(0xFE)

	_, err := a.Duration()

	assert.Error(t, err)
}
