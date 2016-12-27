package commands

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/Sirupsen/logrus"
)

type MultiChannelCmdEndPointReport struct {
	*report
	Dynamic   bool `json:"dynamic"`
	Identical bool `json:"identical"`
	Endpoints int
}

func NewMultiChannelCmdEndPointReport(data []byte) *MultiChannelCmdEndPointReport {
	if len(data) < 2 {

		logrus.Warnf("Failed to decode MultiChannelCmdEndPointReport: Wrong length, got %d expected 2: %x", len(data), data)
		return nil
	}

	type AliasMultiChannelCmdEndPointReport struct {
		//Dynamic   bool
		//Identical bool
		//Spare3    bool
		//Spare4    bool
		//Spare5    bool
		//Spare6    bool
		//Spare7    bool
		//Spare8    bool
		Bits      byte
		Endpoints byte
	}

	ms := &AliasMultiChannelCmdEndPointReport{}

	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, ms)
	if err != nil {
		logrus.Warn("Failed to decode MultiChannelCmdEndPointReport: %s (%x)", err, data)
		return nil
	}

	ret := &MultiChannelCmdEndPointReport{}
	ret.Dynamic = (ms.Bits & 0x80) != 0
	ret.Identical = (ms.Bits & 0x40) != 0
	ret.Endpoints = int(ms.Endpoints)

	return ret
}

func (ms *MultiChannelCmdEndPointReport) String() string {
	return fmt.Sprintf("Dynamic: %b, Identical: %b, Endpoints: %d", ms.Dynamic, ms.Identical, ms.Endpoints)
}
