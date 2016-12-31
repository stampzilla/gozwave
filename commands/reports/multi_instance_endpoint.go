package reports

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type MultiChannelEndpoints struct {
	*report
	Dynamic   bool `json:"dynamic"`   // When true are the number of endpoints dynamic else its always the same
	Identical bool `json:"identical"` // When true are all endpoints identical (same command classes)
	Endpoints int  `json:"endpoints"` // Number of available endpoints
}

func NewMultiChannelEndpoints(data []byte) (*MultiChannelEndpoints, error) {
	if len(data) < 2 {
		return nil, fmt.Errorf("wrong length, got %d expected 2 (data: %x)", len(data), data)
	}

	type aliasMultiChannelEndpoints struct {
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

	ms := &aliasMultiChannelEndpoints{}

	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, ms)
	if err != nil {
		return nil, fmt.Errorf("failed binary decode: %s (data: %x)", err, data)
	}

	ret := &MultiChannelEndpoints{}
	ret.Dynamic = (ms.Bits & 0x80) != 0
	ret.Identical = (ms.Bits & 0x40) != 0
	ret.Endpoints = int(ms.Endpoints)

	return ret, nil
}

func (ms *MultiChannelEndpoints) String() string {
	return fmt.Sprintf("multiChannelEndpoints dynamic:%t identical:%t endpoints:%d", ms.Dynamic, ms.Identical, ms.Endpoints)
}
