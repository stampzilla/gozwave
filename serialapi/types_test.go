package serialapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {

	m := map[byte]string{
		0x00: "None",
		0x02: "DiscoveryNodes",
		0x03: "SerialApiApplNodeInformation",
		0x04: "ApplicationCommandHandler",
		0x05: "GetControllerCapabilities",
		0x06: "SerialApiSetTimeouts",
		0x07: "SerialGetCapabilities",
		0x08: "SerialApiSoftReset",
		0x10: "SetRFReceiveMode",
		0x11: "SetSleepMode",
		0x12: "SendNodeInformation",
		0x13: "SendData",
		0x14: "SendDataMulti",
		0x15: "GetVersion",
		0x16: "SendDataAbort",
		0x17: "RFPowerLevelSet",
		0x18: "SendDataMeta",
		0x20: "MemoryGetId",
		0x21: "MemoryGetByte",
		0x22: "MemoryPutByte",
		0x23: "MemoryGetBuffer",
		0x24: "MemoryPutBuffer",
		0x25: "ReadMemory",
		0x30: "ClockSet",
		0x31: "ClockGet",
		0x32: "ClockCompare",
		0x33: "RtcTimerCreate",
		0x34: "RtcTimerRead",
		0x35: "RtcTimerDelete",
		0x36: "RtcTimerCall",
		0x41: "GetNodeProtocolInfo",
		0x42: "SetDefault",
		0x44: "ReplicationCommandComplete",
		0x45: "ReplicationSendData",
		0x46: "AssignReturnRoute",
		0x47: "DeleteReturnRoute",
		0x48: "RequestNodeNeighborUpdate",
		0x49: "ApplicationUpdate",
		0x4a: "AddNodeToNetwork",
		0x4b: "RemoveNodeFromNetwork",
		0x4c: "CreateNewPrimary",
		0x4d: "ControllerChange",
		0x50: "SetLearnMode",
		0x51: "AssignSucReturnRoute",
		0x52: "EnableSuc",
		0x53: "RequestNetworkUpdate",
		0x54: "SetSucNodeId",
		0x55: "DeleteSucReturnRoute",
		0x56: "GetSucNodeId",
		0x57: "SendSucId",
		0x59: "RediscoveryNeeded",
		0x60: "RequestNodeInfo",
		0x61: "RemoveFailedNodeId",
		0x62: "IsFailedNode",
		0x63: "ReplaceFailedNode",
		0x70: "TimerStart",
		0x71: "TimerRestart",
		0x72: "TimerCancel",
		0x73: "TimerCall",
		0x80: "GetRoutingTableLine",
		0x81: "GetTXCounter",
		0x82: "ResetTXCounter",
		0x83: "StoreNodeInfo",
		0x84: "StoreHomeId",
		0x90: "LockRouteResponse",
		0x91: "SendDataRouteDemo",
		0x95: "SerialApiTest",
		0xa0: "SerialApiSlaveNodeInfo",
		0xa1: "ApplicationSlaveCommandHandler",
		0xa2: "SendSlaveNodeInfo",
		0xa3: "SendSlaveData",
		0xa4: "SetSlaveLearnMode",
		0xa5: "GetVirtualNodes",
		0xa6: "IsVirtualNode",
		0xd0: "SetPromiscuousMode",
	}

	for k, v := range m {
		assert.Equal(t, v, ZWaveFunction(k).String())
	}

}
