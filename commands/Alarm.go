package commands

import "fmt"

type AlarmReport struct {
	*report
	Type           byte
	Level          byte
	SensorSourceId byte
	SensorType     byte
	Event          byte
	Status         byte

	data []byte
}

func NewAlarmReport(data []byte) *AlarmReport {
	alarm := &AlarmReport{data: data}

	if len(data) >= 2 {
		alarm.Type = data[0]
		alarm.Level = data[1]
	}

	if len(data) >= 6 {
		alarm.SensorSourceId = data[2]
		alarm.Status = data[3]
		alarm.SensorType = data[4]
		alarm.Event = data[5]
	}

	// Version 2
	if len(data) >= 9 {

	}
	return alarm
}

func (alarm *AlarmReport) String() string {
	if len(alarm.data) >= 6 {
		return fmt.Sprintf("Alarm: type=%d, level=%d, sensorSrcID=%d, type:%s event:%d, status=%d",
			alarm.Type, alarm.Level, alarm.SensorSourceId, alarm.SensorType, alarm.Event, alarm.Status)
	}
	return fmt.Sprintf("Alarm: type=%d level=%d", alarm.Type, alarm.Level)

}
