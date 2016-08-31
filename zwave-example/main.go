package main

import (
	"flag"
	"log"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave"
)

func main() {

	//direction := flag.Bool("d", false, "true == down, false == up")
	debug := flag.Bool("v", false, "Verbose - show more debuging info")
	flag.Parse()

	logrus.SetLevel(logrus.WarnLevel)
	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// We communicate with the z-stick thru an "serial port"
	// on ubuntu it showes up as an ACM interface
	_, err := gozwave.Connect("/dev/ttyACM0")
	if err != nil {
		log.Println(err)
		return
	}

	//nodes, _ := z.GetNodes()
	//for _, node := range nodes.All() {
	//fmt.Printf("%#v\n", node)
	//}

	//switch node := nodes.Get(2).(type) {
	//case Blind:
	//node.Up()
	//}
	select {}

	//rollup := switchbinary.New().SetNode(2).SetValue(true)
	//for {
	//logrus.Warn("UP")
	//rollup.SetValue(false)
	//<-z.Send(rollup) // Stop previous motion
	//<-time.After(time.Millisecond * 200)
	//<-z.Send(rollup) // Start up

	//<-time.After(time.Second * 5)

	//logrus.Warn("DOWN")
	//rollup.SetValue(true)
	//<-z.Send(rollup)
	//<-time.After(time.Millisecond * 200)
	//<-z.Send(rollup)
	//<-time.After(time.Second * 5)
	//}

}
