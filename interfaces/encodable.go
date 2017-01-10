package interfaces

import (
	"time"

	"github.com/stampzilla/gozwave/commands/reports"
	"github.com/stampzilla/gozwave/serialapi"
)

type Encodable interface {
	Encode() []byte
}

type Writer interface {
	Write(Encodable) error
	WriteWithTimeout(Encodable, time.Duration) (<-chan *serialapi.Message, error)
	WriteAndWaitForReport(Encodable, time.Duration, byte) (<-chan reports.Report, error)
}

type LoadSaveable interface {
	LoadConfigurationFromFile() error
	SaveConfigurationToFile() error
}
