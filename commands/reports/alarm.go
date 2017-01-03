package reports

import "fmt"

//Alarm reports are sent from a node as an event. Is typicaly used for motion sensors and such to notify when motion is detected.
type Alarm struct {
	*report
	Type           byte `json:"type_"`
	Level          byte `json:"level"`
	SensorSourceID byte `json:"sensor_source_id"`
	SensorType     byte `json:"sensor_type"`
	Event          byte `json:"event"`
	Status         byte `json:"status"`

	data []byte
}

// NewAlarm creates a new alarm report.
func NewAlarm(data []byte) (*Alarm, error) {
	a := &Alarm{data: data}

	if len(data) < 2 {
		return nil, fmt.Errorf("To short, expected at least 2 byte")
	}

	a.Type = data[0]
	a.Level = data[1]

	if len(data) >= 6 {
		a.SensorSourceID = data[2]
		a.Status = data[3]
		a.SensorType = data[4]
		a.Event = data[5]
	}

	// Version 2
	if len(data) >= 9 {

	}
	return a, nil
}

func (a *Alarm) String() string {
	if len(a.data) >= 6 {
		return fmt.Sprintf("alarm type:%x level:%x sensorSrcID:%x status:%x sensorType:%x event:%x",
			a.Type, a.Level, a.SensorSourceID, a.Status, a.SensorType, a.Event)
	}
	return fmt.Sprintf("alarm type:%x level:%x", a.Type, a.Level)

}
