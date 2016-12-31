package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManufacturerSpecificNoData(t *testing.T) {
	w, err := NewManufacturerSpecific([]byte{})

	assert.Error(t, err)
	assert.IsType(t, w, &ManufacturerSpecific{})
}

func TestManufacturerSpecific(t *testing.T) {
	w, err := NewManufacturerSpecific([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06})

	assert.NoError(t, err)
	assert.IsType(t, w, &ManufacturerSpecific{})
	assert.Equal(t, w.Manufacturer, "0102", w.Manufacturer)
	assert.Equal(t, w.Type, "0304", w.Type)
	assert.Equal(t, w.ID, "0506", w.ID)
	assert.Equal(t, "manufacturerSpecific manufacturer:0102 type:0304 id:0506 (0102:0304:0506)", w.String())
}
