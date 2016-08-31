package gozwave

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/nodes"
	"github.com/stampzilla/gozwave/serialapi"
)

//type Controller struct {
//Name      string
//Baud      int
//port      io.ReadWriteCloser
//Connected bool

//// Keep track of requests wating a response
//inFlight map[string]chan *Message

//sync.Mutex
//}

func Connect(port string, file ...string) (*Controller, error) {

	c := &Controller{
		Nodes:      nodes.NewList(),
		Connection: serialapi.NewConnection(),
	}

	c.Connection.Name = port
	c.Connection.Baud = 115200
	connected := make(chan error)

	go func() {
		for {
			err := c.Connection.Connect(connected)
			logrus.Error(err)
			<-time.After(time.Second)
		}
	}()

	err := <-connected

	go c.getNodes()

	return c, err
}
