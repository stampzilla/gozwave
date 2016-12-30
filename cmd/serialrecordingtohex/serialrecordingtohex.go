package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"

	"github.com/stampzilla/gozwave/serialrecorder"
)

var file string
var short bool
var read bool
var write bool

func main() {
	flag.StringVar(&file, "file", "", "File to decode")
	flag.BoolVar(&short, "short", false, "Exclude timestamp, direction and writes")
	flag.BoolVar(&read, "read", true, "Include reads")
	flag.BoolVar(&write, "write", true, "Include writes")
	flag.Parse()

	if file == "" {
		log.Fatalln("No file specified. Use -file=")
		return
	}

	// Open the logfile
	file, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}

	// Read and decode all lines in the file
	dec := gob.NewDecoder(file)
	i := 0
	var rows []serialrecorder.Row
	for {
		// Save the line number so we can print a better error message
		i++

		// Decode the row
		var r serialrecorder.Row
		err = dec.Decode(&r)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("decode row %d: %s\n", i, err.Error())
		}

		rows = append(rows, r)
	}

	// Create a list of combined rows
	reads := logRowCombiner{}
	writes := logRowCombiner{}
	var combinedRows rowsByTimestamp
	for _, v := range rows {
		if v.Direction == serialrecorder.DirectionRead {
			c := reads.combine(v)
			if c != nil {
				combinedRows = append(combinedRows, *c)
			}
			continue
		}

		c := writes.combine(v)
		if c != nil {
			combinedRows = append(combinedRows, *c)
		}
	}

	// Sort the rows and print them out
	sort.Sort(combinedRows)
	for _, v := range combinedRows {
		if v.Direction.IsRead() && !read {
			continue
		}
		if v.Direction.IsWrite() && !write {
			continue
		}

		if short {
			fmt.Printf("%x\n", v.Data)
			continue
		}
		fmt.Printf("%s: %s - %x\n", v.Timestamp, v.Direction, v.Data)
	}
}

type row serialrecorder.Row
type rowsByTimestamp []row

func (c rowsByTimestamp) Len() int {
	return len(c)
}
func (c rowsByTimestamp) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c rowsByTimestamp) Less(i, j int) bool {
	return c[i].Timestamp.Before(c[j].Timestamp)
}

type logRowCombiner struct {
	prev serialrecorder.Row
	buf  []byte
}

func (l *logRowCombiner) combine(r serialrecorder.Row) *row {
	// If its the first row, set timestamp to same as first row
	if (l.prev.Timestamp == time.Time{}) {
		l.prev = r
	}

	// Calculate the duration from the last row
	duration := r.Timestamp.Sub(l.prev.Timestamp)

	// If duration is short, just add it to the buffer
	if duration < time.Millisecond {
		l.buf = append(l.buf, r.Data...)
		l.prev = r
		return nil
	}

	combinedRow := &row{
		Timestamp: l.prev.Timestamp,
		Direction: l.prev.Direction,
		Data:      l.buf,
	}

	l.buf = r.Data
	l.prev = r

	return combinedRow
}
