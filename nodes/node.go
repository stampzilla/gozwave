package nodes

import (
	"strings"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/commands/reports"
	"github.com/stampzilla/gozwave/database"
	"github.com/stampzilla/gozwave/events"
	"github.com/stampzilla/gozwave/interfaces"
	"github.com/stampzilla/gozwave/serialapi"
)

func New(address int) *Node {
	n := &Node{
		Id:            address,
		StateBool:     make(map[string]bool),
		StateFloat:    make(map[string]float64),
		pushEventFunc: func(interface{}) {},
	}

	//go n.Worker()

	return n
}

type Node struct {
	Id      int  `json:"id"`
	IsAwake bool `json:"is_awake"`

	protocolInfo        *serialapi.FuncGetNodeProtocolInfo
	ManufacurerSpecific *reports.ManufacturerSpecific

	//Device *database.Device
	Brand       string
	Product     string
	Description string

	CommandClasses []*database.CommandClass
	//TODO make public. move to database/types.go and use here
	//Parameters []*parameter

	Endpoints []*Endpoint

	StateBool  map[string]bool
	StateFloat map[string]float64
	statesOk   bool

	connection    interfaces.Writer
	pushEventFunc func(interface{})

	awake      chan struct{}
	identified bool

	sync.RWMutex
}

//TODO not use empty interface
//type ProcessEventFunc func(commands.Report)

func (n *Node) Setup(connection interfaces.Writer, pushEventFunc func(interface{})) {
	n.connection = connection
	n.pushEventFunc = pushEventFunc
}

func (n *Node) isAwake() chan struct{} {
	if n.IsAwake {
		c := make(chan struct{})
		close(c)
		return c
	}

	if n.awake == nil {
		n.awake = make(chan struct{})
	}

	return n.awake
}

func (n *Node) ProcessEvent(event reports.Report) {
	n.Lock()
	defer n.Unlock()

	switch data := event.(type) {
	case *reports.Alarm:
		if data.Status == 0xFF {
			n.StateBool["alarm"] = true
			if data.Type != 0 {
				n.StateFloat["alarm"] = float64(data.Type)
				n.StateFloat["alarmLastType"] = float64(data.Type)
			}
			n.pushEvent(events.NodeUpdated{
				Address: n.Id,
			})
			n.Unlock()

			<-time.After(time.Second)
			n.Lock()
			n.StateBool["alarm"] = false
			n.StateFloat["alarm"] = 0
			n.pushEvent(events.NodeUpdated{
				Address: n.Id,
			})
		}
	case *reports.WakeUp:
		logrus.Error("NODE RECEIVED WAKEUP")
		if n.awake != nil {
			close(n.awake)
			n.awake = nil
		}
	case *reports.SwitchBinaryV1:
		n.StateBool["on"] = data.CurrentValue
		n.pushEvent(events.NodeUpdated{
			Address: n.Id,
		})
	case *reports.SwitchBinaryV2:
		n.StateBool["on"] = data.CurrentValue
		n.pushEvent(events.NodeUpdated{
			Address: n.Id,
		})
	case *reports.SwitchMultilevelV1:
		n.StateBool["on"] = data.CurrentValue > 0
		n.StateFloat["level"] = float64(data.CurrentValue)
		n.pushEvent(events.NodeUpdated{
			Address: n.Id,
		})
	case *reports.SwitchMultilevelV4:
		n.StateBool["on"] = data.CurrentValue > 0
		n.StateFloat["level"] = float64(data.CurrentValue)
		n.pushEvent(events.NodeUpdated{
			Address: n.Id,
		})
	case *reports.SensorMultiLevel:
		//n.StateFloat[data.TypeString+" ("+data.Unit+")"] = data.Value
		key := strings.ToLower(data.TypeString + "_" + data.Unit)
		n.StateFloat[key] = data.Value
		n.pushEvent(events.NodeUpdated{
			Address: n.Id,
		})
	}

}

func (n *Node) Identify() {
	logrus.Infof("Started identification on node %d", n.Id)
	defer logrus.Infof("Ended identification on node %d", n.Id)

	for {
		if n.ProtocolInfo() == nil {
			logrus.Infof("Identify %d - Step 1 (RequestProtocolInfo)", n.Id)
			resp, err := n.RequestProtocolInfo()
			if err != nil {
				logrus.Errorf("Node ident: Failed RequestProtocolInfo: %s", err.Error())
				<-time.After(time.Second * 10)
				continue
			}

			n.Lock()
			n.protocolInfo = resp
			n.IsAwake = resp.Listening
			n.pushEvent(events.NodeUpdated{
				Address: n.Id,
			})
			n.Unlock()
		}

		// set basic commandClasses
		classes := database.GetMandatoryCommandClasses(n.ProtocolInfo().Generic, n.ProtocolInfo().Specific)
		desc := database.GetDescription(n.ProtocolInfo().Generic, n.ProtocolInfo().Specific)

		n.Lock()
		n.CommandClasses = classes
		n.Description = desc
		n.Unlock()

		//<-self.Connection.SendRaw([]byte{serialapi.GetNodeProtocolInfo, byte(index + 1)}) // Request node information
		//		nodeinfo := self.WaitForGetNodeProtocolInfo()

		// All manufacturer specific information such as vendor, vendors product ID and product type
		if n.ManufacurerSpecific == nil {
			<-n.isAwake()
			logrus.Infof("Identify %d - Step 2 (RequestManufacturerSpecific)", n.Id)
			resp, err := n.RequestManufacturerSpecific()
			if err != nil {
				logrus.Errorf("Node ident: Failed ManufacurerSpecific: %s", err.Error())
				<-time.After(time.Second * 10)
				continue
			}

			n.Lock()
			n.ManufacurerSpecific = resp
			dev := database.New(n.ManufacurerSpecific.Manufacturer, n.ManufacurerSpecific.Type, n.ManufacurerSpecific.ID)
			if dev != nil {
				n.Brand = dev.Brand
				n.Product = dev.Product
				n.Description = dev.Description
				for _, v := range dev.CommandClasses {
					n.CommandClasses = append(n.CommandClasses, v)
				}
			}
			n.pushEvent(events.NodeUpdated{
				Address: n.Id,
			})
			n.Unlock()
		}

		// Get node version
		//resp := <-n.connection.SendRaw([]byte{
		//serialapi.SendData, // Function
		//byte(n.Id),         // Node id
		//0x02,               // Length
		//commands.Version,   // Command
		//0x11,               // VersionCmd_Get
		//0x00,
		////0x05, // TransmitOptions?
		////0x23, // Callback?
		//}, time.Second*10) // Request node information
		//logrus.Println("%#v", resp)

		// associoation with controller
		//if n.Id == 5 {
		//logrus.Println("")
		//logrus.Println("")
		//resp := <-n.connection.SendRaw([]byte{
		//serialapi.SendData,   // Function
		//byte(n.Id),           // Node id
		//0x04,                 // Length
		//commands.Association, // Command
		//0x01,                 // AssociationCmd_Set
		//0x01,                 // Group
		//0x01,                 // Target node (controller)
		//0x25,                 // TransmitOptions?
		////0x23, // Callback?
		//}, time.Second*10) // Request node information
		//logrus.Println("")
		//logrus.Println("")
		//logrus.Println("%#v", resp)
		//}

		// Request node endpoints
		n.RLock()
		if n.Endpoints == nil && n.HasCommand(commands.MultiInstance) {
			n.RUnlock()
			<-n.isAwake()
			logrus.Infof("Identify %d - Step 3 (RequestEndpoints)", n.Id)
			err := n.RequestEndpoints()
			if err != nil {
				<-time.After(time.Second * 10)
				continue
			}

			n.RLock()
			n.pushEvent(events.NodeUpdated{
				Address: n.Id,
			})
			n.RUnlock()
			continue
		}
		n.RUnlock()

		n.RLock()
		if !n.statesOk {
			n.RUnlock()
			<-n.isAwake()
			logrus.Infof("Identify %d - Step 4 (RequestStates)", n.Id)
			err := n.RequestStates()
			if err != nil {
				<-time.After(time.Second * 10)
				continue
			}
			continue
		}
		n.RUnlock()

		//<-self.Connection.SendRaw([]byte{serialapi.IsFailedNode, byte(index + 1)}) // Request is failed node
		//	<-self.SendRaw([]byte{0xa0, byte(index + 1)}) // Request ?

		// All firmware and Z-Wave version information
		// All switching and reporting capabilities including current switching states and sensor values.
		// Number and maximum size of association groups
		// Configuration values if known
		// The Gateways will then do some initial setup:
		// If needed a certain predefined wakeup interval is set for battery operated devices
		// If available and requested certain standard defaults behaviours are configured in the device
		// If association groups are available the gateway will always put its own Node ID into these association groups in order to be informed about status changes and to be prepared for using associations for scene switching
		// The user can change all values. However it needs to be clear that particularly removing the gateways Node ID from the association groups may cause malfunctions of the gateway.

		n.RLock()
		n.pushEvent(events.NodeUpdated{
			Address: n.Id,
		})
		n.RUnlock()

		n.Lock()
		n.identified = true
		n.Unlock()
		return
	}
}

func (n *Node) pushEvent(event interface{}) {
	n.pushEventFunc(event)
}

func (n *Node) HasCommand(c commands.ZWaveCommand) bool {
	if n.CommandClasses == nil {
		return false
	}

	for _, v := range n.CommandClasses {
		if v.ID == c {
			return true
		}
	}

	return false
}

func (n *Node) IsDeviceClass(generic, specific byte) bool {
	if n.ProtocolInfo() == nil {
		return false
	}

	return n.ProtocolInfo().Generic == generic && n.ProtocolInfo().Specific == specific
}
