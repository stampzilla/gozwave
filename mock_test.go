package gozwave

import (
	"encoding/hex"
	"io"
	"log"
	"testing"
)

// generate with impl 'm *mockSerial' io.ReadWriteCloser
type mockSerial struct {
	sendToRead   chan string
	getFromWrite chan []byte
	writeLog     [][]byte
	errorLog     []error
}

func newMockSerial() *mockSerial {
	return &mockSerial{
		sendToRead:   make(chan string),
		getFromWrite: make(chan []byte),
	}
}

func (m *mockSerial) Read(p []byte) (n int, err error) {
	data := <-m.sendToRead

	bytes, err := hex.DecodeString(data)
	if err != nil {
		m.errorLog = append(m.errorLog, err)
		return 0, err
	}

	n = copy(p, bytes)
	if testing.Verbose() {
		log.Printf("MOCK read  %x", p[:n])
	}
	return n, nil
}

func (m *mockSerial) Write(p []byte) (n int, err error) {
	if testing.Verbose() {
		log.Printf("MOCK write %x\n", p)
	}
	m.writeLog = append(m.writeLog, p)
	m.getFromWrite <- p
	return len(p), nil
}

func (m *mockSerial) Close() error {
	log.Printf("CLOSE")
	return nil
}

type mockPortOpener struct {
	mockSerial *mockSerial
}

func (mpo *mockPortOpener) Open() (io.ReadWriteCloser, error) {
	return mpo.mockSerial, nil

}

func newMockPortOpener() *mockPortOpener {
	return &mockPortOpener{
		mockSerial: newMockSerial(),
	}
}
