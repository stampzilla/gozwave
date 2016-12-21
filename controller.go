package gozwave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/functions"
	"github.com/stampzilla/gozwave/nodes"
	"github.com/stampzilla/gozwave/serialapi"
)

type Controller struct {
	Nodes      *nodes.List
	Connection *serialapi.Connection `json:"-"`
	eventQue   chan interface{}

	filename string
}

func (self *Controller) getNodes() (*nodes.List, error) {
	if self.eventQue == nil {
		self.eventQue = make(chan interface{}, 10)
	}
	resp := <-self.Connection.SendRaw([]byte{0x02}, time.Second*5)
	if resp == nil {
		return nil, fmt.Errorf("Send failed, timeut?")
	}

	switch r := resp.Data.(type) {
	case *functions.FuncDiscoveryNodes:
		for index, active := range r.ActiveNodes {
			if !active {
				continue
			}

			if node := self.Nodes.Get(index + 1); node != nil {
				node.Setup(self.Connection, self.eventQue)
				go node.Worker()
				continue
			}

			node := nodes.New((index + 1), self.Connection, self.eventQue)
			self.Nodes.Add(node)
			//<-basicIdentificationDone
		}
	default:
		logrus.Errorf("Got wrong response type: %t", r)
	}

	//spew.Dump(self)

	return self.Nodes, nil
}

func (self *Controller) PushEvent(event interface{}) {
	if self.eventQue == nil {
		self.eventQue = make(chan interface{}, 10)
	}

	select {
	case self.eventQue <- event:
	default:
	}
}

func (self *Controller) GetNextEvent() chan interface{} {
	if self.eventQue == nil {
		self.eventQue = make(chan interface{}, 10)
	}

	return self.eventQue
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
		self.Nodes.Add(v)
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
