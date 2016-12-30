package serialrecorder

import (
	"bytes"
	"encoding/gob"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWriteToLog(t *testing.T) {

	buf := bytes.NewBufferString("")
	sr := New("asfd", 100)
	sr.Logger = buf
	go sr.logWriter()

	sr.writeToLog(DirectionWrite, []byte("asdf"))

	time.Sleep(time.Millisecond * 100)
	log.Println(buf)

	dec := gob.NewDecoder(buf) // Will read from network.

	r := &Row{}
	err := dec.Decode(&r)
	assert.NoError(t, err)

	assert.Equal(t, "asdf", string(r.Data))
	assert.Equal(t, DirectionWrite, r.Direction)

	if time.Now().Before(r.Timestamp) {
		t.Error("Timestamp invalid. To old")
	}
}
