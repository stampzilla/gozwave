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
	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/events"
	"github.com/stampzilla/gozwave/functions"
	"github.com/stampzilla/gozwave/nodes"
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

func (self *Controller) getNodes() (*nodes.List, error) {
	t, err := self.Connection.WriteWithTimeout(functions.NewRaw([]byte{0x02}), time.Second*5)

	if err != nil {
		return nil, err
	}

	resp := <-t

	if resp == nil {
		return nil, fmt.Errorf("Timeout or invalid response")
	}

	switch r := resp.Data.(type) {
	case *functions.FuncDiscoveryNodes:
		for index, active := range r.ActiveNodes {
			if !active {
				continue
			}

			//skip first node (controller)
			if index == 0 {
				continue
			}

			node := self.Nodes.Get(index + 1)
			node = self.initNode(index+1, node)
			go node.Identify()
		}
	default:
		logrus.Errorf("Got wrong response type: %t", r)
	}

	//spew.Dump(self)

	return self.Nodes, nil
}

func (self *Controller) initNode(id int, node *nodes.Node) *nodes.Node {
	if node != nil {
		return node
	}

	node = nodes.New(id)
	self.Nodes.Add(node)
	self.pushEvent(events.NodeDiscoverd{
		Address: node.Id,
	})

	node.Setup(self.Connection, self.pushEvent)
	return node
}

func (self *Controller) pushEvent(event interface{}) {
	select {
	case self.eventQue <- event:
	default:
	}

	go func() {
		switch event.(type) {
		case events.NodeUpdated:
			self.triggerFileSave <- struct{}{}
		}
	}()
}

func (self *Controller) saveDebouncer() {
	<-self.triggerFileSave
	for {
		select {
		case <-self.triggerFileSave:
		case <-time.After(time.Second * 10):
			self.SaveConfigurationToFile()
			<-self.triggerFileSave
		}
	}
}

func (self *Controller) GetNextEvent() chan interface{} {
	self.RLock()
	defer self.RUnlock()

	return self.eventQue
}

func (c *Controller) DeliverReportToNode(node byte, report commands.Report) {
	n := c.Nodes.Get(int(node))
	if n == nil {
		return
	}
	n.ProcessEvent(report)
}

func (self *Controller) SaveConfigurationToFile() error {
	controller, err := self.loadConfigurationFromFile()
	if err == nil {
		if reflect.DeepEqual(self, controller) {
			return nil
		}
	}

	configFile, err := os.Create(self.filename)
	if err != nil {
		logrus.Error("creating config file", err.Error())
	}
	defer configFile.Close()

	logrus.Info("Save config: ", self.filename)
	var out bytes.Buffer
	b, err := json.MarshalIndent(self, "", "\t")
	if err != nil {
		logrus.Error("error marshal json", err)
	}
	json.Indent(&out, b, "", "\t")
	out.WriteTo(configFile)

	return nil
}
func (self *Controller) LoadConfigurationFromFile() error {
	controller, err := self.loadConfigurationFromFile()
	if err != nil {
		return err
	}

	for _, v := range controller.Nodes.All() {
		v.Setup(self.Connection, self.pushEvent)
		self.Nodes.Add(v)
		self.pushEvent(events.NodeDiscoverd{
			Address: v.Id,
		})
	}
	//self.Nodes = controller.Nodes

	return nil
}
func (self *Controller) loadConfigurationFromFile() (*Controller, error) {
	configFile, err := os.Open(self.filename)
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
