package database
//import (
	//"github.com/stampzilla/gozwave/commands"
//)

var definitions = map[Definition][]*CommandClass{
	Definition{
		Generic: protocol.GENERIC_TYPE_AV_CONTROL_POINT,
		Specific: protocol.SPECIFIC_TYPE_NOT_USED,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_DISPLAY,
		Specific: protocol.SPECIFIC_TYPE_SIMPLE_DISPLAY,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_ENTRY_CONTROL,
		Specific: protocol.SPECIFIC_TYPE_SECURE _BARRIER_ADDON,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_ENTRY_CONTROL,
		Specific: protocol.SPECIFIC_TYPE_SECURE_BARRIER_CLOSE_ONLY,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_ENTRY_CONTROL,
		Specific: protocol.SPECIFIC_TYPE_SECURE_BARRIER_OPEN_ONLY,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_ENTRY_CONTROL,
		Specific: protocol.SPECIFIC_TYPE_SECURE_DOOR,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_ENTRY_CONTROL,
		Specific: protocol.SPECIFIC_TYPE_SECURE_GATE,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_ENTRY_CONTROL,
		Specific: protocol.SPECIFIC_TYPE_SECURE_KEYPAD,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_ENTRY_CONTROL,
		Specific: protocol.SPECIFIC_TYPE_SECURE_KEYPAD_DOOR_LOCK,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_ENTRY_CONTROL,
		Specific: protocol.SPECIFIC_TYPE_SECURE_LOCKBOX,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_GENERIC_CONTROLLER,
		Specific: protocol.SPECIFIC_TYPE_PORTABLE_REMOTE_CONTROLLER,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_GENERIC_CONTROLLER,
		Specific: protocol.SPECIFIC_TYPE_REMOTE_CONTROL_AV,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_GENERIC_CONTROLLER,
		Specific: protocol.SPECIFIC_TYPE_REMOTE_CONTROL_SIMPLE,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x5b,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_METER,
		Specific: protocol.SPECIFIC_TYPE_SIMPLE_METER,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x56,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x32,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_METER,
		Specific: protocol.SPECIFIC_TYPE_WHOLE_HOME_METER_SIMPLE,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x56,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x32,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_REPEATER_SLAVE,
		Specific: protocol.SPECIFIC_TYPE_REPEATER_SLAVE,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_SENSOR_MULTILEVEL,
		Specific: protocol.SPECIFIC_TYPE_ROUTING_MULTILEVEL_SENSOR,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x31,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_SENSOR_NOTIFICATION,
		Specific: protocol.SPECIFIC_TYPE_NOTIFICATION_SENSOR,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x71,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_STATIC_CONTROLLER,
		Specific: protocol.SPECIFIC_TYPE_GATEWAY,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x22,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x56,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x98,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_STATIC_CONTROLLER,
		Specific: protocol.SPECIFIC_TYPE_PC_CONTROLLER,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x22,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x56,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x98,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_STATIC_CONTROLLER,
		Specific: protocol.SPECIFIC_TYPE_SET_TOP_BOX,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x22,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x56,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x98,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_STATIC_CONTROLLER,
		Specific: protocol.SPECIFIC_TYPE_SUB_SYSTEM_CONTROLLER,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x22,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x56,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_STATIC_CONTROLLER,
		Specific: protocol.SPECIFIC_TYPE_TV,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x22,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x56,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x98,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_SWITCH_BINARY,
		Specific: protocol.SPECIFIC_TYPE_IRRIGATION_CONTROLLER,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x20,
		},
		&CommandClass{
			ID: 0x25,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_SWITCH_BINARY,
		Specific: protocol.SPECIFIC_TYPE_POWER_STRIP,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x20,
		},
		&CommandClass{
			ID: 0x25,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x60,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_SWITCH_BINARY,
		Specific: protocol.SPECIFIC_TYPE_POWER_SWITCH_BINARY,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x20,
		},
		&CommandClass{
			ID: 0x25,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_SWITCH_BINARY,
		Specific: protocol.SPECIFIC_TYPE_SIREN,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x20,
		},
		&CommandClass{
			ID: 0x25,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_SWITCH_BINARY,
		Specific: protocol.SPECIFIC_TYPE_VALVE_OPEN_CLOSE,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x20,
		},
		&CommandClass{
			ID: 0x25,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_SWITCH_MULTILEVEL,
		Specific: protocol.SPECIFIC_TYPE_CLASS_A_MOTOR_CONTROL,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x20,
		},
		&CommandClass{
			ID: 0x25,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x26,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_SWITCH_MULTILEVEL,
		Specific: protocol.SPECIFIC_TYPE_CLASS_B_MOTOR_CONTROL,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x20,
		},
		&CommandClass{
			ID: 0x25,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x26,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_SWITCH_MULTILEVEL,
		Specific: protocol.SPECIFIC_TYPE_FAN_SWITCH,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x20,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x26,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_SWITCH_MULTILEVEL,
		Specific: protocol.SPECIFIC_TYPE_POWER_SWITCH_MULTILEVEL,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x20,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x26,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_THERMOSTAT,
		Specific: protocol.SPECIFIC_TYPE_SETBACK_THERMOSTAT,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x20,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_THERMOSTAT,
		Specific: protocol.SPECIFIC_TYPE_THERMOSTAT_GENERAL_V2,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x20,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}
	Definition{
		Generic: protocol.GENERIC_TYPE_WALL_CONTROLLER,
		Specific: protocol.SPECIFIC_TYPE_BASIC_WALL_CONTROLLER,
	}: []*CommandClass{
		&CommandClass{
			ID: 0x5e,
		},
		&CommandClass{
			ID: 0x85,
		},
		&CommandClass{
			ID: 0x59,
		},
		&CommandClass{
			ID: 0x5b,
		},
		&CommandClass{
			ID: 0x5a,
		},
		&CommandClass{
			ID: 0x72,
		},
		&CommandClass{
			ID: 0x73,
		},
		&CommandClass{
			ID: 0x86,
		},
	},
}

