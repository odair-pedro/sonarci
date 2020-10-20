package main

import (
	"log"
	"sonarci/cmd"
	"sonarci/net/http"
	"sonarci/sonar/rest/base"
	"time"
)

var version string

func main() {
	//logging.Setup()
	//startRootCommand()

	api := base.NewRestApi(http.NewConnection("https://sonarqube.ambevdevs.com.br", "00500edbf41ebdc63f48646a5e86562fcc0b9692", time.Second*30))
	result, _ := api.GetPullRequestQualityGate("ambev-franqueado-gaming-server", "23195")
	_ = result
}

func startRootCommand() {
	rootCmd := cmd.NewRootCmd()
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
