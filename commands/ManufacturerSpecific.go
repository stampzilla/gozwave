package commands

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/Sirupsen/logrus"
)

type CmdManufacturerSpecific struct {
	Manufacturer uint16
	Type         uint16
	Id           uint16
}

func NewCmdManufacturerSpecific(data []byte) *CmdManufacturerSpecific {
	if len(data) != 6 {

		logrus.Warn("Failed to decode ManufacturerSpecific: Wrong length")
		return nil
	}

	ms := &CmdManufacturerSpecific{}

	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, ms)
	if err != nil {
		logrus.Warn("Failed to decode ManufacturerSpecific: %s", err)
		return nil
	}

	return ms
}

func (ms *CmdManufacturerSpecific) String() string {
	return fmt.Sprintf("%04x:%04x:%04x", ms.Manufacturer, ms.Type, ms.Id)
}
