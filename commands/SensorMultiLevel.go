package commands

type ZWaveSensorType byte

const (
	Undefined = iota
	Temperature
	General
	Luminance
	Power
	RelativeHumidity
	Velocity
	Direction
	AtmosphericPressure
	BarometricPressure
	SolarRadiation
	DewPoint
	RainRate
	TideLevel
	Weight
	Voltage
	Current
	CO2Level
	AirFlow
	TankCapacity
	Distance
	AnglePosition
	Rotation
	WaterTemperature
	SoilTemperature
	SeismicIntensity
	SeismicMagnitude
	Ultraviolet
	ElectricalResistivity
	ElectricalConductivity
	Loudness
	Moisture
)

type Sensor struct {
	Type      byte
	Scale     byte
	Precision byte
}

type CmdSensorMultiLevel struct {
	Values []Sensor

	data []byte
}

func NewCmdSensorMultiLevel(data []byte) *CmdSensorMultiLevel {
	sml := &CmdSensorMultiLevel{data: data}

	return sml
}
