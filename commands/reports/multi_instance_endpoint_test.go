package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiChannelEndpointsNoData(t *testing.T) {
	w, err := NewMultiChannelEndpoints([]byte{})

	assert.Error(t, err)
	assert.Nil(t, w)
}

func TestMultiChannelEndpoints(t *testing.T) {
	w, err := NewMultiChannelEndpoints([]byte{0x40, 0x03})

	assert.NoError(t, err)
	assert.IsType(t, w, &MultiChannelEndpoints{})
	assert.Equal(t, w.Dynamic, false)
	assert.Equal(t, w.Identical, true)
	assert.Equal(t, w.Endpoints, 3)
	assert.Equal(t, w.String(), "multiChannelEndpoints dynamic:false identical:true endpoints:3")
}
