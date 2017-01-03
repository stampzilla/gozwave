package reports

import (
	"fmt"
	"math"

	"github.com/stampzilla/gozwave/commands"
)

// SwitchMultilevelV1 is sent by a zwave node to advertise the status of a multilevel device.
// A known issue is that nodes with multiple endpoints will send a SwichMultilevel report on
// all endpoint multilevel state changes but without encapsulating it in a MultiChannel message.
type SwitchMultilevelV1 struct {
	*report
	CurrentValue float64 `json:"current_value"` // The current value state
	Valid        bool    `json:"valid"`         // If the state in known or not. Some devices report "unknown" state
	data         []byte
}

// NewSwitchMultilevelV1 creates a new SwitchMultilevelV1 report from binary data sent from a
// zwave node
func NewSwitchMultilevelV1(data []byte) (*SwitchMultilevelV1, error) {
	if len(data) != 1 {
		return nil, fmt.Errorf("Wrong length, expected 1 byte got %d", len(data))
	}

	return &SwitchMultilevelV1{
		CurrentValue: math.Min(float64(data[0]), 100),
		Valid:        data[0] != 0xFE,
		data:         data,
	}, nil
}

func (sm SwitchMultilevelV1) String() string {
	if !sm.Valid {
		return fmt.Sprintf("current:unknown")
	}
	return fmt.Sprintf("current:%f", sm.CurrentValue)
}

// SwitchMultilevelV4 is sent by a zwave node to advertise the status of a multilevel device.
// A known issue is that nodes with multiple endpoints will send a SwichMultilevel report on
// all endpoint multilevel state changes but without encapsulating it in a MultiChannel message.
type SwitchMultilevelV4 struct {
	*report
	CurrentValue float64           `json:"current_value"` // The current state of the node
	TargetValue  float64           `json:"target_value"`  // The target state of the node
	Duration     commands.Duration `json:"duration"`      // The transition time between current and target values
	Valid        bool              `json:"valid"`         // If the current state in known or not. Some devices report "unknown" state
	TargetValid  bool              `json:"target_valid"`  // If the target state in known or not. Some devices report "unknown" state
	data         []byte
}

// NewSwitchMultilevelV4 creates a new SwitchMultilevelV1 report from binary data sent from a
// zwave node
func NewSwitchMultilevelV4(data []byte) (*SwitchMultilevelV4, error) {
	if len(data) != 3 {
		return nil, fmt.Errorf("Wrong length, expected 3 byte got %d", len(data))
	}

	return &SwitchMultilevelV4{
		CurrentValue: math.Min(float64(data[0]), 100),
		TargetValue:  math.Min(float64(data[1]), 100),
		Duration:     commands.Duration(data[2]),
		Valid:        data[0] != 0xFE,
		TargetValid:  data[1] != 0xFE,
		data:         data,
	}, nil
}

func (sm SwitchMultilevelV4) String() string {
	ret := fmt.Sprintf("current:%f ", sm.CurrentValue)
	if !sm.Valid {
		ret = fmt.Sprintf("current:unknown ")
	}

	if !sm.TargetValid {
		return ret + fmt.Sprintf("target:unknown duration:%s", sm.Duration)
	}
	return ret + fmt.Sprintf("target:%f duration:%s", sm.TargetValue, sm.Duration)
}
