package gozwave

import (
	"fmt"
	"io"

	"github.com/tarm/serial"
)

type serialRecorder struct {
	serial io.ReadWriteCloser
	port   string
	baud   int
}

// NewSerialRecorder returns a new PortOpener for serial port that will record and output each write and read to stdout
func NewSerialRecorder(port string, baud int) PortOpener {
	return &serialRecorder{
		port: port,
		baud: baud,
	}
}

func (sr *serialRecorder) Read(p []byte) (n int, err error) {
	fmt.Printf("%x", p)
	return sr.serial.Read(p)
}

func (sr *serialRecorder) Write(p []byte) (n int, err error) {
	fmt.Printf("%x", p)
	return sr.serial.Write(p)
}

func (sr *serialRecorder) Close() error {
	return sr.serial.Close()
}
func (sr *serialRecorder) Open() (io.ReadWriteCloser, error) {
	var err error
	sr.serial, err = serial.OpenPort(&serial.Config{Name: sr.port, Baud: sr.baud})
	return sr, err
}
