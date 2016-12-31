package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlarmNoData(t *testing.T) {
	w, err := NewAlarm([]byte{})

	assert.Error(t, err)
	assert.IsType(t, w, &Alarm{})
}

func TestAlarmShort(t *testing.T) {
	w, err := NewAlarm([]byte{0x01, 0x02})

	assert.NoError(t, err)
	assert.IsType(t, w, &Alarm{})
	assert.Equal(t, w.Type, byte(1), w.Type)
	assert.Equal(t, w.Level, byte(2), w.Level)
	assert.Equal(t, w.String(), "alarm type:1 level:2")
}

func TestAlarmLong(t *testing.T) {
	w, err := NewAlarm([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06})

	assert.NoError(t, err)
	assert.IsType(t, w, &Alarm{})
	assert.Equal(t, w.Type, byte(1))
	assert.Equal(t, w.Level, byte(2))
	assert.Equal(t, w.String(), "alarm type:1 level:2 sensorSrcID:3 status:4 sensorType:5 event:6")
}
