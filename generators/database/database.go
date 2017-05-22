package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// http://www.cd-jackson.com/index.php/zwave/zwave-device-database

type deviceData struct {
	Label               string         `json:"label"`
	ManufacturerName    string         `json:"manufacturer_name"`
	ManufacturerID      string         `json:"manufacturer_id"`
	DeviceRefs          []deviceRef    `json:"device_ref"`
	VersionMin          float64        `json:"version_min"`
	VersionMax          float64        `json:"version_max"`
	Decription          string         `json:"decription"`
	Overview            string         `json:"overview"`
	Inclusion           string         `json:"inclusion"`
	Exclusion           string         `json:"exclusion"`
	Wakeup              string         `json:"wakeup"`
	ProtocolVersion     float64        `json:"protocolVersion"`
	Listening           bool           `json:"listening"`
	FrequentlyListening bool           `json:"frequently_listening"`
	Routing             bool           `json:"routing"`
	Beaming             bool           `json:"beaming"`
	LibraryType         idName         `json:"library_type"`
	BasicClass          idName         `json:"basic_class"`
	GenericClass        idName         `json:"generic_class"`
	SpecificClass       idName         `json:"specific_class"`
	Endpoints           []endpoint     `json:"endpoints"`
	CommandClasses      []commandClass `json:"-"`
	Parameters          []parameter    `json:"parameters"`
	Associations        []association  `json:"associations"`
}

type deviceRef struct {
	ProductType string
	ProductID   string
}

func (ref *deviceRef) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	ss := strings.Split(s, ":")
	switch len(ss) {
	case 2:
		ref.ProductID = ss[1]
		fallthrough
	case 1:
		ref.ProductType = ss[0]
	}
	return nil
}

type idName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type endpoint struct {
	BasicClass     idName         `json:"basic_class"`
	GenericClass   idName         `json:"generic_class"`
	SpecificClass  idName         `json:"specific_class"`
	CommandClasses []commandClass `json:"commandclasses"`
}

type commandClass struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	Nif       bool      `json:"nif"`
	Basic     bool      `json:"basic"`
	Secure    bool      `json:"secure"`
	NonSecure bool      `json:"nonsecure"`
	Config    string    `json:"config"`
	Channels  []channel `json:"channels"`
}

type channel struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	//	Config []??? `json:"config"`
}

type option struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

type parameter struct {
	ID          int      `json:"id"`
	Label       string   `json:"label"`
	Description string   `json:"description"`
	Overview    string   `json:"overview"`
	Size        int      `json:"size"`
	Bitmask     string   `json:"bitmask"`
	Default     int      `json:"default"`
	ReadOnly    bool     `json:"read_only"`
	WriteOnly   bool     `json:"write_only"`
	ValueMin    int      `json:"value_min"`
	ValueMax    int      `json:"value_max"`
	Options     []option `json:"options"`
}

type association struct {
	ID          int    `json:"id"`
	Label       string `json:"label"`
	MaxNodes    int    `json:"max_nodes"`
	Controller  bool   `json:"controller"`
	Description string `json:"description"`
	Overview    string `json:"overview"`
}

var escaper = strings.NewReplacer(
	"\\", "\\\\", // backslash
	"\"", "\\\"", // quote
	"\r", "\\r",
	"\n", "\\n",
)

var funcMap = template.FuncMap{
	"escape": func(s string) string {
		return escaper.Replace(s)
	},
	"lowercase": func(s string) string {
		return strings.ToLower(s)
	},
}

var templ = `package {{.Package}}

type Value struct {
	From int64
	To int64
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
	dev := manufacturerID+"_"+productType+"_"+productID
	switch dev {
{{- range $value := .Devices }}
{{- range $ref := $value.DeviceRefs }}
	case "{{ $value.ManufacturerID | lowercase }}_{{ $ref.ProductType | lowercase }}_{{ $ref.ProductID | lowercase }}":
		return New{{ $value.ManufacturerID }}_{{ $ref.ProductType }}_{{ $ref.ProductID }}()
{{- end}}
{{- end}}
	}

	return nil
}
{{- range $value := .Devices }}
{{- range $ref := $value.DeviceRefs }}
func New{{ $value.ManufacturerID }}_{{ $ref.ProductType }}_{{ $ref.ProductID }}() *Device{
	return &Device{
		Brand: "{{ $value.ManufacturerName | escape }}",
		Product: "{{ $value.Label | escape }}",
		Description: "{{ $value.Decription | escape }}",

		ManufacturerID: "{{ $value.ManufacturerID }}",
		ProductType: "{{ $ref.ProductType }}",
		ProductID: "{{ $ref.ProductID }}",
		CommandClasses: []*CommandClass{
{{- range $cmd := $value.CommandClasses }}
			&CommandClass{
				{{- if ne $cmd.ID ""}}
				ID: 0x{{$cmd.ID}},
				{{- end }}
				{{- if $cmd.Nif}}
				InNIF: {{$cmd.Nif}},
				{{- end }}
				{{- if $cmd.Secure}}
				Secure: {{$cmd.Secure}},
				{{- end }}
				{{- if $cmd.NonSecure}}
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
				{{- if gt $cmd.ID 0}}
				ID: {{$cmd.ID}},
				{{- end }}
				Name: "{{$cmd.Label | escape}}",
				Size: {{$cmd.Size}},
				Values: []Value{
{{- range $option := $cmd.Options }}
					Value{
						From: {{$option.Value}},
						To: {{$option.Value}},
						Desc: "{{$option.Label | escape}}",
					},
{{- else}}
					Value{
						From: {{$cmd.ValueMin}},
						To: {{$cmd.ValueMax}},
					},
{{- end}}
				},
			},
{{- end}}
		},
	}
}
{{- end}}
{{- end}}
`

type templates struct {
	Package string
	Devices []deviceData
}

var outputfile string
var databaseDir string
var packageName string

func main() {
	flag.StringVar(&outputfile, "file", "./database/devices.go", "Output file")
	flag.StringVar(&databaseDir, "databasedir", "./database/json", "Directory with xml files")
	flag.StringVar(&packageName, "package", "database", "Package name of generated code")

	flag.Parse()

	devices := &templates{
		Package: packageName,
	}
	files, err := ioutil.ReadDir(databaseDir)
	if err != nil {
		log.Fatal(err)
	}

	idens := map[string]bool{}
	devices.Devices = make([]deviceData, 0, 1024)

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}
		jsonFile, err := os.Open(filepath.Join(databaseDir, file.Name()))
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		decoder := json.NewDecoder(jsonFile)
		var dev deviceData
		if err = decoder.Decode(&dev); err != nil {
			log.Println(file.Name())
			log.Println(err)
			return
		}
		switch len(dev.Endpoints) {
		default:
			// FIXME: How do we support more than one endpoint?
			//        Is that "index" on openzwave setValue? "instance"?
			log.Printf("warning: %s supports more than one endpoint\n", file.Name())
			fallthrough
		case 1:
			dev.CommandClasses = dev.Endpoints[0].CommandClasses
		case 0:
		}

		refs := []deviceRef{}
		for _, ref := range dev.DeviceRefs {
			iden := dev.ManufacturerID + "_" + ref.ProductType + "_" + ref.ProductID
			if idens[iden] {
				// FIXME: How do we handle duplicates?
				log.Printf("warning: %s duplicate iden %s\n", iden, file.Name())
				continue
			}
			idens[iden] = true
			refs = append(refs, ref)
		}

		if len(refs) > 0 {
			dev.DeviceRefs = refs
			devices.Devices = append(devices.Devices, dev)
		}
	}

	t := template.New("t").Funcs(funcMap)
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
