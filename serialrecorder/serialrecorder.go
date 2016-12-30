package serialrecorder

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"time"

	"github.com/tarm/serial"
)

type serialRecorder struct {
	serial io.ReadWriteCloser
	port   string
	baud   int
	Logger io.Writer

	writeBuffer chan *Row
}

// Row logs one serial read or write
type Row struct {
	Timestamp time.Time
	Direction Direction
	Data      []byte
}

// New returns a new serialRecorder that implements PortOpener for serial port that will record and output each write and read to stdout
func New(port string, baud int) *serialRecorder {
	return &serialRecorder{
		port:        port,
		baud:        baud,
		Logger:      os.Stdout,
		writeBuffer: make(chan *Row, 100),
	}
}

func (sr *serialRecorder) Read(p []byte) (n int, err error) {
	n, err = sr.serial.Read(p)
	go sr.writeToLog(DirectionRead, p[:n])
	return
}

func (sr *serialRecorder) Write(p []byte) (n int, err error) {
	n, err = sr.serial.Write(p)
	go sr.writeToLog(DirectionWrite, p[:n])
	return
}

func (sr *serialRecorder) writeToLog(d Direction, data []byte) {
	sr.writeBuffer <- &Row{
		Timestamp: time.Now(),
		Direction: d,
		Data:      data,
	}
}

func (sr *serialRecorder) logWriter() {
	enc := gob.NewEncoder(sr.Logger)

	for row := range sr.writeBuffer {
		err := enc.Encode(row)
		if err != nil {
			log.Println(err)
		}
	}
}

func (sr *serialRecorder) Close() error {
	return sr.serial.Close()
}
func (sr *serialRecorder) Open() (io.ReadWriteCloser, error) {
	go sr.logWriter()

	var err error
	sr.serial, err = serial.OpenPort(&serial.Config{Name: sr.port, Baud: sr.baud})
	return sr, err
}
