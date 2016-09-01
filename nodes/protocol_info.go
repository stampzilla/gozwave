package nodes

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
	"github.com/stampzilla/gozwave/functions"
)

func (n *node) RequestProtocolInfo() error {
	// Todo: Send raw messages here

	resp := <-n.connection.SendRaw([]byte{
		functions.GetNodeProtocolInfo, // Function
		byte(n.Id()),                  // Node id
	}, time.Second*1) // Request node information

	logrus.Infof("RESP: %#v", resp)

	if resp != nil {
		switch r := resp.Data.(type) {
		case *functions.FuncGetNodeProtocolInfo:
			n.ProtocolInfo = r
		default:
			spew.Dump("Wrong type: %t", resp.Data)
		}
	}

	return nil
}
