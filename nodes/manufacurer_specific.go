package nodes

import (
	"fmt"
	"time"

	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/commands/reports"
)

type ManufacurerSpecific struct {
}

func (n *Node) RequestManufacturerSpecific() (*reports.ManufacturerSpecific, error) {
	cmd := commands.NewRaw(
		[]byte{
			commands.ManufacturerSpecific, // Command
			0x04, // MANUFACTURER_SPECIFIC_GET
			0x00,
			//0x05, // TransmitOptions?
			//0x23, // Callback?
		})

	cmd.SetNode(n.Id)

	t, err := n.connection.WriteAndWaitForReport(cmd, time.Second*10, 0x05) // Request node information
	if err != nil {
		return nil, err
	}

	report := <-t

	switch cmd := report.(type) {
	case *reports.ManufacturerSpecific:
		return cmd, nil
	default:
		return nil, fmt.Errorf("ManufacurerSpecific: Wrong report type: %t", report)
	}

	return nil, fmt.Errorf("Failed to get ManufacurerSpecific")
}
