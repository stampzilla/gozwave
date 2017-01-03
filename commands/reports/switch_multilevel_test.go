package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchMultilevelNoData(t *testing.T) {
	w, err := NewSwitchMultilevelV1([]byte{})

	assert.Error(t, err)
	assert.Nil(t, w)
}

func TestSwitchMultilevel0(t *testing.T) {
	w, err := NewSwitchMultilevelV1([]byte{0x00})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchMultilevelV1{})
	assert.Equal(t, w.CurrentValue, 0.0)
	assert.Equal(t, w.Valid, true)
	assert.Equal(t, w.String(), "current:0.000000")
}

func TestSwitchMultilevel100(t *testing.T) {
	w, err := NewSwitchMultilevelV1([]byte{100})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchMultilevelV1{})
	assert.Equal(t, w.CurrentValue, 100.0)
	assert.Equal(t, w.Valid, true)
	assert.Equal(t, w.String(), "current:100.000000")
}

func TestSwitchMultilevelUnknown(t *testing.T) {
	w, err := NewSwitchMultilevelV1([]byte{0xFE})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchMultilevelV1{})
	assert.Equal(t, w.CurrentValue, 100.0)
	assert.Equal(t, w.Valid, false)
	assert.Equal(t, w.String(), "current:unknown")
}

func TestSwitchMultilevelV4NoData(t *testing.T) {
	w, err := NewSwitchMultilevelV4([]byte{})

	assert.Error(t, err)
	assert.Nil(t, w)
}

func TestSwitchMultilevelV40(t *testing.T) {
	w, err := NewSwitchMultilevelV4([]byte{0x00, 0xFF, 0x00})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchMultilevelV4{})
	assert.Equal(t, w.CurrentValue, 0.0)
	assert.Equal(t, w.Valid, true)
	assert.Equal(t, w.String(), "current:0.000000 target:100.000000 duration:0")
}

func TestSwitchMultilevelV4100(t *testing.T) {
	w, err := NewSwitchMultilevelV4([]byte{0xFF, 0x00, 0x01})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchMultilevelV4{})
	assert.Equal(t, w.CurrentValue, 100.0)
	assert.Equal(t, w.TargetValue, 0.0)
	assert.Equal(t, w.Valid, true)
	assert.Equal(t, w.TargetValid, true)
	assert.Equal(t, w.String(), "current:100.000000 target:0.000000 duration:1s")
}

func TestSwitchMultilevelV4Unknown(t *testing.T) {
	w, err := NewSwitchMultilevelV4([]byte{0xFE, 0x00, 0x00})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchMultilevelV4{})
	assert.Equal(t, w.CurrentValue, 100.0)
	assert.Equal(t, w.TargetValue, 0.0)
	assert.Equal(t, w.Valid, false)
	assert.Equal(t, w.TargetValid, true)
	assert.Equal(t, w.String(), "current:unknown target:0.000000 duration:0")
}

func TestSwitchMultilevelV4UnknownTarget(t *testing.T) {
	w, err := NewSwitchMultilevelV4([]byte{0x00, 0xFE, 0x00})

	assert.NoError(t, err)
	assert.IsType(t, w, &SwitchMultilevelV4{})
	assert.Equal(t, w.CurrentValue, 0.0)
	assert.Equal(t, w.TargetValue, 100.0)
	assert.Equal(t, w.Valid, true)
	assert.Equal(t, w.TargetValid, false)
	assert.Equal(t, w.String(), "current:0.000000 target:unknown duration:0")
}
