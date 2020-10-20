package main

import (
	"log"
	"sonarci/cmd"
	"sonarci/logging"
)

var version string

func main() {
	logging.Setup()
	startRootCommand()
}

func startRootCommand() {
	rootCmd := cmd.NewRootCmd()
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
