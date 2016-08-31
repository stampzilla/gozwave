package gozwave

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/functions"
	"github.com/stampzilla/gozwave/nodes"
	"github.com/stampzilla/gozwave/serialapi"
)

type Controller struct {
	Nodes      nodes.List
	Connection *serialapi.Connection
}

func (self *Controller) getNodes() (nodes.List, error) {
	resp := <-self.Connection.SendRaw([]byte{0x02}, time.Second)
	if resp == nil {
		return nil, fmt.Errorf("Send failed, timeut?")
	}

	switch r := resp.Data.(type) {
	case *functions.FuncDiscoveryNodes:
		for index, active := range r.ActiveNodes {
			if !active {
				continue
			}

			node := nodes.New(byte(index+1), self.Connection)
			self.Nodes.Add(node)
		}
	default:
		logrus.Errorf("Got wrong response type: %t", r)
	}

	return self.Nodes, nil
}
