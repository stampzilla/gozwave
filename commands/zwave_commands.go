package commands

type ZWaveCommand byte

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

func (self ZWaveCommand) String() string {

	str := ""
	switch self {
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
