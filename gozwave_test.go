package gozwave

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {

	mockPO := newMockPortOpener()
	controller, err := ConnectWithCustomPortOpener("/test", "", mockPO)
	assert.NoError(t, err)

	log.Println("controller: ", controller)

	time.Sleep(3 * time.Second)
}
