//go:generate go run generators/database.go -file database/test.go -package database -databasedir ./database

package gozwave

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/nodes"
)

func Connect(port string, filename string) (*Controller, error) {

	c := &Controller{
		Nodes:           nodes.NewList(),
		Connection:      NewConnection(),
		filename:        filename,
		eventQue:        make(chan interface{}, 10),
		triggerFileSave: make(chan struct{}),
	}

	c.Connection.Name = port
	c.Connection.Baud = 115200
	c.Connection.reportCallback = c.DeliverReportToNode
	connected := make(chan error)

	//spew.Dump(c)

	if filename != "" {
		c.LoadConfigurationFromFile()
	}

	go func() {
		for {
			err := c.Connection.Connect(connected)
			logrus.Error(err)
			<-time.After(time.Second)
		}
	}()

	err := <-connected

	go c.getNodes()
	go c.saveDebouncer()

	return c, err
}
