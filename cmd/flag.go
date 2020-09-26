package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"time"
)

type persistentFlags struct {
	Server  string
	Token   string
	Timeout time.Duration
}

func getPersistentFlagsFromCmd(cmd *cobra.Command) *persistentFlags {
	flags := cmd.Root().PersistentFlags()
	server, _ := flags.GetString(flagServer)
	if !validateFlag(flagToken, server) {
		return nil
	}

	token, _ := flags.GetString(flagToken)
	if !validateFlag(flagToken, token) {
		return nil
	}

	timeout, _ := flags.GetInt(flagTimout)
	if timeout == 0 {
		timeout = timeoutDefault
	}

	return &persistentFlags{Server: server, Token: token, Timeout: time.Duration(timeout) * time.Millisecond}
}

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
