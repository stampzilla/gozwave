package serialapi

import (
	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/commands/reports"
)

type FuncApplicationCommandHandler struct {
	Command commands.ZWaveCommand `json:"command"`
	Class   byte                  `json:"class"`
	Node    byte                  `json:"node"`

	Report reports.Report `json:"report"`
}

// NewApplicationCommandHandler decodes and creates a funcApplicationCommandHandler. This is the most common message from the controller. Its sent everytime a report are received from the zwave network.
func NewApplicationCommandHandler(data []byte) (*FuncApplicationCommandHandler, error) {
	f := &FuncApplicationCommandHandler{}

	f.Command = commands.ZWaveCommand(data[0])
	f.Class = data[1]

	var err error
	f.Report, err = reports.New(f.Command, f.Class, data[2:])
	if err != nil {
		return nil, err
	}

	if report, ok := f.Report.(reports.Report); ok && report != nil {
		report.SetNode(f.Node)
	}

	return f, nil
}
