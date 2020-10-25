package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sonarci/connection/http"
	"sonarci/sonar"
	sonarFactory "sonarci/sonar/factory"
	"time"
)

const (
	flagServer     = "server"
	flagToken      = "token"
	flagTimout     = "timeout"
	timeoutDefault = 30000
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "sonarci",
		Short: "A simple tool for SonarQube integration",
		Long:  "SonarCI is a CLI library for help you integrate and use SonarQube inspections.",
	}

	rootCmd.PersistentFlags().StringP(flagServer, "s", "", "SonarQube server address")
	rootCmd.PersistentFlags().StringP(flagToken, "o", "", "Authentication Token")
	rootCmd.PersistentFlags().IntP(flagTimout, "t", 0, fmt.Sprintf("Timeout in milliseconds. Default value is %d ms", timeoutDefault))

	_ = rootCmd.MarkPersistentFlagRequired(flagServer)
	_ = rootCmd.MarkPersistentFlagRequired(flagToken)

	rootCmd.AddCommand(NewServerVersionCmd())
	rootCmd.AddCommand(NewSearchCmd())
	rootCmd.AddCommand(NewValidateCmd())

	return rootCmd
}

func createSonarApi(server string, token string, timeout time.Duration) sonar.Api {
	conn := http.NewConnection(server, token, timeout)
	return sonarFactory.CreateSonarApi(conn)
}
