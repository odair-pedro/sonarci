package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"sonarci/sonar/rest/factory"
)

func NewServerVersionCmd() *cobra.Command {
	serverVersionCmd := &cobra.Command{
		Use:   "server-version",
		Short: "Get SonarQube server version",
		Long:  "Get SonarQube server version",
		Run:   serverVersion,
	}

	return serverVersionCmd
}

func serverVersion(cmd *cobra.Command, args []string) {
	_ = args

	pFlags := getPersistentFlagsFromCmd(cmd)
	if pFlags == nil {
		return
	}

	api := factory.GetLatestSonarApi(pFlags.Server, pFlags.Token, pFlags.Timeout)
	version, err := api.GetServerVersion()
	if err != nil {
		log.Fatal("Failure to get server version: ", err)
	}

	log.Println(version)
}
