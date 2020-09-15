package main

import (
	"log"
	"sonarci/cmd"
	"sonarci/logging"
)

func main() {
	logging.Setup()
	startRootCommand()
}

func startRootCommand() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
