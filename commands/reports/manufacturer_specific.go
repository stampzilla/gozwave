package reports

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ManufacturerSpecific struct {
	*report
	Manufacturer string `json:"manufacturer"`
	Type         string `json:"type"`
	ID           string `json:"id"`
}

// NewManufacturerSpecific decodes and returns a manufacureSpecific report. This report is used to identify the manufacturer, model and version of a product.
func NewManufacturerSpecific(data []byte) (*ManufacturerSpecific, error) {
	if len(data) != 6 {
		return nil, fmt.Errorf("Failed to decode ManufacturerSpecific: Wrong length")
	}

	type aliasCmdManufacturerSpecific struct {
		manufacturer uint16
		t            uint16
		id           uint16
	}

	ms := &aliasCmdManufacturerSpecific{}

	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, ms)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode ManufacturerSpecific: %s", err)
	}

	ret := &ManufacturerSpecific{}

	ret.Manufacturer = fmt.Sprintf("%04x", ms.manufacturer)
	ret.Type = fmt.Sprintf("%04x", ms.t)
	ret.ID = fmt.Sprintf("%04x", ms.id)

	return ret, nil
}

func (ms *ManufacturerSpecific) String() string {
	return fmt.Sprintf("%s:%s:%s", ms.Manufacturer, ms.Type, ms.ID)
}
