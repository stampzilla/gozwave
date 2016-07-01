package gozwave

import (
	"time"

	"github.com/Sirupsen/logrus"
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

func Connect(port string) (*serialapi.Connection, error) {

	z := serialapi.NewConnection()
	z.Name = port
	z.Baud = 115200
	connected := make(chan error)

	go func() {
		for {
			err := z.Connect(connected)
			logrus.Error(err)
			<-time.After(time.Second)
		}
	}()

	return z, <-connected
}
