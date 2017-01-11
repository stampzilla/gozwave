package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/xml"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// http://products.z-wavealliance.org/products/1729
// http://www.pepper1.net/zwavedb/

type deviceDataValue struct {
	Value string `xml:"value,attr"`
}

type deviceData struct {
	ManufacturerID  deviceDataValue `xml:"manufacturerId"`
	ProductType     deviceDataValue `xml:"productType"`
	ProductID       deviceDataValue `xml:"productId"`
	LibType         deviceDataValue `xml:"libType"`
	ProtoVersion    deviceDataValue `xml:"protoVersion"`
	ProtoSubVersion deviceDataValue `xml:"protoSubVersion"`
	AppVersion      deviceDataValue `xml:"appVersion"`
	AppSubVersion   deviceDataValue `xml:"appSubVersion"`
	BasicClass      deviceDataValue `xml:"basicClass"`
	GenericClass    deviceDataValue `xml:"genericClass"`
	SpecificClass   deviceDataValue `xml:"specificClass"`
	Optional        deviceDataValue `xml:"optional"`
	Listening       deviceDataValue `xml:"listening"`
	Routing         deviceDataValue `xml:"routing"`
	BeamSensor      string          `xml:"beamSensor"`
	RfFrequency     string          `xml:"rfFrequency"`
}

type commandClass struct {
	ID         string `xml:"id,attr"`
	Controlled string `xml:"controlled,attr"`
	InNIF      string `xml:"inNIF,attr"`
	Secure     string `xml:"secure,attr"`
	NonSecure  string `xml:"nonSecure,attr"`
	Version    string `xml:"version,attr"`
}

func (c commandClass) IDasHex() string {
	return "0x" + c.ID
}

type deviceDescription struct {
	BrandName   string `xml:"brandName"`
	ProductName string `xml:"productName"`
	Description string `xml:"description"`
}

func (dd deviceDescription) EscapedDescription() string {
	return removeNewLines(dd.Description)
}

func (dd deviceDescription) EscapedProductName() string {
	return removeNewLines(dd.ProductName)
}
func (dd deviceDescription) EscapedBrandName() string {
	return removeNewLines(dd.BrandName)
}

type lang struct {
	Language string `xml:"lang,attr"`
	Body     string `xml:",chardata"`
}

type value struct {
	From string `xml:"from,attr"`
	To   string `xml:"to,attr"`
	Desc []lang `xml:"description>lang"`
	Unit string `xml:"unit,attr"`

	size int
}

func (p value) FromDec() string {
	return strToUint(p.From, p.size)
}
func (p value) ToDec() string {
	return strToUint(p.To, p.size)
}

func strToUint(str string, size int) string {
	b, _ := hex.DecodeString(str)
	buf := bytes.NewReader(b)

	var res int

	switch size {
	case 1:
		var i uint8
		binary.Read(buf, binary.BigEndian, &i)
		res = int(i)
	case 2:
		var i uint16
		binary.Read(buf, binary.BigEndian, &i)
		res = int(i)
	case 4:
		var i uint32
		binary.Read(buf, binary.BigEndian, &i)
		res = int(i)
	}

	ret := strconv.Itoa(res)
	return ret
}

func (p value) EscapedDesc() string {
	for _, v := range p.Desc {
		if v.Language == "en" && v.Body != "" {
			return removeNewLines(v.Body)
		}
	}

	return ""
}

type parameter struct {
	ID      string `xml:"number,attr"`
	Name    []lang `xml:"name>lang"`
	Desc    []lang `xml:"description>lang"`
	Type    string `xml:"type,attr"`
	Size    string `xml:"size,attr"`
	Default string `xml:"default,attr"`

	Values []value `xml:"value"`
}

func (p parameter) NameCombined() string {
	for _, v := range p.Name {
		if v.Language == "en" && v.Body != "" {
			return removeNewLines(v.Body)
		}
	}
	for _, v := range p.Desc {
		if v.Language == "en" && v.Body != "" {
			return removeNewLines(v.Body)
		}
	}

	return "Parameter " + p.ID
}

type zWaveDevice struct {
	Filename          string
	Name              string
	XMLName           xml.Name          `xml:"ZWaveDevice"`
	DescriptorVersion string            `xml:"descriptorVersion"`
	DeviceData        deviceData        `xml:"deviceData"`
	DeviceDescription deviceDescription `xml:"deviceDescription"`
	CommandClasses    []commandClass    `xml:"commandClasses>commandClass"`
	Parameters        []parameter       `xml:"configParams>configParam"`
}

var templ = `package {{.Package}}

import (
	"github.com/stampzilla/gozwave/commands"
)

type CommandClass struct {
	ID         commands.ZWaveCommand
	Controlled string
	InNIF      string
	Secure     bool
	NonSecure  bool
	Version    string 
}
type Value struct {
  From int
  To int
  Desc string
  Unit string
}
type parameter struct {
	ID int
	Name string
	Type string
	Description string
	Size int
	Default string

	Values []Value
}
type Device struct{
	Brand string
	Product string
	Description string 

	CommandClasses []*CommandClass
	Parameters []*parameter

	ManufacturerID string
	ProductType string
	ProductID string
}
func New(manufacturerID, productType, productID string) *Device{
	dev := manufacturerID+productType+productID
	switch dev {
{{- range $value := .Devices }}
	case "{{ $value.DeviceData.ManufacturerID.Value }}{{ $value.DeviceData.ProductType.Value }}{{ $value.DeviceData.ProductID.Value }}":
		return New{{ $value.DeviceData.ManufacturerID.Value }}{{ $value.DeviceData.ProductType.Value }}{{ $value.DeviceData.ProductID.Value }}() // {{ $value.Filename }} | {{ $value.Name }}
{{- end}}
	}

	return nil
}
{{- range $value := .Devices }}
func New{{ $value.DeviceData.ManufacturerID.Value }}{{ $value.DeviceData.ProductType.Value }}{{ $value.DeviceData.ProductID.Value }}() *Device{
	return &Device{
		Brand: "{{ $value.DeviceDescription.EscapedBrandName }}",
		Product: "{{ $value.DeviceDescription.EscapedProductName }}",
		Description: "{{ $value.DeviceDescription.EscapedDescription }}",

		ManufacturerID: "{{$value.DeviceData.ManufacturerID.Value }}",
		ProductType: "{{$value.DeviceData.ProductType.Value }}",
		ProductID: "{{$value.DeviceData.ProductID.Value }}",
		CommandClasses: []*CommandClass{
{{- range $cmd := $value.CommandClasses }}
			&CommandClass{
				{{- if ne $cmd.ID ""}}
				ID: {{$cmd.IDasHex}},
				{{- end }}
				{{- if ne $cmd.Controlled "" }}
				Controlled: "{{$cmd.Controlled}}",
				{{- end }}
				{{- if ne $cmd.InNIF ""}}
				InNIF: "{{$cmd.InNIF}}",
				{{- end }}
				{{- if ne $cmd.Secure ""}}
				Secure: {{$cmd.Secure}},
				{{- end }}
				{{- if ne $cmd.NonSecure ""}}
				NonSecure: {{$cmd.NonSecure}},
				{{- end }}
				{{- if ne $cmd.Version ""}}
				Version: "{{$cmd.Version}}",
				{{- end }}
			},
{{- end}}
		},
		Parameters: []*parameter{
{{- range $cmd := $value.Parameters }}
			&parameter{
				{{- if ne $cmd.ID ""}}
				ID: {{$cmd.ID}},
				{{- end }}
				Name: "{{$cmd.NameCombined}}",
				Type: "{{$cmd.Type}}",
				Size: {{$cmd.Size}},
				Values: []Value{
{{- range $value := $cmd.Values }}
					Value{
						From: {{$value.FromDec }},
						To: {{$value.ToDec}},
						Desc: "{{$value.EscapedDesc}}",
						Unit: "{{$value.Unit}}",
					},
{{- end }}
				},
			},
{{- end}}
		},
	}
}
{{- end}}
`

type templates struct {
	Devices map[string]*zWaveDevice
	Package string
}

var outputfile string
var databaseDir string
var packageName string

func main() {
	flag.StringVar(&outputfile, "file", "", "Output file")
	flag.StringVar(&databaseDir, "databasedir", "./database", "Directory with xml files")
	flag.StringVar(&packageName, "package", "devices", "Package name of generated code")

	flag.Parse()

	devices := &templates{
		Package: packageName,
	}
	files, err := ioutil.ReadDir(databaseDir)
	if err != nil {
		log.Fatal(err)
	}

	devices.Devices = make(map[string]*zWaveDevice)

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".xml" {
			continue
		}
		var dev *zWaveDevice
		xmlFile, err := os.Open(filepath.Join(databaseDir, file.Name()))
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		decoder := xml.NewDecoder(xmlFile)

		err = decoder.Decode(&dev)
		xmlFile.Close()
		if err != nil {
			log.Println(file.Name())
			log.Println(err)
			return
		}

		dev.Filename = file.Name()
		dev.Name =
			dev.DeviceData.ManufacturerID.Value +
				dev.DeviceData.ProductType.Value +
				dev.DeviceData.ProductID.Value

		// Copy the parameter size to each value (used to decode the values)
		for k, p := range dev.Parameters {
			s, _ := strconv.Atoi(p.Size)
			for k2, _ := range p.Values {
				dev.Parameters[k].Values[k2].size = s
			}
		}

		//spew.Dump(dev)
		//TODO wee need to be more specific in our template above
		// 135-0115-1000-0001-06-03-16-01-04.xml
		// 132-0115-1000-0001-06-03-16-01-05.xml
		// has same ManufacturerID-ProductType-ProductID ?
		devices.Devices[dev.Name] = dev

		//spew.Dump(dev.Parameters)
	}

	t := template.New("t")
	t, err = t.Parse(templ)
	if err != nil {
		panic(err)
	}

	if outputfile == "" {
		err = t.Execute(os.Stdout, devices)
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
		err = t.Execute(file, devices)
		file.Close()
	}
	if err != nil {
		panic(err)
	}

}

func removeNewLines(s string) string {
	s = strings.Replace(s, "\\", `\\`, -1)
	s = strings.Replace(s, "\n", `\n`, -1)
	s = strings.Replace(s, "\r", `\n`, -1)
	return s
}
