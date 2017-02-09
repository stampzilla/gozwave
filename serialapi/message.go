package serialapi

import (
	"fmt"

	"github.com/Sirupsen/logrus"
)

// Message is the decoded message received from the controller
type Message struct {
	Length      byte
	MessageType byte
	Function    ZWaveFunction
	NodeID      byte

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
func Decode(data []byte) (length int, msg *Message, err error) {
	// SOF	0x01	1	Start Of Frame
	// ACK	0x06	6	Message Ack
	// NAK	0x15	21	Message NAK
	// CAN	0x18	24	Cancel - Resend request

	switch data[0] {
	case 0x01: // SOF
		if len(data) < 2 {
			return 2, nil, nil
		}
		length := int(data[1])
		if len(data) < length+2 { // Make shure we have the full message
			return length + 2, nil, nil
		}

		checksum := data[length+1]
		checksumData := append(data[1:length+1], byte(0x00))
		if checksum != GenerateChecksum(checksumData) {
			return -1, nil, fmt.Errorf("Invalid checksum, is 0x%x should be 0x%x, (len=%d)", checksum, GenerateChecksum(checksumData), length)
		}

		msg, err := NewMessage(data)

		return length + 2, msg, err
	case 0x06, 0x15, 0x18: // ACK, NAC, CAN
		msg, err := NewMessage(data)
		if err != nil {
			logrus.Error(err)
		}

		return 1, msg, nil
	}
	return -1, nil, nil // Not a valid start char
}

// GenerateChecksum calculates the checksum for a serialapi message
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

// CompileMessage is used to calculate length and checksum for messages that are going to be sent to the controller
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

// NewMessage tries to decode a binary message from the controller in to a message struct
func NewMessage(data []byte) (*Message, error) {
	message := &Message{}
	message.startByte = data[0]
	if len(data) <= 1 {
		return message, nil
	}
	message.Length = data[1]
	message.MessageType = data[2]
	message.Function = ZWaveFunction(data[3])

	var err error
	switch message.Function {
	case ApplicationCommandHandler:
		message.NodeID = data[5]

		message.Data, err = NewApplicationCommandHandler(data[7 : 7+data[6]])
	case DiscoveryNodes:

		message.Data, err = NewDiscoverdNodes(data[7 : 7+data[6]])
	case GetNodeProtocolInfo:

		message.Data, err = NewGetNodeProtocolInfo(data[4:])
	default:
		err = fmt.Errorf("Dropping message function='%s' with data %x (not implemented)", message.Function, data[4:])
	}

	if err != nil {
		return nil, err
	}

	logrus.Debugf("Message: %+v Message data: %+v", message, message.Data)
	return message, err

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
