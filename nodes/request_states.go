package nodes

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/stampzilla/gozwave/commands"
)

func (self *Node) RequestStates() error {
	self.RLock()
	cc := self.Device.CommandClasses
	self.RUnlock()

	for _, v := range cc {
		switch v.ID {
		case commands.SwitchBinary:
			cmd := commands.NewRaw([]byte{
				commands.SwitchBinary, // Command class
				0x02, // Command: GET
				0x25, // TransmitOptions?
				//0x23, // Callback?
			})
			cmd.SetNode(self.Id)
			self.connection.Write(cmd)
		case commands.SwitchMultilevel:
			cmd := commands.NewRaw([]byte{
				commands.SwitchMultilevel, // Command class
				0x02, // Command: GET
				0x25, // TransmitOptions?
				//0x23, // Callback?
			})
			cmd.SetNode(self.Id)
			self.connection.Write(cmd)
		case commands.SensorMultiLevel:
			cmd := commands.NewRaw([]byte{
				commands.SensorMultiLevel, // Command class
				0x01, // Command: SupportedGet
				0x25, // TransmitOptions?
				//0x23, // Callback?
			})
			cmd.SetNode(self.Id)
			t, _ := self.connection.WriteAndWaitForReport(cmd, time.Second*2, 0x02) // Wait for SupportedReport (0x02)

			report := <-t
			//for k, v := range report.sensors {
			// Request value for each sensor endpoint
			//}

			//logrus.Errorf("Sensors supportedget:")
			spew.Dump(report)
		}
	}

	self.Lock()
	self.statesOk = true
	self.Unlock()
	return nil
}
