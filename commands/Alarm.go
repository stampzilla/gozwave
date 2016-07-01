package commands

import "fmt"

type CmdAlarm struct {
	Type           byte
	Level          byte
	SensorSourceId byte
	SensorType     byte
	Event          byte
	Status         byte

	data []byte
}

func NewCmdAlarm(data []byte) *CmdAlarm {
	alarm := &CmdAlarm{data: data}

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

func (self *CmdAlarm) String() string {
	if len(self.data) >= 6 {
		return fmt.Sprintf("Alarm: type=%d, level=%d, sensorSrcID=%d, type:%s event:%d, status=%d",
			self.Type, self.Level, self.SensorSourceId, self.SensorType, self.Event, self.Status)
	}
	return fmt.Sprintf("Alarm: type=%d level=%d", self.Type, self.Level)

}
