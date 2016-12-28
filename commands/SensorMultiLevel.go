package commands

import (
	"bytes"
	"encoding/binary"
	"math"
	"strconv"

	"github.com/Sirupsen/logrus"
)

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

type SensorMultiLevelReport struct {
	*report
	Type       byte
	Size       byte
	Scale      byte
	Precision  byte
	Value      float64
	TypeString string
	Unit       string

	data []byte
}

func NewSensorMultiLevelReport(data []byte) *SensorMultiLevelReport {
	sml := &SensorMultiLevelReport{data: data}

	sml.Type = data[0]
	sml.Size = (data[1] & 0x07)              // Size
	sml.Scale = (data[1] & 0x18) >> 0x03     // Scale
	sml.Precision = (data[1] & 0xE0) >> 0x05 // Precision

	buf := bytes.NewReader(data[2:])
	var err error

	switch sml.Size {
	case 1:
		val := int8(0)
		err = binary.Read(buf, binary.BigEndian, &val)
		sml.Value = float64(val)
	case 2:
		val := int16(0)
		err = binary.Read(buf, binary.BigEndian, &val)
		sml.Value = float64(val)
	case 4:
		val := int32(0)
		err = binary.Read(buf, binary.BigEndian, &val)
		sml.Value = float64(val)
	case 8:
		val := int64(0)
		err = binary.Read(buf, binary.BigEndian, &val)
		sml.Value = float64(val)
	}

	if err != nil {
		logrus.Warn("Failed to decode SensorMultiLevelReport: %s", err)
		return nil
	}

	if sml.Precision > 0 {
		sml.Value /= math.Pow(10, float64(sml.Precision))
	}

	switch sml.Type {
	case Temperature:
		sml.TypeString = "Temperature"
		switch sml.Scale {
		case 0:
			sml.Unit = "C"
		case 1:
			sml.Unit = "F"
		}
	case General:
		sml.TypeString = "General"
		switch sml.Scale {
		case 0:
			sml.Unit = "%"
		}
	case Luminance:
		sml.TypeString = "Luminance"
		switch sml.Scale {
		case 0:
			sml.Unit = "%"
		case 1:
			sml.Unit = "lux"
		}
	case Power:
		sml.TypeString = "Power"
		switch sml.Scale {
		case 0:
			sml.Unit = "W"
		case 1:
			sml.Unit = "BTU/h"
		}
	case RelativeHumidity:
		sml.TypeString = "RelativeHumidity"
		switch sml.Scale {
		case 0:
			sml.Unit = "%"
		}
	case Velocity:
		sml.TypeString = "Velocity"
		switch sml.Scale {
		case 0:
			sml.Unit = "m/s"
		case 1:
			sml.Unit = "mph"
		}
	case Direction:
		sml.TypeString = "Direction"
	case AtmosphericPressure:
		sml.TypeString = "AtmosphericPressure"
		switch sml.Scale {
		case 0:
			sml.Unit = "kPa"
		case 1:
			sml.Unit = "inHg"
		}
	case BarometricPressure:
		sml.TypeString = "BarometricPressure"
		switch sml.Scale {
		case 0:
			sml.Unit = "kPa"
		case 1:
			sml.Unit = "inHg"
		}
	case SolarRadiation:
		sml.TypeString = "SolarRadiation"
		sml.Unit = "W/m2"
	case DewPoint:
		sml.TypeString = "DewPoint"
		switch sml.Scale {
		case 0:
			sml.Unit = "C"
		case 1:
			sml.Unit = "F"
		}
	case RainRate:
		sml.TypeString = "RainRate"
		switch sml.Scale {
		case 0:
			sml.Unit = "mm/h"
		case 1:
			sml.Unit = "in/h"
		}
	case TideLevel:
		sml.TypeString = "TideLevel"
		switch sml.Scale {
		case 0:
			sml.Unit = "m"
		case 1:
			sml.Unit = "ft"
		}
	case Weight:
		sml.TypeString = "Weight"
		switch sml.Scale {
		case 0:
			sml.Unit = "kg"
		case 1:
			sml.Unit = "lb"
		}
	case Voltage:
		sml.TypeString = "Voltage"
		switch sml.Scale {
		case 0:
			sml.Unit = "V"
		case 1:
			sml.Unit = "mV"
		}
	case Current:
		sml.TypeString = "Current"
		switch sml.Scale {
		case 0:
			sml.Unit = "A"
		case 1:
			sml.Unit = "mA"
		}
	case CO2Level:
		sml.TypeString = "CO2"
		sml.Unit = "ppm"
	case AirFlow:
		sml.TypeString = "AirFlow"
		switch sml.Scale {
		case 0:
			sml.Unit = "m3/h"
		case 1:
			sml.Unit = "cfm"
		}
	case TankCapacity:
		sml.TypeString = "TankCapacity"
		switch sml.Scale {
		case 0:
			sml.Unit = "l"
		case 1:
			sml.Unit = "cbm"
		case 2:
			sml.Unit = "gal"
		}
	case Distance:
		sml.TypeString = "Distance"
		switch sml.Scale {
		case 0:
			sml.Unit = "m"
		case 1:
			sml.Unit = "cm"
		case 2:
			sml.Unit = "ft"
		}
	case AnglePosition:
		sml.TypeString = "AnglePosition"
		switch sml.Scale {
		case 0:
			sml.Unit = "%"
		case 1:
			sml.Unit = "deg N"
		case 2:
			sml.Unit = "deg S"
		}
	case Rotation:
		sml.TypeString = "Rotation"
		switch sml.Scale {
		case 0:
			sml.Unit = "rpm"
		case 1:
			sml.Unit = "hz"
		}
	case WaterTemperature:
		sml.TypeString = "WaterTemperature"
		switch sml.Scale {
		case 0:
			sml.Unit = "C"
		case 1:
			sml.Unit = "F"
		}
	case SoilTemperature:
		sml.TypeString = "SoilTemperature"
		switch sml.Scale {
		case 0:
			sml.Unit = "C"
		case 1:
			sml.Unit = "F"
		}
	case SeismicIntensity:
		sml.TypeString = "SeismicIntensity"
		switch sml.Scale {
		case 0:
			sml.Unit = "mercalli"
		case 1:
			sml.Unit = "EU macroseismic"
		case 2:
			sml.Unit = "liedu"
		case 3:
			sml.Unit = "shindo"
		}
	case SeismicMagnitude:
		sml.TypeString = "SeismicMagnitude"
		switch sml.Scale {
		case 0:
			sml.Unit = "local"
		case 1:
			sml.Unit = "moment"
		case 2:
			sml.Unit = "surface wave"
		case 3:
			sml.Unit = "body wave"
		}
	case Ultraviolet:
		sml.TypeString = "Ultraviolet"
		sml.Unit = ""
	case ElectricalResistivity:
		sml.TypeString = "ElectricalResistivity"
		sml.Unit = "ohm"
	case ElectricalConductivity:
		sml.TypeString = "ElectricalConductivity"
		sml.Unit = "siemens/m"
	case Loudness:
		sml.TypeString = "Loudness"
		switch sml.Scale {
		case 0:
			sml.Unit = "db"
		case 1:
			sml.Unit = "dBA"
		}
	case Moisture:
		sml.TypeString = "Moisture"
		switch sml.Scale {
		case 0:
			sml.Unit = "%"
		case 1:
			sml.Unit = "content"
		case 2:
			sml.Unit = "k ohm"
		case 3:
			sml.Unit = "water activity"
		}
	default:
		sml.TypeString = "Unknown (" + strconv.Itoa(int(sml.Type)) + ")"
	}

	return sml
}
