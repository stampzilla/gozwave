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

func (self *FuncGetNodeProtocolInfo) Decode(data []byte) error {
	self.data = data
	if len(data) < 6 {
		return fmt.Errorf("GetNodeProtocolInfo message is to short, should be at least 6 bytes")
	}

	// Capabilities
	self.Listening = data[0]&0x80 != 0
	self.Routing = data[0]&0x40 != 0
	self.Version = data[0]&0x07 + 1

	self.MaxBaud = 9600
	if data[0]&0x38 == 0x10 {
		self.MaxBaud = 40000
	}

	// Security
	self.Flirs = data[1]&0x60 != 0
	self.Beaming = data[1]&0x10 != 0
	self.Security = data[1]&0x01 != 0

	// data[2] - reserved

	self.Basic = data[3]
	self.Generic = data[4]
	self.Specific = data[5]

	return nil
}
