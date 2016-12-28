package nodes

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/functions"
)

func (n *Node) RequestProtocolInfo() (*functions.FuncGetNodeProtocolInfo, error) {
	cmd := functions.NewRaw(
		[]byte{
			functions.GetNodeProtocolInfo, // Function
			byte(n.Id),                    // Node id
		})

	t, err := n.connection.WriteWithTimeout(cmd, time.Second*10)
	if err != nil {
		return nil, err
	}

	resp := <-t

	logrus.Debugf("RequestProtocolInfo RESP: %#v", resp)

	if resp != nil {
		switch r := resp.Data.(type) {
		case *functions.FuncGetNodeProtocolInfo:
			return r, nil
		default:
			return nil, fmt.Errorf("Wrong type: %t", resp.Data)
		}
	}

	return nil, fmt.Errorf("Reponse was nil")
}
