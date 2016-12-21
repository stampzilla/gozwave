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

	Data interface{}
}

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
		if checksum != generateChecksum(checksumData) {
			logrus.Errorf("Invalid checksum, is %b should be %b, (len=%d)", checksum, generateChecksum(checksumData), length)
			logrus.Debug(hex.Dump(checksumData))
			return -1, nil
		}

		//logrus.Debug("Found SOF")

		return length + 2, CreateMessage(data)
	case 0x06: // ACK
		//logrus.Debug("Found ACK")
		return 1, nil
	case 0x15: // NAK
		logrus.Warn("Found NAK")
		return 1, nil
	case 0x18: // CAN
		logrus.Warn("Found CAN")
		return 1, nil
	default:
		return -1, nil // Not a valid start char
	}

	return
}

func generateChecksum(data []byte) byte {
	var offset int = 0
	var ret byte = data[offset]
	for i := offset + 1; i < len(data)-1; i++ {
		// Xor bytes
		ret ^= data[i]
	}
	// Not result
	ret = (byte)(^ret)
	return ret
}

func CreateMessage(data []byte) *Message {
	message := &Message{}
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
		logrus.Warnf("Dropping message function='%s' with data %x (not implemented)", message.Function, data[4:])
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

	return nil
}
