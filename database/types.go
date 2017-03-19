package database

import "github.com/stampzilla/gozwave/commands"

type CommandClass struct {
	ID         commands.ZWaveCommand
	Controlled string
	InNIF      string
	Secure     bool
	NonSecure  bool
	Version    string
}

type Definition struct {
	Generic  byte
	Specific byte
}
