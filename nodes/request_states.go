package nodes

import (
	"time"

	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/functions"
)

func (self *Node) RequestStates() error {

	for _, v := range self.Device.CommandClasses {
		switch v.ID {
		case commands.SwitchBinary:
			<-self.connection.SendRaw([]byte{
				functions.SendData, // Function
				byte(self.Id),      // Node id
				0x02,               // Length
				commands.SwitchBinary, // Command class
				0x02, // Command: GET
				0x25, // TransmitOptions?
				//0x23, // Callback?
			}, time.Second*2)
		case commands.SwitchMultilevel:
			<-self.connection.SendRaw([]byte{
				functions.SendData, // Function
				byte(self.Id),      // Node id
				0x02,               // Length
				commands.SwitchMultilevel, // Command class
				0x02, // Command: GET
				0x25, // TransmitOptions?
				//0x23, // Callback?
			}, time.Second*2)
		case commands.SensorMultiLevel:
			//<-time.After(time.Second * 10)
			<-self.connection.SendRawAndWaitForResponse([]byte{
				functions.SendData, // Function
				byte(self.Id),      // Node id
				0x02,               // Length
				commands.SensorMultiLevel, // Command class
				0x01, // Command: SupportedGet
				0x25, // TransmitOptions?
				//0x23, // Callback?
			}, time.Second*2, 0x02) // Wait for SupportedReport (0x02)

			//logrus.Errorf("Sensors supportedget:")
			//spew.Dump(resp)
		}
	}

	self.statesOk = true
	return nil
}
