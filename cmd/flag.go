package cmd

import "log"

func validateFlag(flagName string, flagValue interface{}) bool {
	var isValid bool

	switch flagValue.(type) {
	case int:
		isValid = flagValue.(int) > 0
	case string:
		isValid = flagValue.(string) != ""
	}

	if !isValid {
		log.Println("Invalid value for flag: " + flagName)
	}

	return isValid
}
