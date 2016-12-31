package reports

import "fmt"

type SwitchBinary struct {
	*report
	Status bool `json:"status"`
	data   []byte
}

// NewSwitchBinary creates a new switchBinary report that contains the switch status from a zwave node
func NewSwitchBinary(data []byte) (*SwitchBinary, error) {
	if len(data) < 1 {
		data = []byte{0x00}
	}

	return &SwitchBinary{
		Status: data[0] == 0xFF,
		data:   data,
	}, nil
}

func (sb *SwitchBinary) String() string {
	return fmt.Sprintf("Status: %t", sb.Status)
}
