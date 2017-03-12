package serialapi

import "fmt"

type FuncGetNodeProtocolInfo struct {
	Listening bool `json:"listening"`
	Routing   bool `json:"routing"`
	Beaming   bool `json:"beaming"`

	Version byte `json:"version"`

	Flirs    bool `json:"flirs"`
	Security bool `json:"security"`
	MaxBaud  int  `json:"max_baud"`

	Basic    byte `json:"basic"`
	Generic  byte `json:"generic"`
	Specific byte `json:"specific"`
}

// NewGetNodeProtocolInfo decodes and cretes a funcGetNodeProtocolInfo. It is sent from the controller and contanins all basic information that the controller knows about a node.
func NewGetNodeProtocolInfo(data []byte) (*FuncGetNodeProtocolInfo, error) {
	pi := &FuncGetNodeProtocolInfo{}

	if len(data) < 6 {
		return nil, fmt.Errorf("wrong length, should be at least 6 bytes got len %d data: %x", len(data), data)
	}

	// Capabilities
	pi.Listening = data[0]&0x80 != 0
	pi.Routing = data[0]&0x40 != 0
	pi.Version = data[0]&0x07 + 1

	pi.MaxBaud = 9600
	if data[0]&0x38 == 0x10 {
		pi.MaxBaud = 40000
	}

	// Security
	pi.Flirs = data[1]&0x60 != 0
	pi.Beaming = data[1]&0x10 != 0
	pi.Security = data[1]&0x01 != 0

	// data[2] - reserved

	pi.Basic = data[3]
	pi.Generic = data[4]
	pi.Specific = data[5]

	return pi, nil
}
