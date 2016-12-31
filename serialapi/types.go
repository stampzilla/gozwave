package serialapi

type ZWaveFunction byte

const (
	None                           = 0x00
	DiscoveryNodes                 = 0x02
	SerialApiApplNodeInformation   = 0x03
	ApplicationCommandHandler      = 0x04
	GetControllerCapabilities      = 0x05
	SerialApiSetTimeouts           = 0x06
	SerialGetCapabilities          = 0x07
	SerialApiSoftReset             = 0x08
	SetRFReceiveMode               = 0x10
	SetSleepMode                   = 0x11
	SendNodeInformation            = 0x12
	SendData                       = 0x13
	SendDataMulti                  = 0x14
	GetVersion                     = 0x15
	SendDataAbort                  = 0x16
	RFPowerLevelSet                = 0x17
	SendDataMeta                   = 0x18
	MemoryGetId                    = 0x20
	MemoryGetByte                  = 0x21
	MemoryPutByte                  = 0x22
	MemoryGetBuffer                = 0x23
	MemoryPutBuffer                = 0x24
	ReadMemory                     = 0x25
	ClockSet                       = 0x30
	ClockGet                       = 0x31
	ClockCompare                   = 0x32
	RtcTimerCreate                 = 0x33
	RtcTimerRead                   = 0x34
	RtcTimerDelete                 = 0x35
	RtcTimerCall                   = 0x36
	GetNodeProtocolInfo            = 0x41
	SetDefault                     = 0x42
	ReplicationCommandComplete     = 0x44
	ReplicationSendData            = 0x45
	AssignReturnRoute              = 0x46
	DeleteReturnRoute              = 0x47
	RequestNodeNeighborUpdate      = 0x48
	ApplicationUpdate              = 0x49
	AddNodeToNetwork               = 0x4a
	RemoveNodeFromNetwork          = 0x4b
	CreateNewPrimary               = 0x4c
	ControllerChange               = 0x4d
	SetLearnMode                   = 0x50
	AssignSucReturnRoute           = 0x51
	EnableSuc                      = 0x52
	RequestNetworkUpdate           = 0x53
	SetSucNodeId                   = 0x54
	DeleteSucReturnRoute           = 0x55
	GetSucNodeId                   = 0x56
	SendSucId                      = 0x57
	RediscoveryNeeded              = 0x59
	RequestNodeInfo                = 0x60
	RemoveFailedNodeId             = 0x61
	IsFailedNode                   = 0x62
	ReplaceFailedNode              = 0x63
	TimerStart                     = 0x70
	TimerRestart                   = 0x71
	TimerCancel                    = 0x72
	TimerCall                      = 0x73
	GetRoutingTableLine            = 0x80
	GetTXCounter                   = 0x81
	ResetTXCounter                 = 0x82
	StoreNodeInfo                  = 0x83
	StoreHomeId                    = 0x84
	LockRouteResponse              = 0x90
	SendDataRouteDemo              = 0x91
	SerialApiTest                  = 0x95
	SerialApiSlaveNodeInfo         = 0xa0
	ApplicationSlaveCommandHandler = 0xa1
	SendSlaveNodeInfo              = 0xa2
	SendSlaveData                  = 0xa3
	SetSlaveLearnMode              = 0xa4
	GetVirtualNodes                = 0xa5
	IsVirtualNode                  = 0xa6
	SetPromiscuousMode             = 0xd0
)

func (b ZWaveFunction) String() string {
	str := ""
	switch b {
	case 0x00:
		str = "None"
	case 0x02:
		str = "DiscoveryNodes"
	case 0x03:
		str = "SerialApiApplNodeInformation"
	case 0x04:
		str = "ApplicationCommandHandler"
	case 0x05:
		str = "GetControllerCapabilities"
	case 0x06:
		str = "SerialApiSetTimeouts"
	case 0x07:
		str = "SerialGetCapabilities"
	case 0x08:
		str = "SerialApiSoftReset"
	case 0x10:
		str = "SetRFReceiveMode"
	case 0x11:
		str = "SetSleepMode"
	case 0x12:
		str = "SendNodeInformation"
	case 0x13:
		str = "SendData"
	case 0x14:
		str = "SendDataMulti"
	case 0x15:
		str = "GetVersion"
	case 0x16:
		str = "SendDataAbort"
	case 0x17:
		str = "RFPowerLevelSet"
	case 0x18:
		str = "SendDataMeta"
	case 0x20:
		str = "MemoryGetId"
	case 0x21:
		str = "MemoryGetByte"
	case 0x22:
		str = "MemoryPutByte"
	case 0x23:
		str = "MemoryGetBuffer"
	case 0x24:
		str = "MemoryPutBuffer"
	case 0x25:
		str = "ReadMemory"
	case 0x30:
		str = "ClockSet"
	case 0x31:
		str = "ClockGet"
	case 0x32:
		str = "ClockCompare"
	case 0x33:
		str = "RtcTimerCreate"
	case 0x34:
		str = "RtcTimerRead"
	case 0x35:
		str = "RtcTimerDelete"
	case 0x36:
		str = "RtcTimerCall"
	case 0x41:
		str = "GetNodeProtocolInfo"
	case 0x42:
		str = "SetDefault"
	case 0x44:
		str = "ReplicationCommandComplete"
	case 0x45:
		str = "ReplicationSendData"
	case 0x46:
		str = "AssignReturnRoute"
	case 0x47:
		str = "DeleteReturnRoute"
	case 0x48:
		str = "RequestNodeNeighborUpdate"
	case 0x49:
		str = "ApplicationUpdate"
	case 0x4a:
		str = "AddNodeToNetwork"
	case 0x4b:
		str = "RemoveNodeFromNetwork"
	case 0x4c:
		str = "CreateNewPrimary"
	case 0x4d:
		str = "ControllerChange"
	case 0x50:
		str = "SetLearnMode"
	case 0x51:
		str = "AssignSucReturnRoute"
	case 0x52:
		str = "EnableSuc"
	case 0x53:
		str = "RequestNetworkUpdate"
	case 0x54:
		str = "SetSucNodeId"
	case 0x55:
		str = "DeleteSucReturnRoute"
	case 0x56:
		str = "GetSucNodeId"
	case 0x57:
		str = "SendSucId"
	case 0x59:
		str = "RediscoveryNeeded"
	case 0x60:
		str = "RequestNodeInfo"
	case 0x61:
		str = "RemoveFailedNodeId"
	case 0x62:
		str = "IsFailedNode"
	case 0x63:
		str = "ReplaceFailedNode"
	case 0x70:
		str = "TimerStart"
	case 0x71:
		str = "TimerRestart"
	case 0x72:
		str = "TimerCancel"
	case 0x73:
		str = "TimerCall"
	case 0x80:
		str = "GetRoutingTableLine"
	case 0x81:
		str = "GetTXCounter"
	case 0x82:
		str = "ResetTXCounter"
	case 0x83:
		str = "StoreNodeInfo"
	case 0x84:
		str = "StoreHomeId"
	case 0x90:
		str = "LockRouteResponse"
	case 0x91:
		str = "SendDataRouteDemo"
	case 0x95:
		str = "SerialApiTest"
	case 0xa0:
		str = "SerialApiSlaveNodeInfo"
	case 0xa1:
		str = "ApplicationSlaveCommandHandler"
	case 0xa2:
		str = "SendSlaveNodeInfo"
	case 0xa3:
		str = "SendSlaveData"
	case 0xa4:
		str = "SetSlaveLearnMode"
	case 0xa5:
		str = "GetVirtualNodes"
	case 0xa6:
		str = "IsVirtualNode"
	case 0xd0:
		str = "SetPromiscuousMode"
	}

	return str
}
