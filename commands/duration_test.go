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
	assert.Equal(t, time.Duration(0), d)
	assert.Equal(t, time.Duration(0).String(), a.String())
}

func TestDuration1s(t *testing.T) {
	a := Duration(0x01)

	d, err := a.Duration()

	assert.NoError(t, err)
	assert.Equal(t, 1*time.Second, d)
	assert.Equal(t, time.Duration(1*time.Second).String(), a.String())
}
func TestDuration127s(t *testing.T) {
	a := Duration(0x7F)

	d, err := a.Duration()

	assert.NoError(t, err)
	assert.Equal(t, 127*time.Second, d)
	assert.Equal(t, time.Duration(127*time.Second).String(), a.String())
}

func TestDuration1m(t *testing.T) {
	a := Duration(0x80)

	d, err := a.Duration()

	assert.NoError(t, err)
	assert.Equal(t, 1*time.Minute, d)
	assert.Equal(t, time.Duration(1*time.Minute).String(), a.String())
}

func TestDuration126m(t *testing.T) {
	a := Duration(0xFD)

	d, err := a.Duration()

	assert.NoError(t, err)
	assert.Equal(t, 126*time.Minute, d)
	assert.Equal(t, time.Duration(126*time.Minute).String(), a.String())
}

func TestDurationUnknown(t *testing.T) {
	a := Duration(0xFE)

	_, err := a.Duration()

	assert.Error(t, err)
	assert.Equal(t, "Factory default duration", a.String())
}
