package commands

type ZWaveCommand byte

// more found here: http://wiki.micasaverde.com/index.php/ZWave_Command_Classes
//COMMAND_CLASS_NO_OPERATION	0x00	0
//COMMAND_CLASS_BASIC	0x20	32
//COMMAND_CLASS_CONTROLLER_REPLICATION	0x21	33
//COMMAND_CLASS_APPLICATION_STATUS	0x22	34
//COMMAND_CLASS_ZIP_SERVICES	0x23	35
//COMMAND_CLASS_ZIP_SERVER	0x24	36
//COMMAND_CLASS_SWITCH_BINARY	0x25	37
//COMMAND_CLASS_SWITCH_MULTILEVEL	0x26	38
//COMMAND_CLASS_SWITCH_MULTILEVEL_V2	0x26	38
//COMMAND_CLASS_SWITCH_ALL	0x27	39
//COMMAND_CLASS_SWITCH_TOGGLE_BINARY	0x28	40
//COMMAND_CLASS_SWITCH_TOGGLE_MULTILEVEL	0x29	41
//COMMAND_CLASS_CHIMNEY_FAN	0x2A	42
//COMMAND_CLASS_SCENE_ACTIVATION	0x2B	43
//COMMAND_CLASS_SCENE_ACTUATOR_CONF	0x2C	44
//COMMAND_CLASS_SCENE_CONTROLLER_CONF	0x2D	45
//COMMAND_CLASS_ZIP_CLIENT	0x2E	46
//COMMAND_CLASS_ZIP_ADV_SERVICES	0x2F	47
//COMMAND_CLASS_SENSOR_BINARY	0x30	48
//COMMAND_CLASS_SENSOR_MULTILEVEL	0x31	49
//COMMAND_CLASS_SENSOR_MULTILEVEL_V2	0x31	49
//COMMAND_CLASS_METER	0x32	50
//COMMAND_CLASS_ZIP_ADV_SERVER	0x33	51
//COMMAND_CLASS_ZIP_ADV_CLIENT	0x34	52
//COMMAND_CLASS_METER_PULSE	0x35	53
//COMMAND_CLASS_METER_TBL_CONFIG	0x3C	60
//COMMAND_CLASS_METER_TBL_MONITOR	0x3D	61
//COMMAND_CLASS_METER_TBL_PUSH	0x3E	62
//COMMAND_CLASS_THERMOSTAT_HEATING	0x38	56
//COMMAND_CLASS_THERMOSTAT_MODE	0x40	64
//COMMAND_CLASS_THERMOSTAT_OPERATING_STATE	0x42	66
//COMMAND_CLASS_THERMOSTAT_SETPOINT	0x43	67
//COMMAND_CLASS_THERMOSTAT_FAN_MODE	0x44	68
//COMMAND_CLASS_THERMOSTAT_FAN_STATE	0x45	69
//COMMAND_CLASS_CLIMATE_CONTROL_SCHEDULE	0x46	70
//COMMAND_CLASS_THERMOSTAT_SETBACK	0x47	71
//COMMAND_CLASS_DOOR_LOCK_LOGGING	0x4C	76
//COMMAND_CLASS_SCHEDULE_ENTRY_LOCK	0x4E	78
//COMMAND_CLASS_BASIC_WINDOW_COVERING	0x50	80
//COMMAND_CLASS_MTP_WINDOW_COVERING	0x51	81
//COMMAND_CLASS_ASSOCIATION_GRP_INFO	0x59	89
//COMMAND_CLASS_DEVICE_RESET_LOCALLY	0x5A	90
//COMMAND_CLASS_CENTRAL_SCENE	0x5B	91
//COMMAND_CLASS_IP_ASSOCIATION	0x5C	92
//COMMAND_CLASS_ANTITHEFT	0x5D	93
//COMMAND_CLASS_ZWAVEPLUS_INFO	0x5E	94
//COMMAND_CLASS_MULTI_CHANNEL_V2	0x60	96
//COMMAND_CLASS_MULTI_INSTANCE	0x60	96
//COMMAND_CLASS_DOOR_LOCK	0x62	98
//COMMAND_CLASS_USER_CODE	0x63	99
//COMMAND_CLASS_BARRIER_OPERATOR	0x66	102
//COMMAND_CLASS_CONFIGURATION	0x70	112
//COMMAND_CLASS_CONFIGURATION_V2	0x70	112
//COMMAND_CLASS_ALARM	0x71	113
//COMMAND_CLASS_MANUFACTURER_SPECIFIC	0x72	114
//COMMAND_CLASS_POWERLEVEL	0x73	115
//COMMAND_CLASS_PROTECTION	0x75	117
//COMMAND_CLASS_PROTECTION_V2	0x75	117
//COMMAND_CLASS_LOCK	0x76	118
//COMMAND_CLASS_NODE_NAMING	0x77	119
//COMMAND_CLASS_FIRMWARE_UPDATE_MD	0x7A	122
//COMMAND_CLASS_GROUPING_NAME	0x7B	123
//COMMAND_CLASS_REMOTE_ASSOCIATION_ACTIVATE	0x7C	124
//COMMAND_CLASS_REMOTE_ASSOCIATION	0x7D	125
//COMMAND_CLASS_BATTERY	0x80	128
//COMMAND_CLASS_CLOCK	0x81	129
//COMMAND_CLASS_HAIL	0x82	130
//COMMAND_CLASS_WAKE_UP	0x84	132
//COMMAND_CLASS_WAKE_UP_V2	0x84	132
//COMMAND_CLASS_ASSOCIATION	0x85	133
//COMMAND_CLASS_ASSOCIATION_V2	0x85	133
//COMMAND_CLASS_VERSION	0x86	134
//COMMAND_CLASS_INDICATOR	0x87	135
//COMMAND_CLASS_PROPRIETARY	0x88	136
//COMMAND_CLASS_LANGUAGE	0x89	137
//COMMAND_CLASS_TIME	0x8A	138
//COMMAND_CLASS_TIME_PARAMETERS	0x8B	139
//COMMAND_CLASS_GEOGRAPHIC_LOCATION	0x8C	140
//COMMAND_CLASS_COMPOSITE	0x8D	141
//COMMAND_CLASS_MULTI_CHANNEL_ASSOCIATION_V2	0x8E	142
//COMMAND_CLASS_MULTI_INSTANCE_ASSOCIATION	0x8E	142
//COMMAND_CLASS_MULTI_CMD	0x8F	143
//COMMAND_CLASS_ENERGY_PRODUCTION	0x90	144
//COMMAND_CLASS_MANUFACTURER_PROPRIETARY	0x91	145
//COMMAND_CLASS_SCREEN_MD	0x92	146
//COMMAND_CLASS_SCREEN_MD_V2	0x92	146
//COMMAND_CLASS_SCREEN_ATTRIBUTES	0x93	147
//COMMAND_CLASS_SCREEN_ATTRIBUTES_V2	0x93	147
//COMMAND_CLASS_SIMPLE_AV_CONTROL	0x94	148
//COMMAND_CLASS_AV_CONTENT_DIRECTORY_MD	0x95	149
//COMMAND_CLASS_AV_RENDERER_STATUS	0x96	150
//COMMAND_CLASS_AV_CONTENT_SEARCH_MD	0x97	151
//COMMAND_CLASS_SECURITY	0x98	152
//COMMAND_CLASS_AV_TAGGING_MD	0x99	153
//COMMAND_CLASS_IP_CONFIGURATION	0x9A	154
//COMMAND_CLASS_ASSOCIATION_COMMAND_CONFIGURATION	0x9B	155
//COMMAND_CLASS_SENSOR_ALARM	0x9C	156
//COMMAND_CLASS_SILENCE_ALARM	0x9D	157
//COMMAND_CLASS_SENSOR_CONFIGURATION	0x9E	158
//COMMAND_CLASS_MARK	0xEF	239
//COMMAND_CLASS_NON_INTEROPERABLE	0xF0	240

const (
	Basic                           = 0x20
	ControllerReplication           = 0x21
	ApplicationStatus               = 0x22
	SwitchBinary                    = 0x25
	SwitchMultilevel                = 0x26
	SwitchAll                       = 0x27
	SwitchToggleBinary              = 0x28
	SwitchToggleMultilevel          = 0x29
	SceneActivation                 = 0x2B
	SceneActuatorConf               = 0x2C
	SensorBinary                    = 0x30
	SensorMultiLevel                = 0x31
	Meter                           = 0x32
	MeterPulse                      = 0x35
	ThermostatMode                  = 0x40
	ThermostatOperatingState        = 0x42
	ThermostatSetpoint              = 0x43
	ThermostatFanMode               = 0x44
	ThermostatFanState              = 0x45
	ClimateControlSchedule          = 0x46
	BasicWindowCovering             = 0x50
	AssociationGRPInfo              = 0x59
	MultiInstance                   = 0x60
	Configuration                   = 0x70
	Alarm                           = 0x71
	ManufacturerSpecific            = 0x72
	PowerLevel                      = 0x73
	Protection                      = 0x75
	Lock                            = 0x76
	NodeNaming                      = 0x77
	Battery                         = 0x80
	Clock                           = 0x81
	Hail                            = 0x82
	WakeUp                          = 0x84
	Association                     = 0x85
	Version                         = 0x86
	Indicator                       = 0x87
	Proprietary                     = 0x88
	Language                        = 0x89
	MultiInstanceAssociation        = 0x8e
	MultiCmd                        = 0x8F
	EnergyProduction                = 0x90
	ManufacturerProprietary         = 0x91
	AssociationCommandConfiguration = 0x9b
	Mark                            = 0xE
)

func (cmd ZWaveCommand) String() string {

	str := ""
	switch cmd {
	case 0x20:
		str = "Basic"
	case 0x21:
		str = "ControllerReplication"
	case 0x22:
		str = "ApplicationStatus"
	case 0x25:
		str = "SwitchBinary"
	case 0x26:
		str = "SwitchMultilevel"
	case 0x27:
		str = "SwitchAll"
	case 0x28:
		str = "SwitchToggleBinary"
	case 0x29:
		str = "SwitchToggleMultilevel"
	case 0x2B:
		str = "SceneActivation"
	case 0x2C:
		str = "SceneActuatorConf"
	case 0x30:
		str = "SensorBinary"
	case 0x31:
		str = "SensorMultiLevel"
	case 0x32:
		str = "Meter"
	case 0x35:
		str = "MeterPulse"
	case 0x40:
		str = "ThermostatMode"
	case 0x42:
		str = "ThermostatOperatingState"
	case 0x43:
		str = "ThermostatSetpoint"
	case 0x44:
		str = "ThermostatFanMode"
	case 0x45:
		str = "ThermostatFanState"
	case 0x46:
		str = "ClimateControlSchedule"
	case 0x50:
		str = "BasicWindowCovering"
	case 0x60:
		str = "MultiInstance"
	case 0x70:
		str = "Configuration"
	case 0x71:
		str = "Alarm"
	case 0x72:
		str = "ManufacturerSpecific"
	case 0x73:
		str = "PowerLevel"
	case 0x75:
		str = "Protection"
	case 0x76:
		str = "Lock"
	case 0x77:
		str = "NodeNaming"
	case 0x80:
		str = "Battery"
	case 0x81:
		str = "Clock"
	case 0x82:
		str = "Hail"
	case 0x84:
		str = "WakeUp"
	case 0x85:
		str = "Association"
	case 0x86:
		str = "Version"
	case 0x87:
		str = "Indicator"
	case 0x88:
		str = "Proprietary"
	case 0x89:
		str = "Language"
	case 0x8e:
		str = "MultiInstanceAssociation"
	case 0x8F:
		str = "MultiCmd"
	case 0x90:
		str = "EnergyProduction"
	case 0x91:
		str = "ManufacturerProprietary"
	case 0x9b:
		str = "AssociationCommandConfiguration"
	case 0xEF:
		str = "Mark"
	}
	return str

}
