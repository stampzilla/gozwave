package functions

import "fmt"

type FuncGetNodeProtocolInfo struct {
	Listening bool
	Routing   bool
	Beaming   bool

	Version byte

	Flirs    bool
	Security bool
	MaxBaud  int

	Basic    byte
	Generic  byte
	Specific byte

	data []byte
}

func NewGetNodeProtocolInfo(data []byte) *FuncGetNodeProtocolInfo {
	f := &FuncGetNodeProtocolInfo{}
	f.Decode(data)

	return f
}

func (pi *FuncGetNodeProtocolInfo) Decode(data []byte) error {
	pi.data = data
	if len(data) < 6 {
		return fmt.Errorf("GetNodeProtocolInfo message is to short, should be at least 6 bytes")
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

	return nil
}
