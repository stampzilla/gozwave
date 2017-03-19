package database

// GetMandatoryCommandClasses generates a list of mandatory commandclasses
func GetMandatoryCommandClasses(generic, specific byte) []*CommandClass {
	for k, v := range definitions {
		if k.Generic == generic && k.Specific == specific {
			return v
		}
	}
	return nil
}
