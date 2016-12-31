package reports

import "fmt"

type SwitchMultilevel struct {
	*report
	Level byte `json:"level"`
	data  []byte
}

func NewSwitchMultilevel(data []byte) (*SwitchMultilevel, error) {
	if len(data) < 1 {
		data = []byte{0x00}
	}

	return &SwitchMultilevel{
		Level: data[0],
		data:  data,
	}, nil
}

func (sm SwitchMultilevel) String() string {
	return fmt.Sprintf("Level: %d", sm.Level)
}
