package reports

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSwitchBinaryV1NoData(t *testing.T) {
	w, err := NewSwitchBinaryV1([]byte{})

	assert.Error(t, err)
	assert.Nil(t, w)
}

func TestSwitchBinaryV1False(t *testing.T) {
	w, err := NewSwitchBinaryV1([]byte{0x00})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchBinaryV1{})
	assert.Equal(t, w.CurrentValue, false)
	assert.Equal(t, w.String(), "current:false")
}

func TestSwitchBinaryV1True(t *testing.T) {
	w, err := NewSwitchBinaryV1([]byte{0xFF})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchBinaryV1{})
	assert.Equal(t, w.CurrentValue, true)
	assert.Equal(t, w.String(), "current:true")
}

func TestSwitchBinaryV2NoData(t *testing.T) {
	w, err := NewSwitchBinaryV2([]byte{})

	assert.Error(t, err)
	assert.Nil(t, w)
}

func TestSwitchBinaryV2False(t *testing.T) {
	w, err := NewSwitchBinaryV2([]byte{0x00, 0xFF, 0x00})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchBinaryV2{})
	assert.Equal(t, w.CurrentValue, false)
	assert.Equal(t, w.TargetValue, true)
	assert.Equal(t, w.String(), "current:false target:true duration:0")

	d, err := w.Duration.Duration()
	assert.NoError(t, err)
	assert.Equal(t, d, time.Duration(0))
}

func TestSwitchBinaryV2True(t *testing.T) {
	w, err := NewSwitchBinaryV2([]byte{0xFF, 0x00, 0x01})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchBinaryV2{})
	assert.Equal(t, w.CurrentValue, true)
	assert.Equal(t, w.TargetValue, false)
	assert.Equal(t, w.String(), "current:true target:false duration:1s")

	d, err := w.Duration.Duration()
	assert.NoError(t, err)
	assert.Equal(t, d, time.Second)
}
