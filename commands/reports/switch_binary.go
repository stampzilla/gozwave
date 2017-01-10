package reports

import (
	"fmt"

	"github.com/stampzilla/gozwave/commands"
)

// SwitchBinaryV1 is sent from a zwave node to advertise the status of a device with On/Off or
// Enable/Disable capability. A known issue is that nodes with multiple endpoints will send
// a SwichBinary report on all endpoint binary state changes but without encapsulating it in a
// MultiChannel message.
type SwitchBinaryV1 struct {
	*report
	CurrentValue bool `json:"current_value"` // The current state of the node
	Valid        bool `json:"valid"`         // If the state in known or not. Some devices report "unknown" state

	data []byte
}

// NewSwitchBinaryV1 creates a new SwitchBinary report that contains the switch status from a
// zwave node
func NewSwitchBinaryV1(data []byte) (*SwitchBinaryV1, error) {
	if len(data) != 1 {
		return nil, fmt.Errorf("Wrong length, expected 1 byte got %d", len(data))
	}

	return &SwitchBinaryV1{
		CurrentValue: data[0] == 0xFF && data[0] != 0xFE,
		Valid:        data[0] != 0xFE,
		data:         data,
	}, nil
}

func (sb *SwitchBinaryV1) String() string {
	return fmt.Sprintf("current:%t", sb.CurrentValue)
}

// SwitchBinaryV2 is sent from a zwave node to advertise the status of a device with On/Off or
// Enable/Disable capability. A known issue is that nodes with multiple endpoints will send
// a SwichBinary report on all endpoint state changes but without encapsulating it in a
// MultiChannel message.
type SwitchBinaryV2 struct {
	*report
	CurrentValue bool              `json:"current_value"` // The current state of the node
	TargetValue  bool              `json:"target_value"`  // The target state of the node
	Duration     commands.Duration `json:"duration"`      // The transition time between current and target values

	data []byte
}

// NewSwitchBinaryV2 creates a new switchBinary report that contains the switch status from a
// zwave node
func NewSwitchBinaryV2(data []byte) (*SwitchBinaryV2, error) {
	if len(data) != 3 {
		return nil, fmt.Errorf("Wrong length, expected 3 byte got %d", len(data))
	}

	return &SwitchBinaryV2{
		CurrentValue: data[0] != 0x00,
		TargetValue:  data[1] != 0x00,
		Duration:     commands.Duration(data[2]),
		data:         data,
	}, nil
}

func (sb *SwitchBinaryV2) String() string {
	return fmt.Sprintf("current:%t target:%t duration:%s", sb.CurrentValue, sb.TargetValue, sb.Duration)
}
