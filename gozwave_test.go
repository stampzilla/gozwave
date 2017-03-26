package gozwave

import (
	"fmt"
	"io"
	"os"
	"testing"
	"text/tabwriter"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/serialapi"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	if testing.Verbose() {
		logrus.SetLevel(logrus.DebugLevel)
	}

	mockPO := newMockPortOpener()
	controller, err := ConnectWithCustomPortOpener("/test", "", mockPO)

	assert.NoError(t, err)
	//log.Println("controller: ", controller)

	reply(t, mockPO.mockSerial.getFromWrite, mockPO.mockSerial.sendToRead) // Start up a conversation loop

	// TODO: make something better than a sleep here
	time.Sleep(100 * time.Millisecond)

	n := controller.Nodes.Get(8)
	assert.NotNil(t, n)
	if assert.NotNil(t, n.ProtocolInfo()) {
		assert.Equal(t,
			serialapi.FuncGetNodeProtocolInfo{
				Listening: false,
				Routing:   true,
				Beaming:   true,
				Version:   0x04,
				Flirs:     false,
				Security:  false,
				MaxBaud:   40000,
				Basic:     0x04,
				Generic:   0x40,
				Specific:  0x0,
			},
			*n.ProtocolInfo(),
		)
	}
}

func reply(t *testing.T, c chan []byte, w chan string) {
	replies := map[string][]string{
		"06": []string{},
		"01030002fe": []string{ // Request discovery nodes
			"0125010205001d8000000000000000000000000000000000000000000000000000000000050044", // Answer with node 8 active
		},
		"0104004108b2": []string{ // Request node information node 8
			"06", // ack
			"01080141539c0004403c", // Answer with node information
		},
	}

	reads := map[string]int{} // Count the number of reads received

	// Receive messages from gozwave and answer
	for {
		select {
		case v := <-c:
			s := fmt.Sprintf("%x", v) // Encode to hex string

			if r, ok := replies[s]; ok {
				// Keep a count of how many times we have received each message
				if _, ok := reads[s]; !ok {
					reads[s] = 0
				}
				reads[s]++

				for _, rep := range r {
					w <- rep // Send the response
				}

				if len(reads) == len(replies) { // We got all messages expected. Exit the loop
					//<-time.After(time.Second)
					return
				}
				break
			}

			t.Fatalf("Unexpected read %s", s)
			return
		case <-time.After(time.Second * 3):
			fmt.Println("Message counts\n----------------------------------")
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.AlignRight|tabwriter.Debug)
			for k := range replies {
				fmt.Fprintf(w, "%d\t%s\n", reads[k], k)
			}
			w.Flush()
			fmt.Println("----------------------------------")

			for k := range replies {
				if _, ok := reads[k]; !ok {
					t.Errorf("Test timeout. Have not received message %s", k)
				}
			}
			return
		}
	}
}

type poMock struct{}

func (po *poMock) Open() (io.ReadWriteCloser, error) {
	return nil, fmt.Errorf("custom mock connect error")
}

func TestConnectWithCustomPortOpener(t *testing.T) {
	// Test failed portopener
	c, err := ConnectWithCustomPortOpener("", "", nil)

	assert.EqualError(t, err, "portopener is nil")
	assert.Nil(t, c)

	// Test failed connect
	c, err = ConnectWithCustomPortOpener("", "", &poMock{})

	assert.EqualError(t, err, "custom mock connect error")
	assert.Nil(t, c)
}
