package commands

import (
	"fmt"
	"time"
)

// Duration used to describe time between 0-127 seconds and 1-127 minutes in a byte
type Duration byte

// Duration converts zwave duration type into time.Duration
func (d Duration) Duration() (time.Duration, error) {
	switch {
	case byte(d) == 0x00:
		return time.Duration(0), nil
	case byte(d) <= 0x7F:
		return time.Duration(d) * time.Second, nil
	case d <= 0xFD:
		return time.Duration(d-0x7F) * time.Minute, nil
	}
	return time.Duration(0), fmt.Errorf("Factory default duration")
}

func (d Duration) String() string {
	t, err := d.Duration()

	if err != nil {
		return err.Error()
	}

	return t.String()
}
