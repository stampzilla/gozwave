package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/stampzilla/gozwave/commands"
)

type templates struct {
	CommandClasses map[string][]string
	Package        string
}

var outputfile string
var mandatoryFile string
var packageName string

func main() {
	flag.StringVar(&outputfile, "file", "", "Output file")
	flag.StringVar(&mandatoryFile, "mandatoryfile", "./database", "the text file with specification")
	flag.StringVar(&packageName, "package", "devices", "Package name of generated code")

	flag.Parse()

	data := &templates{
		CommandClasses: make(map[string][]string),
		Package:        packageName,
	}
	//data.CommandClasses[database.Definition{Generic: 0x01, Specific: 0x02}] = []string{
	//"0x10",
	//"0x11",
	//}

	file, err := os.Open(mandatoryFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	defs := [][]string{}
	classes := []string{}
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			for _, def := range defs {
				data.CommandClasses[def[0]+"|"+def[1]] = classes
			}
			defs = [][]string{}
			classes = []string{}
			continue
		}

		if strings.HasPrefix(text, "#") {
			tmp := strings.Split(text, "-")
			if strings.Contains(text, "SPECIFIC_TYPE_NOT_USED") {
				//TODO skip this for now. Do something smarter if AV control point is needed
				continue
			}
			defs = append(defs, []string{
				strings.TrimSpace(tmp[len(tmp)-2]),
				strings.TrimSpace(tmp[len(tmp)-1]),
			})
		}

		if strings.HasPrefix(text, "-") {
			class := strings.TrimLeft(text, "-")
			class = strings.TrimSpace(class)
			x := StringToByte(class)
			if x == 0 {
				continue
			}
			classes = append(classes, fmt.Sprintf("0x%x", x))
			//fmt.Println(text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	t := template.New("t")
	t.Funcs(template.FuncMap{
		"getGeneric": func(s string) string {
			tmp := strings.Split(s, "|")
			return tmp[0]
		},
		"getSpecific": func(s string) string {
			tmp := strings.Split(s, "|")
			return tmp[1]
		},
	})
	t, err = t.Parse(templ)
	if err != nil {
		panic(err)
	}

	if outputfile == "" {
		err = t.Execute(os.Stdout, data)
	} else {
		var file *os.File
		file, err = os.Create(outputfile)
		if err != nil {
			log.Println(err)
			return
		}
		if err != nil {
			log.Println("Unable to open file", err)
			return
		}
		err = t.Execute(file, data)
		file.Close()
	}
	if err != nil {
		panic(err)
	}

}

func StringToByte(s string) byte {

	//- Application Status
	if strings.Contains(s, "Application Status") {
		return commands.ApplicationStatus
	}

	//- Association Group Information
	if strings.Contains(s, "Association Group Information") {
		return commands.AssociationGRPInfo
	}

	//- Association (version 2 or newer)
	if strings.Contains(s, "Association") {
		return commands.Association
	}

	//- Basic
	if s == "Basic" {
		return commands.Basic
	}

	//- Binary Switch
	if strings.Contains(s, "Binary Switch") {
		return commands.SwitchBinary
	}
	//- Central Scene
	if strings.Contains(s, "Central Scene") {
		return commands.CentralScenev2
	}
	//- CRC16
	if strings.Contains(s, "CRC16") {
		return commands.CRC16Encap
	}
	//- Device Reset Locally (*)
	//- Device Reset Locally (if the device can be reset, refer to [1])
	if strings.Contains(s, "Device Reset Locally") {
		return commands.DeviceResetLocally
	}
	//- Manufacturer Specific
	if s == "Manufacturer Specific" {
		return commands.ManufacturerSpecific
	}
	//- Meter (version 2 or newer)
	if strings.Contains(s, "Meter") {
		return commands.Meter
	}
	//- Multichannel Association (version 3 or newer)
	if strings.Contains(s, "Multichannel Association") {
		return commands.MultiInstanceAssociation
	}

	//- Multi Channel (version 4 or newer)
	if strings.Contains(s, "Multi Channel") {
		return commands.MultiInstance
	}

	//- Multilevel Sensor
	if strings.Contains(s, "Multilevel Sensor") {
		return commands.SensorMultiLevel
	}

	//- Multilevel Switch
	//- Multilevel Switch (v3 or newer)
	if strings.Contains(s, "Multilevel Switch") {
		return commands.SwitchMultilevel
	}
	//- Notification
	if s == "Notification" {
		return commands.Alarm
	}
	//- Power Level
	if s == "Power Level" {
		return commands.PowerLevel
	}
	//- Screen Meta Data

	//- Security
	//- Security - if CSC
	if strings.Contains(s, "Security") {
		return commands.Security
	}

	//- Simple AV
	//- Simple AV Control
	//- Thermostat Mode
	//- Thermostat Setback
	//- Thermostat Set Point
	//- Version (v2 or newer)
	if strings.Contains(s, "Version") {
		return commands.Version
	}

	//- Z-Wave Plus Info
	if strings.Contains(s, "Z-Wave Plus Info") {
		return commands.ZwavePlusInfov2
	}

	log.Println("CommandClass", s, " not found")
	//os.Exit(1)
	return 0
}

var templ = `package {{.Package}}
import (
	"github.com/stampzilla/gozwave/protocol"
)

var definitions = map[Definition][]*CommandClass{
{{- range $key, $value := .CommandClasses }}
	Definition{
		Generic: protocol.{{ getGeneric $key }},
		Specific: protocol.{{ getSpecific $key }},
	}: []*CommandClass{
{{- range $cmdClass := $value }}
		&CommandClass{
			ID: {{ $cmdClass }},
		},
{{- end }}
	},

{{- end}}
}

`
