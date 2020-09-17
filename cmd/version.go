package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"sonarci/sonar"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get SonarQube server version",
	Long:  "Get SonarQube server version",
	Run:   version,
}

func version(cmd *cobra.Command, args []string) {
	_ = args

	pFlags := getPersistentFlagsFromCmd(cmd)
	if pFlags == nil {
		return
	}

	api := sonar.MakeApi(pFlags.Server, pFlags.Token, pFlags.Timeout)
	version, err := api.GetServerVersion()
	if err != nil {
		log.Fatalln("Failure to get server version: ", err)
	}

	log.Println(version)
}
