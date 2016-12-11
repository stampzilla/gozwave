package interfaces

type Encodable interface {
	Encode() []byte
}

type LoadSaveable interface {
	LoadConfigurationFromFile() error
	SaveConfigurationToFile() error
}
