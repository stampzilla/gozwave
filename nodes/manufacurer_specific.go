package nodes

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/functions"
)

type ManufacurerSpecific struct {
}

func (n *Node) RequestManufacturerSpecific() error {
	// Todo: Send raw messages here

	resp := <-n.connection.SendRawAndWaitForResponse([]byte{
		functions.SendData, // Function
		byte(n.Id),         // Node id
		0x02,               // Length
		commands.ManufacturerSpecific, // Command
		0x04, // MANUFACTURER_SPECIFIC_GET
		0x00,
		//0x05, // TransmitOptions?
		//0x23, // Callback?
	}, time.Second*10, 0x05) // Request node information

	if resp != nil {
		switch r := resp.Data.(type) {
		case *functions.FuncApplicationCommandHandler:
			switch cmd := r.Data.(type) {
			case *commands.CmdManufacturerSpecific:
				n.ManufacurerSpecific = cmd
				return nil
			default:
				spew.Dump("Wrong type: %t", r.Command)
			}
		}
	}

	return fmt.Errorf("Failed to get ManufacurerSpecific")
}
