package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Root struct {
	Typedefs        []TypeDef         `xml:"Typedef"`
	Structs         []Struct          `xml:"Struct"`
	Fields          []Field           `xml:"Field"`
	FundamentalType []FundamentalType `xml:"FundamentalType"`
	ArrayType       []ArrayType       `xml:"ArrayType"`
	PointerType     []PointerType     `xml:"PointerType"`
}

type TypeDef struct {
	ID      string `xml:"id,attr"`
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"`
	Context string `xml:"context,attr"`
}

type Struct struct {
	ID      string `xml:"id,attr"`
	Name    string `xml:"name,attr"`
	Context string `xml:"context,attr"`
	Members string `xml:"members,attr"`
}

type Field struct {
	ID      string `xml:"id,attr"`
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"`
	Context string `xml:"context,attr"`
	Offset  string `xml:"offset,attr"`
}

type FundamentalType struct {
	ID    string `xml:"id,attr"`
	Name  string `xml:"name,attr"`
	Size  string `xml:"size,attr"`
	Align string `xml:"align,attr"`
}

type ArrayType struct {
	ID   string `xml:"id,attr"`
	Min  string `xml:"min,attr"`
	Max  string `xml:"max,attr"`
	Type string `xml:"type,attr"`
}

type PointerType struct {
	ID   string `xml:"id,attr"`
	Type string `xml:"type,attr"`
}

type tree struct {
	list map[string]interface{}
}

func main() {
	f, err := os.Open("ZW_classcmd.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	b, _ := ioutil.ReadAll(f)

	var d Root
	xml.Unmarshal(b, &d)

	t := makeTree(d)

	fmt.Printf("%#v\n", t)
}

func makeTree(r Root) tree {
	t := tree{
		list: make(map[string]interface{}),
	}

	for _, v := range r.Typedefs {
		t.list[v.ID] = v
	}
	for _, v := range r.Structs {
		t.list[v.ID] = v
	}
	for _, v := range r.Fields {
		t.list[v.ID] = v
	}
	for _, v := range r.FundamentalType {
		t.list[v.ID] = v
	}
	for _, v := range r.ArrayType {
		t.list[v.ID] = v
	}
	for _, v := range r.PointerType {
		t.list[v.ID] = v
	}

	return t
}
