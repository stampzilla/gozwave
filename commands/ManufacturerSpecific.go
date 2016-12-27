package commands

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/Sirupsen/logrus"
)

type ManufacturerSpecificReport struct {
	*report
	Manufacturer string `json:"manufacturer"`
	Type         string `json:"type"`
	Id           string `json:"id"`
}

func NewManufacturerSpecificReport(data []byte) *ManufacturerSpecificReport {
	if len(data) != 6 {
		logrus.Warn("Failed to decode ManufacturerSpecific: Wrong length")
		return nil
	}

	type AliasCmdManufacturerSpecific struct {
		Manufacturer uint16
		Type         uint16
		Id           uint16
	}

	ms := &AliasCmdManufacturerSpecific{}

	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, ms)
	if err != nil {
		logrus.Warn("Failed to decode ManufacturerSpecific: %s", err)
		return nil
	}

	ret := &ManufacturerSpecificReport{}

	ret.Manufacturer = fmt.Sprintf("%04x", ms.Manufacturer)
	ret.Type = fmt.Sprintf("%04x", ms.Type)
	ret.Id = fmt.Sprintf("%04x", ms.Id)

	return ret
}

func (ms *ManufacturerSpecificReport) String() string {
	return fmt.Sprintf("%s:%s:%s", ms.Manufacturer, ms.Type, ms.Id)
}
