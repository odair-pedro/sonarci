package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sonarqubeci",
	Short: "A tool for SonarQube integration",
	Long:  "SonarQubeFast is a CLI library for help you integrate and use SonarQube inspections.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//log.Print("init was called")

	rootCmd.PersistentFlags().StringP("username", "u", "", "Test usage")

	rootCmd.AddCommand(qualityGateCmd)
}
