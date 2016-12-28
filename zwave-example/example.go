// An example that connects to an usb connected controller and starts a
// webserver at port 8080.
//
// http://localhost:8080/nodes
// http://localhost:8080/nodes/2
// http://localhost:8080/nodes/2/on
// http://localhost:8080/nodes/2/off

package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stampzilla/gozwave"
	"github.com/stampzilla/gozwave/events"
	"github.com/stampzilla/gozwave/nodes"
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
	r.GET("/nodes", getNodes(z))
	r.GET("/nodes/:id", getNode(z))
	r.GET("/nodes/:id/:action", control(z))
	go r.Run() // listen and serve on 0.0.0.0:8080

	for {
		select {
		case event := <-z.GetNextEvent():
			log.Println("----------------------------------------")
			log.Printf("Event: %#v\n", event)
			switch e := event.(type) {
			case events.NodeDiscoverd:
				znode := z.Nodes.Get(e.Address)
				znode.RLock()
				log.Printf("Node: %#v\n", znode)
				znode.RUnlock()

			case events.NodeUpdated:
				znode := z.Nodes.Get(e.Address)
				znode.RLock()
				log.Printf("Node: %#v\n", znode)
				znode.RUnlock()
			}
			log.Println("----------------------------------------")
		}
	}
}

func getNodes(z *gozwave.Controller) func(*gin.Context) {
	return func(c *gin.Context) {
		nodes := make(map[string]*nodes.Node)

		for k, v := range z.Nodes.All() {
			id := strconv.Itoa(k)
			nodes[id] = v
		}

		c.JSON(200, nodes)
	}
}

func getNode(z *gozwave.Controller) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(500, err)
			return
		}
		log.Println(id)
		device := z.Nodes.Get(id)
		if device == nil {
			c.JSON(404, "Device not found")
			return
		}

		c.JSON(200, device)
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
