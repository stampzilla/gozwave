package nodes

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/functions"
)

func (n *Node) RequestEndpoints() error {
	// Todo: Send raw messages here

	if n.Device == nil || n.Device.CommandClasses == nil {
		n.Endpoints = make([]*Endpoint, 0)
		return fmt.Errorf("Failed 'RequestEndpoints', no commandclasses exists")
	}

	for _, v := range n.Device.CommandClasses {
		logrus.Errorf("Request endpoint %d:%x - %s", n.Id, byte(v.ID), v.ID)
	}
	fmt.Println("")
	fmt.Println("")
	logrus.Errorf("Request endpoint %d", n.Id)
	resp := <-n.connection.SendRawAndWaitForResponse([]byte{
		functions.SendData, // Function
		byte(n.Id),         // Node id
		0x02,               // Length
		commands.MultiInstance, // Command class
		0x07, // Command: MultiInstanceCmd_Get
		//byte(v.ID),
		0x25, // TransmitOptions?
		//0x23, // Callback?
	}, time.Second*2, 0x08) // Request node information

	logrus.Errorf("Request endpoint %d RESP: %#v", n.Id, resp)

	if resp != nil {
		switch r := resp.Data.(type) {
		case *functions.FuncApplicationCommandHandler:
			switch cmd := r.Data.(type) {
			case *commands.MultiChannelCmdEndPointReport:
				//n.ManufacurerSpecific = cmd
				logrus.Info(cmd.String())

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
				return nil
			default:
				logrus.Errorf("Wrong type: %t", r.Command)
			}
		}
	}

	//}

	return fmt.Errorf("Failed to request endpoints")
}
