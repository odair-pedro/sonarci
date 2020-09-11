package main

import (
	"log"
	"os"
	"sonarci/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
