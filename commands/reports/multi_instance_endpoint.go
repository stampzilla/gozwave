package reports

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type MultiChannelEndPoints struct {
	*report
	Dynamic   bool `json:"dynamic"`   // When true are the number of endpoints dynamic else its always the same
	Identical bool `json:"identical"` // When true are all endpoints identical (same command classes)
	Endpoints int  `json:"endpoints"` // Number of available endpoints
}

func NewMultiChannelEndPoints(data []byte) (*MultiChannelEndPoints, error) {
	if len(data) < 2 {
		return nil, fmt.Errorf("Failed to decode MultiChannelEndPoint: Wrong length, got %d expected 2: %x", len(data), data)
	}

	type aliasMultiChannelEndPoints struct {
		//Dynamic   bool
		//Identical bool
		//Spare3    bool
		//Spare4    bool
		//Spare5    bool
		//Spare6    bool
		//Spare7    bool
		//Spare8    bool
		bits      byte
		endpoints byte
	}

	ms := &aliasMultiChannelEndPoints{}

	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, ms)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode MultiChannelCmdEndPoint: %s (%x)", err, data)
	}

	ret := &MultiChannelEndPoints{}
	ret.Dynamic = (ms.bits & 0x80) != 0
	ret.Identical = (ms.bits & 0x40) != 0
	ret.Endpoints = int(ms.endpoints)

	return ret, nil
}

func (ms *MultiChannelEndPoints) String() string {
	return fmt.Sprintf("Dynamic: %t, Identical: %t, Endpoints: %d", ms.Dynamic, ms.Identical, ms.Endpoints)
}
