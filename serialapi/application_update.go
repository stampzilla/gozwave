package serialapi

import (
	"fmt"

	"github.com/Sirupsen/logrus"
)

const (
	UPDATE_STATE_SUC_ID               = 0x10
	UPDATE_STATE_DELETE_DONE          = 0x20
	UPDATE_STATE_NEW_ID_ASSIGNED      = 0x40
	UPDATE_STATE_ROUTING_PENDING      = 0x80
	UPDATE_STATE_NODE_INFO_REQ_FAILED = 0x81
	UPDATE_STATE_NODE_INFO_REQ_DONE   = 0x82
	UPDATE_STATE_NODE_INFO_RECEIVED   = 0x84
)

type FuncApplicationUpdate struct {
	Event        byte
	Supported    []byte
	Controllable []byte
}

func NewApplicationUpdate(data []byte) (*FuncApplicationUpdate, error) {
	au := &FuncApplicationUpdate{}

	switch data[0] { // Report type
	case UPDATE_STATE_SUC_ID:
	case UPDATE_STATE_DELETE_DONE:
	case UPDATE_STATE_NEW_ID_ASSIGNED:
	case UPDATE_STATE_ROUTING_PENDING:
	case UPDATE_STATE_NODE_INFO_REQ_FAILED:

	case UPDATE_STATE_NODE_INFO_REQ_DONE:
	case UPDATE_STATE_NODE_INFO_RECEIVED:
		afterMark := false

		logrus.Infof("DATA 0x%X", data)
		for _, v := range data[1:] {
			if v == 0xEF { // End of supported command classes. The rest are those that can be controlled by the device.
				afterMark = true
				continue
			}

			if afterMark {
				// controllable by the device
				au.Controllable = append(au.Controllable, v)
				continue
			}

			// supported by the device
			au.Supported = append(au.Supported, v)
		}
	default:
		return nil, fmt.Errorf("Unknow response type 0x%X", data[0])
	}
	logrus.Warnf("AU: %#v", au)
	return au, nil
}
