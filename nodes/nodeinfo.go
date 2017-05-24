package nodes

import (
	"fmt"
	"time"

	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/database"
	"github.com/stampzilla/gozwave/serialapi"
)

func (n *Node) RequestNodeInfo() ([]*database.CommandClass, error) {
	cmd := serialapi.NewRaw(
		[]byte{
			serialapi.RequestNodeInfo, // Command
			byte(n.Id),                // Node id
		})

	t, err := n.connection.WriteWithTimeout(cmd, time.Second*10)
	if err != nil {
		return nil, err
	}

	report := <-t

	if report == nil {
		return nil, fmt.Errorf("Failed to get NodeInfo (nil)")
	}

	switch cmd := report.Data.(type) {
	case *serialapi.FuncApplicationUpdate:
		cc := make([]*database.CommandClass, 0)
		for _, v := range cmd.Supported {
			cc = append(cc, &database.CommandClass{
				ID:         commands.ZWaveCommand(v),
				Controlled: "false",
				InNIF:      "true",
			})
		}

		for _, v := range cmd.Controllable {
			cc = append(cc, &database.CommandClass{
				ID:         commands.ZWaveCommand(v),
				Controlled: "true",
				InNIF:      "true",
			})
		}

		return cc, nil
	default:
		return nil, fmt.Errorf("NodeInfo: Wrong report type: %t", report)
	}

	return nil, fmt.Errorf("Failed to get NodeInfo")
}
