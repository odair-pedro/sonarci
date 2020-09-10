package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var qualityGateCmd = &cobra.Command{
	Use:   "quality",
	Short: "Quality short description",
	Long:  "Quality long description",
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("quality was called")
	},
}

func init() {
	qualityGateCmd.Flags().StringP("pull-request", "p", "", "Usage for pull-request")
}
