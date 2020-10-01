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
	cmd.SetVersion(version)
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
