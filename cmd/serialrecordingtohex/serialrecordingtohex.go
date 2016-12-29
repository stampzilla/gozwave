package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/stampzilla/gozwave/serialrecorder"
)

var file string

func init() {
	flag.StringVar(&file, "file", "", "File to decode")
	flag.Parse()
}

func main() {
	if file == "" {
		log.Fatalln("No file specified. Use -file=")
		return
	}

	file, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}

	dec := gob.NewDecoder(file)

	i := 0
	var rows []serialrecorder.Row

	for {
		var r serialrecorder.Row

		i++
		err = dec.Decode(&r)
		rows = append(rows, r)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("decode row %d: %s\n", i, err.Error())
		}
	}

	for _, v := range rows {
		fmt.Printf("%s: %s - %x\n", v.Timestamp, v.Direction, v.Data)
	}
}
