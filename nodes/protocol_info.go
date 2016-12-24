package nodes

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/functions"
)

func (n *Node) RequestProtocolInfo() error {
	// Todo: Send raw messages here

	resp := <-n.connection.SendRaw([]byte{
		functions.GetNodeProtocolInfo, // Function
		byte(n.Id),                    // Node id
	}, time.Second*10) // Request node information

	logrus.Infof("RESP: %#v", resp)

	if resp != nil {
		switch r := resp.Data.(type) {
		case *functions.FuncGetNodeProtocolInfo:
			n.Lock()
			n.ProtocolInfo = r
			n.Unlock()
			return nil
		default:
			return fmt.Errorf("Wrong type: %t", resp.Data)
		}
	}

	return fmt.Errorf("Reponse was nil")
}
