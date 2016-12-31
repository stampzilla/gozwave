package gozwave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/commands/reports"
	"github.com/stampzilla/gozwave/events"
	"github.com/stampzilla/gozwave/nodes"
	"github.com/stampzilla/gozwave/serialapi"
)

type Controller struct {
	Nodes      *nodes.List
	Connection *Connection `json:"-"`
	eventQue   chan interface{}

	filename        string
	triggerFileSave chan struct{}

	sync.RWMutex
}

func NewController() *Controller {
	c := &Controller{
		Nodes:           nodes.NewList(),
		Connection:      NewConnection(),
		eventQue:        make(chan interface{}, 10),
		triggerFileSave: make(chan struct{}),
	}
	return c
}

func (c *Controller) getNodes() (*nodes.List, error) {
	t, err := c.Connection.WriteWithTimeout(serialapi.NewRaw([]byte{0x02}), time.Second*5)

	if err != nil {
		return nil, err
	}

	resp := <-t

	if resp == nil {
		return nil, fmt.Errorf("Timeout or invalid response")
	}

	switch r := resp.Data.(type) {
	case *serialapi.DiscoverdNodes:
		for index, active := range r.ActiveNodes {
			if !active {
				continue
			}

			//skip first node (controller)
			if index == 0 {
				continue
			}

			node := c.Nodes.Get(index + 1)
			node = c.initNode(index+1, node)
			go node.Identify()
		}
	default:
		logrus.Errorf("Got wrong response type: %t", r)
	}

	return c.Nodes, nil
}

func (c *Controller) initNode(id int, node *nodes.Node) *nodes.Node {
	if node != nil {
		return node
	}

	node = nodes.New(id)
	c.Nodes.Add(node)
	c.pushEvent(events.NodeDiscoverd{
		Address: node.Id,
	})

	node.Setup(c.Connection, c.pushEvent)
	return node
}

func (c *Controller) pushEvent(event interface{}) {
	select {
	case c.eventQue <- event:
	default:
	}

	go func() {
		switch event.(type) {
		case events.NodeUpdated:
			c.triggerFileSave <- struct{}{}
		}
	}()
}

func (c *Controller) saveDebouncer() {
	<-c.triggerFileSave
	for {
		select {
		case <-c.triggerFileSave:
		case <-time.After(time.Second * 10):
			c.SaveConfigurationToFile()
			<-c.triggerFileSave
		}
	}
}

func (c *Controller) GetNextEvent() chan interface{} {
	c.RLock()
	defer c.RUnlock()

	return c.eventQue
}

func (c *Controller) DeliverReportToNode(node byte, report reports.Report) {
	n := c.Nodes.Get(int(node))
	if n == nil {
		return
	}
	n.ProcessEvent(report)
}

func (c *Controller) SaveConfigurationToFile() error {
	controller, err := c.loadConfigurationFromFile()
	if err == nil {
		if reflect.DeepEqual(c, controller) {
			return nil
		}
	}

	configFile, err := os.Create(c.filename)
	if err != nil {
		logrus.Error("creating config file", err.Error())
	}
	defer configFile.Close()

	logrus.Info("Save config: ", c.filename)
	var out bytes.Buffer
	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		logrus.Error("error marshal json", err)
	}
	json.Indent(&out, b, "", "\t")
	out.WriteTo(configFile)

	return nil
}
func (c *Controller) LoadConfigurationFromFile() error {
	controller, err := c.loadConfigurationFromFile()
	if err != nil {
		return err
	}

	for _, v := range controller.Nodes.All() {
		v.Setup(c.Connection, c.pushEvent)
		c.Nodes.Add(v)
		c.pushEvent(events.NodeDiscoverd{
			Address: v.Id,
		})
	}

	return nil
}
func (c *Controller) loadConfigurationFromFile() (*Controller, error) {
	configFile, err := os.Open(c.filename)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	controller := &Controller{}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&controller); err != nil {
		return nil, err
	}

	return controller, nil
}
