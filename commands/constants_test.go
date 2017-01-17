package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZwaveCommand(t *testing.T) {
	m := map[byte]string{
	0x20: "Basic",
	0x21: "ControllerReplication",
	0x22: "ApplicationStatus",
	0x25: "SwitchBinary",
	0x26: "SwitchMultilevel",
	0x27: "SwitchAll",
	0x28: "SwitchToggleBinary",
	0x29: "SwitchToggleMultilevel",
	0x2B: "SceneActivation",
	0x2C: "SceneActuatorConf",
	0x30: "SensorBinary",
	0x31: "SensorMultiLevel",
	0x32: "Meter",
	0x35: "MeterPulse",
	0x40: "ThermostatMode",
	0x42: "ThermostatOperatingState",
	0x43: "ThermostatSetpoint",
	0x44: "ThermostatFanMode",
	0x45: "ThermostatFanState",
	0x46: "ClimateControlSchedule",
	0x50: "BasicWindowCovering",
	0x60: "MultiInstance",
	0x70: "Configuration",
	0x71: "Alarm",
	0x72: "ManufacturerSpecific",
	0x73: "PowerLevel",
	0x75: "Protection",
	0x76: "Lock",
	0x77: "NodeNaming",
	0x80: "Battery",
	0x81: "Clock",
	0x82: "Hail",
	0x84: "WakeUp",
	0x85: "Association",
	0x86: "Version",
	0x87: "Indicator",
	0x88: "Proprietary",
	0x89: "Language",
	0x8e: "MultiInstanceAssociation",
	0x8F: "MultiCmd",
	0x90: "EnergyProduction",
	0x91: "ManufacturerProprietary",
	0x9b: "AssociationCommandConfiguration",
	0xEF: "Mark",
	}

	for k, v := range m {
		assert.Equal(t, v, ZWaveCommand(k).String())
	}

}
