package gozwave

/************ Basic Device Class identifiers **************/
const BASIC_TYPE_CONTROLLER = 0x01 /*Node is a portable
controller */
const BASIC_TYPE_ROUTING_SLAVE = 0x04     /*Node is a slave with routing capabilities*/
const BASIC_TYPE_SLAVE = 0x03             /*Node is a slave*/
const BASIC_TYPE_STATIC_CONTROLLER = 0x02 /*Node is a static controller*/

/***** Generic and Specific Device Class identifiers ******/

/* Device class - Generic Controller */
const GENERIC_TYPE_GENERIC_CONTROLLER = 0x01           /*Remote Controller*/
const SPECIFIC_TYPE_GENERIC_CONTROLLER_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_PORTABLE_REMOTE_CONTROLLER = 0x01  /*Remote Control (Multi Purpose) Device Type*/
const SPECIFIC_TYPE_PORTABLE_SCENE_CONTROLLER = 0x02   /*Portable Scene Controller*/
const SPECIFIC_TYPE_PORTABLE_INSTALLER_TOOL = 0x03
const SPECIFIC_TYPE_REMOTE_CONTROL_AV = 0x04     /*Remote Control (AV) Device Type*/
const SPECIFIC_TYPE_REMOTE_CONTROL_SIMPLE = 0x06 /*Remote Control (Simple) Device Type*/

/* Device class - Static Controller */
const GENERIC_TYPE_STATIC_CONTROLLER = 0x02           /*Static Controller*/
const SPECIFIC_TYPE_STATIC_CONTROLLER_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_PC_CONTROLLER = 0x01              /*Central Controller Device Type*/
const SPECIFIC_TYPE_SCENE_CONTROLLER = 0x02           /*Scene Controller*/
const SPECIFIC_TYPE_STATIC_INSTALLER_TOOL = 0x03
const SPECIFIC_TYPE_SET_TOP_BOX = 0x04           /*Set Top Box Device Type*/
const SPECIFIC_TYPE_SUB_SYSTEM_CONTROLLER = 0x05 /*Sub System Controller Device Type*/
const SPECIFIC_TYPE_TV = 0x06                    /*TV Device Type*/
const SPECIFIC_TYPE_GATEWAY = 0x07               /*Gateway Device Type*/

/* Device class - Av Control Point */
const GENERIC_TYPE_AV_CONTROL_POINT = 0x03           /*AV Control Point*/
const SPECIFIC_TYPE_AV_CONTROL_POINT_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_DOORBELL = 0x12
const SPECIFIC_TYPE_SATELLITE_RECEIVER = 0x04    /*Satellite Receiver*/
const SPECIFIC_TYPE_SATELLITE_RECEIVER_V2 = 0x11 /*Satellite Receiver V2*/

/* Device class - Display */
const GENERIC_TYPE_DISPLAY = 0x04
const SPECIFIC_TYPE_DISPLAY_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_SIMPLE_DISPLAY = 0x01   /*Display (simple) Device Type*/

/* Device class - Network Extender */
const GENERIC_TYPE_NETWORK_EXTENDER = 0x05           /*Network Extender Generic Device Class*/
const SPECIFIC_TYPE_NETWORK_EXTENDER_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_SECURE_EXTENDER = 0x01           /*Specific Device Secure Extender*/

/* Device class - Appliance */
const GENERIC_TYPE_APPLIANCE = 0x06
const SPECIFIC_TYPE_APPLIANCE_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_GENERAL_APPLIANCE = 0x01
const SPECIFIC_TYPE_KITCHEN_APPLIANCE = 0x02
const SPECIFIC_TYPE_LAUNDRY_APPLIANCE = 0x03

/* Device class - Sensor Notification */
const GENERIC_TYPE_SENSOR_NOTIFICATION = 0x07
const SPECIFIC_TYPE_SENSOR_NOTIFICATION_NOT_USED = 0x00 /*Specific Device Class not used*/
const SPECIFIC_TYPE_NOTIFICATION_SENSOR = 0x01

/* Device class - Thermostat */
const GENERIC_TYPE_THERMOSTAT = 0x08                   /*Thermostat*/
const SPECIFIC_TYPE_THERMOSTAT_NOT_USED = 0x00         /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_SETBACK_SCHEDULE_THERMOSTAT = 0x03 /*Setback Schedule Thermostat*/
const SPECIFIC_TYPE_SETBACK_THERMOSTAT = 0x05          /*Thermostat (Setback) Device Type*/
const SPECIFIC_TYPE_SETPOINT_THERMOSTAT = 0x04
const SPECIFIC_TYPE_THERMOSTAT_GENERAL = 0x02    /*Thermostat General*/
const SPECIFIC_TYPE_THERMOSTAT_GENERAL_V2 = 0x06 /*Thermostat (HVAC) Device Type*/
const SPECIFIC_TYPE_THERMOSTAT_HEATING = 0x01    /*Thermostat Heating*/

/* Device class - Window covering */
const GENERIC_TYPE_WINDOW_COVERING = 0x09           /*Window Covering*/
const SPECIFIC_TYPE_WINDOW_COVERING_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_SIMPLE_WINDOW_COVERING = 0x01   /*Simple Window Covering Control*/

/* Device class - Repeater Slave */
const GENERIC_TYPE_REPEATER_SLAVE = 0x0F           /*Repeater Slave*/
const SPECIFIC_TYPE_REPEATER_SLAVE_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_REPEATER_SLAVE = 0x01          /*Basic Repeater Slave*/
const SPECIFIC_TYPE_VIRTUAL_NODE = 0x02

/* Device class - Switch Binary */
const GENERIC_TYPE_SWITCH_BINARY = 0x10           /*Binary Switch*/
const SPECIFIC_TYPE_SWITCH_BINARY_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_POWER_SWITCH_BINARY = 0x01    /*On/Off Power Switch Device Type*/
const SPECIFIC_TYPE_SCENE_SWITCH_BINARY = 0x03    /*Binary Scene Switch*/
const SPECIFIC_TYPE_POWER_STRIP = 0x04            /*Power Strip Device Type*/
const SPECIFIC_TYPE_SIREN = 0x05                  /*Siren Device Type*/
const SPECIFIC_TYPE_VALVE_OPEN_CLOSE = 0x06       /*Valve (open/close) Device Type*/
const SPECIFIC_TYPE_COLOR_TUNABLE_BINARY = 0x02
const SPECIFIC_TYPE_IRRIGATION_CONTROLLER = 0x07

/* Device class - Switch Multilevel */
const GENERIC_TYPE_SWITCH_MULTILEVEL = 0x11           /*Multilevel Switch*/
const SPECIFIC_TYPE_SWITCH_MULTILEVEL_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_CLASS_A_MOTOR_CONTROL = 0x05      /*Window Covering No Position/Endpoint Device Type*/
const SPECIFIC_TYPE_CLASS_B_MOTOR_CONTROL = 0x06      /*Window Covering Endpoint Aware Device Type*/
const SPECIFIC_TYPE_CLASS_C_MOTOR_CONTROL = 0x07      /*Window Covering Position/Endpoint Aware Device Type*/
const SPECIFIC_TYPE_MOTOR_MULTIPOSITION = 0x03        /*Multiposition Motor*/
const SPECIFIC_TYPE_POWER_SWITCH_MULTILEVEL = 0x01    /*Light Dimmer Switch Device Type*/
const SPECIFIC_TYPE_SCENE_SWITCH_MULTILEVEL = 0x04    /*Multilevel Scene Switch*/
const SPECIFIC_TYPE_FAN_SWITCH = 0x08                 /*Fan Switch Device Type*/
const SPECIFIC_TYPE_COLOR_TUNABLE_MULTILEVEL = 0x02   /* Device class Switch Remote */

/* Device class - Remote Switch */
const GENERIC_TYPE_SWITCH_REMOTE = 0x12                    /*Remote Switch*/
const SPECIFIC_TYPE_SWITCH_REMOTE_NOT_USED = 0x00          /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_SWITCH_REMOTE_BINARY = 0x01            /*Binary Remote Switch*/
const SPECIFIC_TYPE_SWITCH_REMOTE_MULTILEVEL = 0x02        /*Multilevel Remote Switch*/
const SPECIFIC_TYPE_SWITCH_REMOTE_TOGGLE_BINARY = 0x03     /*Binary Toggle Remote Switch*/
const SPECIFIC_TYPE_SWITCH_REMOTE_TOGGLE_MULTILEVEL = 0x04 /*Multilevel Toggle Remote Switch*/

/* Device class - Switch Toggle */
const GENERIC_TYPE_SWITCH_TOGGLE = 0x13             /*Toggle Switch*/
const SPECIFIC_TYPE_SWITCH_TOGGLE_NOT_USED = 0x00   /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_SWITCH_TOGGLE_BINARY = 0x01     /*Binary Toggle Switch*/
const SPECIFIC_TYPE_SWITCH_TOGGLE_MULTILEVEL = 0x02 /*Multilevel Toggle Switch*/

/* Device class - Zip Node */
const GENERIC_TYPE_ZIP_NODE = 0x15
const SPECIFIC_TYPE_ZIP_NODE_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_ZIP_ADV_NODE = 0x02
const SPECIFIC_TYPE_ZIP_TUN_NODE = 0x01 /* Device class Wall Controller */

/* Device class - Ventilation */
const GENERIC_TYPE_VENTILATION = 0x16
const SPECIFIC_TYPE_VENTILATION_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_RESIDENTIAL_HRV = 0x01      /* Device class Window Covering */

/* Device class - Security Panel */
const GENERIC_TYPE_SECURITY_PANEL = 0x17
const SPECIFIC_TYPE_SECURITY_PANEL_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_ZONED_SECURITY_PANEL = 0x01    /* Device class Semi Interoperable */

/* Device class - Wall controller */
const GENERIC_TYPE_WALL_CONTROLLER = 0x18
const SPECIFIC_TYPE_WALL_CONTROLLER_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_BASIC_WALL_CONTROLLER = 0x01    /*Wall Controller Device Type*/

/* Device class - Sensor Binary */
const GENERIC_TYPE_SENSOR_BINARY = 0x20           /*Binary Sensor*/
const SPECIFIC_TYPE_SENSOR_BINARY_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_ROUTING_SENSOR_BINARY = 0x01  /*Routing Binary Sensor*/

/* Device class - Sensor Multilevel */
const GENERIC_TYPE_SENSOR_MULTILEVEL = 0x21           /*Multilevel Sensor*/
const SPECIFIC_TYPE_SENSOR_MULTILEVEL_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_ROUTING_SENSOR_MULTILEVEL = 0x01  /*Sensor (Multilevel) Device Type*/
const SPECIFIC_TYPE_CHIMNEY_FAN = 0x02                /* Device class Static Controller */

/* Device class - Meter Pulse */
const GENERIC_TYPE_METER_PULSE = 0x30           /*Pulse Meter*/
const SPECIFIC_TYPE_METER_PULSE_NOT_USED = 0x00 /*Specific Device Class Not Used*/

/* Device class - Meter */
const GENERIC_TYPE_METER = 0x31                    /*Meter*/
const SPECIFIC_TYPE_METER_NOT_USED = 0x00          /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_SIMPLE_METER = 0x01            /*Sub Energy Meter Device Type*/
const SPECIFIC_TYPE_ADV_ENERGY_CONTROL = 0x02      /*Whole Home Energy Meter (Advanced) Device Type*/
const SPECIFIC_TYPE_WHOLE_HOME_METER_SIMPLE = 0x03 /*Whole Home Meter (Simple) Device Type*/

/* Device class - Entry Control */
const GENERIC_TYPE_ENTRY_CONTROL = 0x40                     /*Entry Control*/
const SPECIFIC_TYPE_ENTRY_CONTROL_NOT_USED = 0x00           /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_DOOR_LOCK = 0x01                        /*Door Lock*/
const SPECIFIC_TYPE_ADVANCED_DOOR_LOCK = 0x02               /*Advanced Door Lock*/
const SPECIFIC_TYPE_SECURE_KEYPAD_DOOR_LOCK = 0x03          /*Door Lock (keypad–lever) Device Type*/
const SPECIFIC_TYPE_SECURE_KEYPAD_DOOR_LOCK_DEADBOLT = 0x04 /*Door Lock (keypad– deadbolt) Device Type*/
const SPECIFIC_TYPE_SECURE_DOOR = 0x05                      /*Barrier Operator Specific Device Class*/
const SPECIFIC_TYPE_SECURE_GATE = 0x06                      /*Barrier Operator Specific Device Class*/
const SPECIFIC_TYPE_SECURE_BARRIER_ADDON = 0x07             /*Barrier Operator Specific Device Class*/
const SPECIFIC_TYPE_SECURE_BARRIER_OPEN_ONLY = 0x08         /*Barrier Operator Specific Device Class*/
const SPECIFIC_TYPE_SECURE_BARRIER_CLOSE_ONLY = 0x09        /*Barrier Operator Specific Device Class*/
const SPECIFIC_TYPE_SECURE_LOCKBOX = 0x0A                   /*SDS12724*/
const SPECIFIC_TYPE_SECURE_KEYPAD = 0x0B                    /* Device class Generic Controller */

/* Device class - Semi Interoperable */
const GENERIC_TYPE_SEMI_INTEROPERABLE = 0x50           /*Semi Interoperable*/
const SPECIFIC_TYPE_SEMI_INTEROPERABLE_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_ENERGY_PRODUCTION = 0x01           /*Energy Production*/

/* Device class - Sensor Alarm */
const GENERIC_TYPE_SENSOR_ALARM = 0xA1
const SPECIFIC_TYPE_SENSOR_ALARM_NOT_USED = 0x00 /*Specific Device Class Not Used*/
const SPECIFIC_TYPE_ADV_ZENSOR_NET_ALARM_SENSOR = 0x05
const SPECIFIC_TYPE_ADV_ZENSOR_NET_SMOKE_SENSOR = 0x0A
const SPECIFIC_TYPE_BASIC_ROUTING_ALARM_SENSOR = 0x01
const SPECIFIC_TYPE_BASIC_ROUTING_SMOKE_SENSOR = 0x06
const SPECIFIC_TYPE_BASIC_ZENSOR_NET_ALARM_SENSOR = 0x03
const SPECIFIC_TYPE_BASIC_ZENSOR_NET_SMOKE_SENSOR = 0x08
const SPECIFIC_TYPE_ROUTING_ALARM_SENSOR = 0x02
const SPECIFIC_TYPE_ROUTING_SMOKE_SENSOR = 0x07
const SPECIFIC_TYPE_ZENSOR_NET_ALARM_SENSOR = 0x04
const SPECIFIC_TYPE_ZENSOR_NET_SMOKE_SENSOR = 0x09
const SPECIFIC_TYPE_ALARM_SENSOR = 0x0B /*Sensor (Alarm) Device Type*/

/* Device class - Non Interoperable */
const GENERIC_TYPE_NON_INTEROPERABLE = 0xFF           /*Non interoperable*/
const SPECIFIC_TYPE_NON_INTEROPERABLE_NOT_USED = 0x00 /*Specific Device Class Not Used*/
