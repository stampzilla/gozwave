package gozwave

import (
	"fmt"
	"log"
	"os"
	"testing"
	"text/tabwriter"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {

	mockPO := newMockPortOpener()

	controller, err := ConnectWithCustomPortOpener("/test", "", mockPO)
	assert.NoError(t, err)

	log.Println("controller: ", controller)

	reply(t, mockPO.mockSerial.getFromWrite, mockPO.mockSerial.sendToRead) // Start up a conversation loop
	//time.Sleep(10 * time.Second)
}

func reply(t *testing.T, c chan []byte, w chan string) {
	replies := map[string][]string{
		"06": []string{},
		"01030002fe": []string{ // Request discovery nodes
			"0125010205001d8000000000000000000000000000000000000000000000000000000000050044", // Answer with node 8 active
		},
		"0104004108b2": []string{ // Request node information node 8
			"06", // ack
			"01090141539c000440003d", // Answer with node information
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
