package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stampzilla/gozwave"
	"github.com/stampzilla/gozwave/events"
)

var port string

func init() {
	flag.StringVar(&port, "port", "/dev/ttyACM0", "SerialAPI communication port (to controller)")
	flag.Parse()

}

func main() {
	z, err := gozwave.Connect(port, "")
	if err != nil {
		log.Println(err)
	}

	r := gin.Default()
	r.GET("/nodes", nodes(z))
	r.GET("/nodes/:id/:action", control(z))
	go r.Run() // listen and serve on 0.0.0.0:8080

	for {
		select {
		case event := <-z.GetNextEvent():
			log.Printf("Event: %#v\n", event)
			switch e := event.(type) {
			case events.NodeDiscoverd:
				znode := z.Nodes.Get(e.Address)
				znode.RLock()
				log.Printf("%#v\n", znode)
				znode.RUnlock()

			case events.NodeUpdated:
				znode := z.Nodes.Get(e.Address)
				znode.RLock()
				log.Printf("%#v\n", znode)
				znode.RUnlock()
			}
		}
	}
}

func nodes(z *gozwave.Controller) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, z.Nodes.All())
	}
}

func control(z *gozwave.Controller) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(500, err)
			return
		}

		device := z.Nodes.Get(id)
		if device == nil {
			c.JSON(404, "Device not found")
			return
		}

		switch c.Param("action") {
		case "on":
			device.On()
		case "off":
			device := z.Nodes.Get(id)
			device.Off()
		default:
			c.JSON(404, "Action not found")
			return

		}

		c.AbortWithStatus(200)
	}
}
