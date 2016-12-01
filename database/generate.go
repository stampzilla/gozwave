package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// http://products.z-wavealliance.org/products/1729
// http://www.pepper1.net/zwavedb/

type manufacturerId struct {
	Value string `xml:"value,attr"`
}

type productType struct {
	Value string `xml:"value,attr"`
}
type productId struct {
	Value string `xml:"value,attr"`
}
type libType struct {
	Value string `xml:"value,attr"`
}
type protoVersion struct {
	Value string `xml:"value,attr"`
}
type protoSubVersion struct {
	Value string `xml:"value,attr"`
}
type appVersion struct {
	Value string `xml:"value,attr"`
}
type appSubVersion struct {
	Value string `xml:"value,attr"`
}
type basicClass struct {
	Value string `xml:"value,attr"`
}
type genericClass struct {
	Value string `xml:"value,attr"`
}
type specificClass struct {
	Value string `xml:"value,attr"`
}
type optional struct {
	Value string `xml:"value,attr"`
}
type listening struct {
	Value string `xml:"value,attr"`
}
type routing struct {
	Value string `xml:"value,attr"`
}

type deviceData struct {
	ManufacturerId  manufacturerId  `xml:"manufacturerId"`
	ProductType     productType     `xml:"productType"`
	ProductId       productId       `xml:"productId"`
	LibType         libType         `xml:"libType"`
	ProtoVersion    protoVersion    `xml:"protoVersion"`
	ProtoSubVersion protoSubVersion `xml:"protoSubVersion"`
	AppVersion      appVersion      `xml:"appVersion"`
	AppSubVersion   appSubVersion   `xml:"appSubVersion"`
	BasicClass      basicClass      `xml:"basicClass"`
	GenericClass    genericClass    `xml:"genericClass"`
	SpecificClass   specificClass   `xml:"specificClass"`
	Optional        optional        `xml:"optional"`
	Listening       listening       `xml:"listening"`
	Routing         routing         `xml:"routing"`
	BeamSensor      string          `xml:"beamSensor"`
	RfFrequency     string          `xml:"rfFrequency"`
}

type commandClass struct {
	Id         string `xml:"id,attr"`
	Controlled string `xml:"controlled,attr"`
	InNIF      string `xml:"inNIF,attr"`
	Secure     string `xml:"secure,attr"`
	NonSecure  string `xml:"nonSecure,attr"`
	Version    string `xml:"version,attr"`
}

type zWaveDevice struct {
	XMLName           xml.Name       `xml:"ZWaveDevice"`
	DescriptorVersion string         `xml:"descriptorVersion"`
	DeviceData        deviceData     `xml:"deviceData"`
	CommandClasses    []commandClass `xml:"commandClasses>commandClass"`
	//ManufacturerId    string     `xml:"deviceData>manufacturerId,attr"`
}

var templ = `package {{.Package}}
type commandClass struct {
	Id         string
	Controlled string
	InNIF      string
	Secure     bool
	NonSecure  bool
	Version    string 
}
type Device struct{
	CommandClasses []commandClass
	ManufacturerId string
	ProductType string
	ProductId string
}
func New(manufacturerId, productType, productId string) *Device{
	dev := manufacturerId+productType+productId
	switch dev {
{{- range $value := .Devices }}
	case "{{ $value.DeviceData.ManufacturerId.Value }}{{ $value.DeviceData.ProductType.Value }}{{ $value.DeviceData.ProductId.Value }}":
		return New{{ $value.DeviceData.ManufacturerId.Value }}{{ $value.DeviceData.ProductType.Value }}{{ $value.DeviceData.ProductId.Value }}()
{{- end}}
	}
}
{{- range $value := .Devices }}
func New{{ $value.DeviceData.ManufacturerId.Value }}{{ $value.DeviceData.ProductType.Value }}{{ $value.DeviceData.ProductId.Value }}() *Device{
	return &Device{
		ManufacturerId: "{{$value.DeviceData.ManufacturerId.Value }}",
		ProductType: "{{$value.DeviceData.ProductType.Value }}",
		ProductId: "{{$value.DeviceData.ProductId.Value }}",
		CommandClasses: []commandClass{
{{- range $cmd := $value.CommandClasses }}
			&commandClass{
				{{- if ne $cmd.Id ""}}
				Id: "{{$cmd.Id}}",
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
	}
}
{{- end}}
`

//TODO wee need to be more specific in our template above
// 135-0115-1000-0001-06-03-16-01-04.xml
// 132-0115-1000-0001-06-03-16-01-05.xml
// has same ManufacturerId-ProductType-ProductId ?

type templates struct {
	Devices []*zWaveDevice
	Package string
}

var outputfile string
var packageName string

func init() {
	flag.StringVar(&outputfile, "file", "", "Output file")
	flag.StringVar(&packageName, "package", "devices", "Package name of generated code")
	flag.Parse()
}

func main() {
	devices := &templates{
		Package: packageName,
	}
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".xml" {
			continue
		}
		var q *zWaveDevice
		xmlFile, err := os.Open(file.Name())
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		decoder := xml.NewDecoder(xmlFile)

		err = decoder.Decode(&q)
		xmlFile.Close()
		if err != nil {
			log.Println(file.Name())
			log.Println(err)
			return
		}

		//spew.Dump(q)
		devices.Devices = append(devices.Devices, q)
	}

	t := template.New("t")
	t, err = t.Parse(templ)
	if err != nil {
		panic(err)
	}

	if outputfile == "" {
		err = t.Execute(os.Stdout, devices)
	} else {
		file, err := os.Create(outputfile)
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
