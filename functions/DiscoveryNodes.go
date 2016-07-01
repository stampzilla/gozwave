package functions

import "fmt"

type FuncDiscoveryNodes struct {
	ActiveNodes [232]bool

	data []byte
}

func NewDiscoveryNodes(data []byte) *FuncDiscoveryNodes {
	f := &FuncDiscoveryNodes{}
	f.Decode(data)

	return f
}

func (self *FuncDiscoveryNodes) Decode(data []byte) error {
	self.data = data
	if len(data) < 6 {
		return fmt.Errorf("GetNodeProtocolInfo message is to short, should be at least 6 bytes")
	}

	for index, bits := range data {
		self.ActiveNodes[index*8+0] = bits&0x01 != 0
		self.ActiveNodes[index*8+1] = bits&0x02 != 0
		self.ActiveNodes[index*8+2] = bits&0x04 != 0
		self.ActiveNodes[index*8+3] = bits&0x08 != 0

		self.ActiveNodes[index*8+4] = bits&0x10 != 0
		self.ActiveNodes[index*8+5] = bits&0x20 != 0
		self.ActiveNodes[index*8+6] = bits&0x40 != 0
		self.ActiveNodes[index*8+7] = bits&0x80 != 0
	}

	return nil
}
