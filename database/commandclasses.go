package database

import "fmt"

// GetMandatoryCommandClasses generates a list of mandatory commandclasses
func GetMandatoryCommandClasses(generic, specific byte) []*CommandClass {
	for k, v := range definitions {
		if k.Generic == generic && k.Specific == specific {
			return v
		}
	}
	return nil
}

// GetDescription returns a generic description of the device
func GetDescription(generic, specific byte) string {
	for k, v := range descriptions {
		if k.Generic == generic && k.Specific == specific {
			return v
		}
	}
	return fmt.Sprintf("Unknown (generic=0x%X specific=0x%X)", generic, specific)
}
