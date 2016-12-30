package serialapi

import (
	"encoding/hex"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/functions"
)

type Message struct {
	Length      byte
	MessageType byte
	Function    functions.ZWaveFunction
	NodeId      byte

	startByte byte
	Data      interface{}
}

// IsACK checks if message is ACK type
func (m *Message) IsACK() bool {
	return m.startByte == 0x06
}

// IsNAK checks if message is NAK type
func (m *Message) IsNAK() bool {
	return m.startByte == 0x15
}

// IsCAN checks if message is CANCEL type
func (m *Message) IsCAN() bool {
	return m.startByte == 0x18
}

// Decode decodes serialapi raw data
func Decode(data []byte) (length int, msg *Message) {
	// SOF	0x01	1	Start Of Frame
	// ACK	0x06	6	Message Ack
	// NAK	0x15	21	Message NAK
	// CAN	0x18	24	Cancel - Resend request

	switch data[0] {
	case 0x01: // SOF
		if len(data) < 2 {
			return 2, nil
		}
		length := int(data[1])
		if len(data) < length+2 { // Make shure we have the full message
			return length + 2, nil
		}

		checksum := data[length+1]
		checksumData := append(data[1:length+1], byte(0x00))
		if checksum != GenerateChecksum(checksumData) {
			logrus.Errorf("Invalid checksum, is %b should be %x, (len=%d)", checksum, GenerateChecksum(checksumData), length)
			logrus.Debug(hex.Dump(checksumData))
			return -1, nil
		}

		//logrus.Debug("Found SOF")

		return length + 2, NewMessage(data)
	case 0x06, 0x15, 0x18: // ACK, NAC, CAN
		return 1, NewMessage(data)
	}
	return -1, nil // Not a valid start char
}

func GenerateChecksum(data []byte) byte {
	var offset int
	ret := data[offset]
	for i := offset + 1; i < len(data)-1; i++ {
		// Xor bytes
		ret ^= data[i]
	}
	// Not result
	ret = (byte)(^ret)
	return ret
}

func CompileMessage(data []byte) []byte {
	// Compile message
	msg := append([]byte{0x00, 0x00}, data...)
	msg = append(msg, 0x00)

	// Add length
	msg[0] = byte(len(msg) - 1)

	// Add checksum
	msg[len(msg)-1] = GenerateChecksum(msg)

	// Add header
	msg = append([]byte{0x01}, msg...)

	return msg
}

func NewMessage(data []byte) *Message {
	message := &Message{}
	message.startByte = data[0]
	if len(data) <= 1 {
		return message
	}
	message.Length = data[1]
	message.MessageType = data[2]
	message.Function = functions.ZWaveFunction(data[3])

	switch message.Function {
	case functions.ApplicationCommandHandler:
		message.NodeId = data[5]

		message.Data = functions.NewApplicationCommandHandler(data[7 : 7+data[6]])
	case functions.DiscoveryNodes:

		message.Data = functions.NewDiscoveryNodes(data[7 : 7+data[6]])
	case functions.GetNodeProtocolInfo:

		message.Data = functions.NewGetNodeProtocolInfo(data[4:])
	default:
		if len(data) > 5 {
			message.Data = data[5:]
		}
		logrus.Debugf("Dropping message function='%s' with data %x (not implemented)", message.Function, data[4:])
	}

	logrus.Debugf("Message: %+v Message data: %+v", message, message.Data)
	return message

	// 1 FrameHeader (type of message)
	// 1 Message length
	// 1 Request (00) or Response (01)
	// 1 Function
	//
	// 1 (Node ID)
	// 1 (Command - length)
	// - (Command)
	// - (Transmit options)
	// - (Callback id)
	// 1 Checksum

	// Type
	//Request = 0x00,
	//Response = 0x01,
	//GetVersion = 0x15,
	//MemoryGetId = 0x20,
	//ClockSet = 0x30

	// Transmit options
	//Ack = 0x01,
	//LowPower = 0x02,
	//AutoRoute = 0x04,
	//ForceRoute = 0x08

}
