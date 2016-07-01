package gozwave

import (
	"encoding/hex"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/pborman/uuid"
	"github.com/stampzilla/gozwave/functions"
	"github.com/tarm/goserial"
)

type Controller struct {
	Name      string
	Baud      int
	port      io.ReadWriteCloser
	Connected bool

	// Keep track of requests wating a response
	inFlight map[string]chan *Message

	sync.Mutex
}

func New(port string) *Controller {
	z := &Controller{
		inFlight: make(map[string]chan *Message),
	}
	z.Name = port
	z.Baud = 115200

	return z
}

func (self *Controller) Connect() error {

	connected := make(chan error)

	go func() {
		for {
			err := self.connect(connected)
			logrus.Error(err)
			<-time.After(time.Second)
		}
	}()

	return <-connected
}

func (self *Controller) Send(data []byte) chan *Message {
	returnChan := make(chan *Message)

	// Compile message
	msg := append([]byte{0x00, 0x00}, data...)
	msg = append(msg, 0x00)

	// Add length
	msg[0] = byte(len(msg) - 1)

	// Add checksum
	msg[len(msg)-1] = generateChecksum(msg)

	// Add header
	msg = append([]byte{0x01}, msg...)

	fmt.Println("Sending: ")
	fmt.Println(hex.Dump(msg))
	self.Lock()
	self.port.Write(msg)
	self.Unlock()

	self.inFlight[uuid.New()] = returnChan

	return returnChan
}

func (self *Controller) GetNodes() (err error) {

	resp := <-self.Send([]byte{0x02})
	switch r := resp.Data.(type) {
	case *functions.FuncDiscoveryNodes:
		for index, active := range r.ActiveNodes {
			if !active {
				continue
			}

			<-self.Send([]byte{0x41, byte(index + 1)}) // Request node information
			<-self.Send([]byte{0x62, byte(index + 1)}) // Request is failed node
			<-self.Send([]byte{0xa0, byte(index + 1)}) // Request is failed node
		}
	default:
		logrus.Errorf("Got wrong response type: %t", r)
	}

	return nil
}

func (self *Controller) connect(connectChan chan error) (err error) {
	defer func() {
		fmt.Print("Disonnected\n\n")
		self.Connected = false
	}()

	self.Lock()
	c := &serial.Config{Name: self.Name, Baud: self.Baud}
	self.port, err = serial.OpenPort(c)
	self.Unlock()

	if err != nil {
		select {
		case connectChan <- err:
		default:
		}
		return
	}

	select {
	case connectChan <- nil:
	default:
	}
	self.Connected = true
	fmt.Print("Connected\n\n")

	//<-time.After(time.Millisecond * 100)
	////self.port.Write([]byte{0x01, 0x09, 0x00, 0x13, 0x06, 0x03, 0x20, 0x01, 0xFF, 0x05, 0x3B})
	//msg := []byte{0x01, 0x03, 0x00, 0x07, 0xfb}
	//self.port.Write(msg)

	incomming := make([]byte, 0)

	for {
		buf := make([]byte, 128)

		n, err := self.port.Read(buf)
		if err != nil {
			logrus.Error("Serial read failed: ", err)
			self.port.Close()
			return err
		}

		//fmt.Println(hex.Dump(buf[:n]))
		incomming = append(incomming, buf[:n]...)

		for len(incomming) > 0 {
			l, msg := Decode(incomming)

			if msg != nil {
				for index, c := range self.inFlight {
					select {
					case c <- msg: // Try to deliver message
					default: // Else close and remove
						delete(self.inFlight, index)
						close(c)
					}
				}
			}

			if l <= len(incomming) { // If complete message was decoded, remove it from buffer
				incomming = incomming[l:]
				if l > 1 {
					self.port.Write([]byte{0x06}) // Send ACK to stick
				}
			} else {
				break
			}
		}
		//fmt.Println(n)

		//for {
		//n := strings.Index(incomming, footer)

		//if n < 0 {
		//break
		//}

		//msg := incomming[:n]
		//incomming = incomming[n+len(footer):]

		////go callback(msg)
		//}
	}

	return nil
}
