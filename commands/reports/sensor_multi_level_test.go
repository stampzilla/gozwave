package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensorMultiLevelNoData(t *testing.T) {
	w, err := NewSensorMultiLevel([]byte{})

	assert.Error(t, err)
	assert.Nil(t, w)
}

func TestSensorMultiLevelShort(t *testing.T) {
	w, err := NewSensorMultiLevel([]byte{0x01, 0x01})

	assert.Error(t, err)
	assert.Nil(t, w)
}

func TestSensorMultiLevel1(t *testing.T) {
	w, err := NewSensorMultiLevel([]byte{0x01, 0x21, 0x01})

	assert.NoError(t, err)
	assert.IsType(t, w, &SensorMultiLevel{})
}
func TestSensorMultiLevel2(t *testing.T) {
	w, err := NewSensorMultiLevel([]byte{0x01, 0x22, 0x00, 0x01})

	assert.NoError(t, err)
	assert.IsType(t, w, &SensorMultiLevel{})
}
func TestSensorMultiLevel4(t *testing.T) {
	w, err := NewSensorMultiLevel([]byte{0x01, 0x24, 0x00, 0x00, 0x00, 0x01})

	assert.NoError(t, err)
	assert.IsType(t, w, &SensorMultiLevel{})
}

func TestSensorMultiLevelTypes(t *testing.T) {
	d := map[byte]map[byte]string{
		0x01: map[byte]string{
			0x00: "Temperature C",
			0x01: "Temperature F",
		},
		0x02: map[byte]string{
			0x00: "General %",
		},
		0x03: map[byte]string{
			0x00: "Luminance %",
			0x01: "Luminance lux",
		},
		0x04: map[byte]string{
			0x00: "Power W",
			0x01: "Power BTU/h",
		},
		0x05: map[byte]string{
			0x00: "RelativeHumidity %",
		},
		0x06: map[byte]string{
			0x00: "Velocity m/s",
			0x01: "Velocity mph",
		},
		0x07: map[byte]string{
			0x00: "Direction ",
		},
		0x08: map[byte]string{
			0x00: "AtmosphericPressure kPa",
			0x01: "AtmosphericPressure inHg",
		},
		0x09: map[byte]string{
			0x00: "BarometricPressure kPa",
			0x01: "BarometricPressure inHg",
		},
		0x0A: map[byte]string{
			0x00: "SolarRadiation W/m2",
		},
		0x0B: map[byte]string{
			0x00: "DewPoint C",
			0x01: "DewPoint F",
		},
		0x0C: map[byte]string{
			0x00: "RainRate mm/h",
			0x01: "RainRate in/h",
		},
		0x0D: map[byte]string{
			0x00: "TideLevel m",
			0x01: "TideLevel ft",
		},
		0x0E: map[byte]string{
			0x00: "Weight kg",
			0x01: "Weight lb",
		},
		0x0F: map[byte]string{
			0x00: "Voltage V",
			0x01: "Voltage mV",
		},
		0x10: map[byte]string{
			0x00: "Current A",
			0x01: "Current mA",
		},
		0x11: map[byte]string{
			0x00: "CO2 ppm",
		},
		0x12: map[byte]string{
			0x00: "AirFlow m3/h",
			0x01: "AirFlow cfm",
		},
		0x13: map[byte]string{
			0x00: "TankCapacity l",
			0x01: "TankCapacity cbm",
			0x02: "TankCapacity gal",
		},
		0x14: map[byte]string{
			0x00: "Distance m",
			0x01: "Distance cm",
			0x02: "Distance ft",
		},
		0x15: map[byte]string{
			0x00: "AnglePosition %",
			0x01: "AnglePosition deg N",
			0x02: "AnglePosition deg S",
		},
		0x16: map[byte]string{
			0x00: "Rotation rpm",
			0x01: "Rotation hz",
		},
		0x17: map[byte]string{
			0x00: "WaterTemperature C",
			0x01: "WaterTemperature F",
		},
		0x18: map[byte]string{
			0x00: "SoilTemperature C",
			0x01: "SoilTemperature F",
		},
		0x19: map[byte]string{
			0x00: "SeismicIntensity mercalli",
			0x01: "SeismicIntensity EU macroseismic",
			0x02: "SeismicIntensity liedu",
			0x03: "SeismicIntensity shindo",
		},
		0x1A: map[byte]string{
			0x00: "SeismicMagnitude local",
			0x01: "SeismicMagnitude moment",
			0x02: "SeismicMagnitude surface wave",
			0x03: "SeismicMagnitude body wave",
		},
		0x1B: map[byte]string{
			0x00: "Ultraviolet ",
		},
		0x1C: map[byte]string{
			0x00: "ElectricalResistivity ohm",
		},
		0x1D: map[byte]string{
			0x00: "ElectricalConductivity siemens/m",
		},
		0x1E: map[byte]string{
			0x00: "Loudness db",
			0x01: "Loudness dBA",
		},
		0x1F: map[byte]string{
			0x00: "Moisture %",
			0x01: "Moisture content",
			0x02: "Moisture k ohm",
			0x03: "Moisture water activity",
		},
		0x20: map[byte]string{
			0x00: "Unknown (32) ",
		},
	}

	for ty, scales := range d {
		for scale, expectedTypeString := range scales {
			w, err := NewSensorMultiLevel([]byte{ty, 0x04 + (scale << 0x03), 0x00, 0x00, 0x00, 0x01})

			assert.NoError(t, err)
			assert.IsType(t, &SensorMultiLevel{}, w)
			assert.Equal(t, "1.000000 "+expectedTypeString, w.String())
		}
	}

}
