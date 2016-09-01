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
	Nodes      *nodes.List
	Connection *serialapi.Connection
	eventQue   chan interface{}
}

func (self *Controller) getNodes() (*nodes.List, error) {
	if self.eventQue == nil {
		self.eventQue = make(chan interface{}, 10)
	}

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

			node, basicIdentificationDone := nodes.New(byte(index+1), self.Connection, self.eventQue)
			self.Nodes.Add(node)
			<-basicIdentificationDone
		}
	default:
		logrus.Errorf("Got wrong response type: %t", r)
	}

	return self.Nodes, nil
}

func (self *Controller) PushEvent(event interface{}) {
	if self.eventQue == nil {
		self.eventQue = make(chan interface{}, 10)
	}

	select {
	case self.eventQue <- event:
	default:
	}
}

func (self *Controller) GetNextEvent() chan interface{} {
	if self.eventQue == nil {
		self.eventQue = make(chan interface{}, 10)
	}

	return self.eventQue
}
