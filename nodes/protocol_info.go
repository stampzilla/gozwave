package nodes

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/serialapi"
)

func (n *Node) ProtocolInfo() *serialapi.FuncGetNodeProtocolInfo {
	n.RLock()
	defer n.RUnlock()

	return n.protocolInfo
}

func (n *Node) RequestProtocolInfo() (*serialapi.FuncGetNodeProtocolInfo, error) {
	cmd := serialapi.NewRaw(
		[]byte{
			serialapi.GetNodeProtocolInfo, // Function
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
		case *serialapi.FuncGetNodeProtocolInfo:
			return r, nil
		default:
			return nil, fmt.Errorf("Wrong type: %t", resp.Data)
		}
	}

	return nil, fmt.Errorf("Reponse was nil")
}
