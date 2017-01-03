package reports

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// ManufacturerSpecific reports are used to identify the manufacturer, model and version of a product.
type ManufacturerSpecific struct {
	*report
	Manufacturer string `json:"manufacturer"`
	Type         string `json:"type"`
	ID           string `json:"id"`
}

// NewManufacturerSpecific decodes and returns a manufacureSpecific report.
func NewManufacturerSpecific(data []byte) (*ManufacturerSpecific, error) {
	if len(data) != 6 {
		return nil, fmt.Errorf("Failed to decode ManufacturerSpecific: Wrong length")
	}

	type aliasCmdManufacturerSpecific struct {
		Manufacturer uint16
		Type         uint16
		ID           uint16
	}

	ms := &aliasCmdManufacturerSpecific{}

	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, ms)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode ManufacturerSpecific: %s", err)
	}

	ret := &ManufacturerSpecific{}

	ret.Manufacturer = fmt.Sprintf("%04x", ms.Manufacturer)
	ret.Type = fmt.Sprintf("%04x", ms.Type)
	ret.ID = fmt.Sprintf("%04x", ms.ID)

	return ret, nil
}

func (ms *ManufacturerSpecific) String() string {
	return fmt.Sprintf("manufacturerSpecific manufacturer:%s type:%s id:%s (%s:%s:%s)", ms.Manufacturer, ms.Type, ms.ID, ms.Manufacturer, ms.Type, ms.ID)
}
