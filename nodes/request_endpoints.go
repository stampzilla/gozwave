package nodes

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/commands"
)

func (n *Node) RequestEndpoints() error {
	// Todo: Send raw messages here
	n.RLock()
	if n.Device == nil || n.Device.CommandClasses == nil {
		n.RUnlock()
		n.Lock()
		n.Endpoints = make([]*Endpoint, 0)
		n.Unlock()
		return fmt.Errorf("Failed 'RequestEndpoints', no commandclasses exists")
	}

	//for _, v := range n.Device.CommandClasses {
	//logrus.Errorf("Request endpoint %d:%x - %s", n.Id, byte(v.ID), v.ID)
	//}
	n.RUnlock()

	cmd := commands.NewRaw(
		[]byte{
			commands.MultiInstance, // Command class
			0x07, // Command: MultiInstanceCmd_Get
			//byte(v.ID),
			0x25, // TransmitOptions?
			//0x23, // Callback?
		})
	cmd.SetNode(n.Id)

	fmt.Println("")
	fmt.Println("")
	logrus.Debugf("Request endpoint %d", n.Id)

	t, _ := n.connection.WriteAndWaitForReport(cmd, time.Second*2, 0x08) // Request node information
	report := <-t

	logrus.Debugf("Request endpoint %d report: %#v", n.Id, report)

	if report != nil {
		switch cmd := report.(type) {
		case *commands.MultiChannelCmdEndPointReport:
			//n.ManufacurerSpecific = cmd
			logrus.Debug(cmd.String())
			n.Lock()
			n.Endpoints = make([]*Endpoint, 0)
			for i := 1; i < cmd.Endpoints; i++ {
				n.Endpoints = append(n.Endpoints, &Endpoint{
					Id:             i,
					CommandClasses: n.Device.CommandClasses,

					StateBool:  make(map[string]bool),
					StateFloat: make(map[string]float64),

					node: n,
				})
			}
			n.Unlock()
			return nil
		default:
			logrus.Errorf("Wrong type: %t", cmd)
		}
	}

	//}

	return fmt.Errorf("Failed to request endpoints")
}
