package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sonarqubeci",
	Short: "A simple tool for SonarQube integration",
	Long:  "SonarQubeFast is a CLI library for help you integrate and use SonarQube inspections.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("username", "u", "", "Username and password or a Token for authentication")
	rootCmd.PersistentFlags().StringP("server", "s", "", "SonarQube server address")
	rootCmd.PersistentFlags().StringP("project", "p", "", "SonarQube project key")

	rootCmd.MarkPersistentFlagRequired("server")
	rootCmd.MarkPersistentFlagRequired("project")

	rootCmd.AddCommand(qualityGateCmd)
	rootCmd.AddCommand(searchCmd)
}
