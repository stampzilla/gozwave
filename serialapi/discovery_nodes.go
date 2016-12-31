package serialapi

import "fmt"

type DiscoverdNodes struct {
	ActiveNodes [232]bool `json:"active_nodes"`

	data []byte
}

// NewDiscoverdNodes decodes and creates a discoveryNode message. This message is received from the controller and contains a list of all available nodes in the network
func NewDiscoverdNodes(data []byte) (*DiscoverdNodes, error) {
	f := &DiscoverdNodes{
		data: data,
	}

	if len(data) < 6 {
		return nil, fmt.Errorf("message is to short, should be at least 6 bytes")
	}

	for index, bits := range data {
		f.ActiveNodes[index*8+0] = bits&0x01 != 0
		f.ActiveNodes[index*8+1] = bits&0x02 != 0
		f.ActiveNodes[index*8+2] = bits&0x04 != 0
		f.ActiveNodes[index*8+3] = bits&0x08 != 0

		f.ActiveNodes[index*8+4] = bits&0x10 != 0
		f.ActiveNodes[index*8+5] = bits&0x20 != 0
		f.ActiveNodes[index*8+6] = bits&0x40 != 0
		f.ActiveNodes[index*8+7] = bits&0x80 != 0
	}

	return f, nil
}
