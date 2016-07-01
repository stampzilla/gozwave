package main

import "github.com/stampzilla/gozwave"

func main() {

	// We communicate with the z-stick thru an "serial port"
	// on ubuntu it showes up as an ACM interface
	z := gozwave.New("/dev/ttyACM0")

	z.Connect()

	z.GetNodes()

	select {}

}
