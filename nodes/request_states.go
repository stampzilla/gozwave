package nodes

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/stampzilla/gozwave/commands"
)

func (n *Node) RequestStates() error {
	n.RLock()
	cc := n.Device.CommandClasses
	n.RUnlock()

	for _, v := range cc {
		switch v.ID {
		case commands.SwitchBinary:
			cmd := commands.NewRaw([]byte{
				commands.SwitchBinary, // Command class
				0x02, // Command: GET
				0x25, // TransmitOptions?
				//0x23, // Callback?
			})
			cmd.SetNode(n.Id)
			n.connection.Write(cmd)
		case commands.SwitchMultilevel:
			cmd := commands.NewRaw([]byte{
				commands.SwitchMultilevel, // Command class
				0x02, // Command: GET
				0x25, // TransmitOptions?
				//0x23, // Callback?
			})
			cmd.SetNode(n.Id)
			n.connection.Write(cmd)
		case commands.SensorMultiLevel:
			cmd := commands.NewRaw([]byte{
				commands.SensorMultiLevel, // Command class
				0x01, // Command: SupportedGet
				0x25, // TransmitOptions?
				//0x23, // Callback?
			})
			cmd.SetNode(n.Id)
			t, _ := n.connection.WriteAndWaitForReport(cmd, time.Second*2, 0x02) // Wait for SupportedReport (0x02)

			report := <-t
			//for k, v := range report.sensors {
			// Request value for each sensor endpoint
			//}

			//logrus.Errorf("Sensors supportedget:")
			spew.Dump(report)
		}
	}

	n.Lock()
	n.statesOk = true
	n.Unlock()
	return nil
}
