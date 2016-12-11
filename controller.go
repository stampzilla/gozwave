package gozwave

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/coreos/fleet/log"
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
	/*
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

				node, basicIdentificationDone := nodes.New((index + 1), self.Connection, self.eventQue)
				self.Nodes.Add(node)
				<-basicIdentificationDone
			}
		default:
			logrus.Errorf("Got wrong response type: %t", r)
		}*/

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
	return nil

	configFile, err := os.Create(self.filename)
	if err != nil {
		log.Error("creating config file", err.Error())
	}

	log.Info("Save config: ", self.filename)
	var out bytes.Buffer
	b, err := json.MarshalIndent(self, "", "\t")
	if err != nil {
		log.Error("error marshal json", err)
	}
	json.Indent(&out, b, "", "\t")
	out.WriteTo(configFile)

	return nil
}
func (self *Controller) LoadConfigurationFromFile() error {
	configFile, err := os.Open(self.filename)
	if err != nil {
		return err
	}

	controller := &Controller{}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&controller); err != nil {
		return err
	}

	for _, v := range controller.Nodes.All() {
		self.Nodes.Add(v)
	}
	//self.Nodes = controller.Nodes

	return nil
}
