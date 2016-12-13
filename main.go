//go:generate go run generators/database.go -file database/test.go -package database -databasedir ./database

package gozwave

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/nodes"
	"github.com/stampzilla/gozwave/serialapi"
)

func Connect(port string, filename string) (*Controller, error) {

	c := &Controller{
		Nodes:      nodes.NewList(),
		Connection: serialapi.NewConnection(),
		filename:   filename,
	}

	c.Connection.Name = port
	c.Connection.Baud = 115200
	c.Connection.ConfigController = c
	connected := make(chan error)

	//spew.Dump(c)

	c.Nodes.SetConnection(c.Connection)
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

	return c, err
}
