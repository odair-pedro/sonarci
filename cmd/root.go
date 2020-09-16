package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	flagServer = "server"
	flagToken  = "token"
	flagTimout = "timeout"

	timeoutDefault = 30000
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
	rootCmd.PersistentFlags().StringP(flagServer, "s", "", "SonarQube server address")
	rootCmd.PersistentFlags().StringP(flagToken, "o", "", "Authentication Token")
	rootCmd.PersistentFlags().IntP(flagTimout, "t", 0, fmt.Sprintf("Timeout in milliseconds. Default value is %d ms", timeoutDefault))

	_ = rootCmd.MarkPersistentFlagRequired(flagServer)
	_ = rootCmd.MarkPersistentFlagRequired(flagToken)

	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(versionCmd)
}
