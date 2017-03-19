//go:generate go run generators/database/database.go -file database/test.go -package database -databasedir ./database
//go:generate go run generators/commandclasses/commandclasses.go -file database/mandatory.go -package database -mandatoryfile ./database/mandatory.txt

package gozwave

import (
	"fmt"
	"io"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/tarm/serial"
)

func ConnectWithCustomPortOpener(port string, filename string, po PortOpener) (*Controller, error) {
	if po == nil {
		return nil, fmt.Errorf("portopener is nil")
	}

	c := NewController()
	c.filename = filename
	c.Connection.portOpener = po

	var err error

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

	err = <-connected
	if err != nil {
		return nil, err
	}

	go c.getNodes()
	go c.saveDebouncer()

	return c, err
}
func Connect(port string, filename string) (*Controller, error) {
	return ConnectWithCustomPortOpener(port, filename, newSerialPortOpener(port, 115200))
}

type PortOpener interface {
	/* TODO: add methods */
	Open() (io.ReadWriteCloser, error)
}

type serialPortOpener struct {
	port string
	baud int
}

func (spo *serialPortOpener) Open() (io.ReadWriteCloser, error) {
	return serial.OpenPort(&serial.Config{Name: spo.port, Baud: spo.baud})
}

func newSerialPortOpener(port string, baud int) *serialPortOpener {
	return &serialPortOpener{
		port: port,
		baud: baud,
	}
}
