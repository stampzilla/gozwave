package gozwave

import (
	"io"
	"log"
)

// generate with impl 'm *mockSerial' io.ReadWriteCloser
type mockSerial struct {
	sendToRead chan []byte
	writeLog   [][]byte
}

func newMockSerial() *mockSerial {
	return &mockSerial{
		sendToRead: make(chan []byte),
	}
}

func (m *mockSerial) Read(p []byte) (n int, err error) {
	data := <-m.sendToRead
	n = copy(p, data)
	log.Printf("%x", p)
	return len(p), nil
}

func (m *mockSerial) Write(p []byte) (n int, err error) {
	log.Printf("Writing %x\n", p)
	m.writeLog = append(m.writeLog, p)
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
