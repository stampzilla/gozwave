package nodes

type ManufacurerSpecific struct {
}

func (n *node) RequestManufacturerSpecific() error {
	// Todo: Send raw messages here

	//resp := <-n.connection.SendRaw([]byte{
	//functions.SendData, // Function
	//byte(index + 1),    // Node id
	//0x02,               // Length
	//commands.ManufacturerSpecific, // Command
	//0x04, // MANUFACTURER_SPECIFIC_GET
	//0x05, // TransmitOptions?
	//0x23, // Callback?
	//}) // Request node information

	//logrus.Infof("RESP: %#v", resp)

	return nil
}
