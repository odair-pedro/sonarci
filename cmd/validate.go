package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"sonarci/sonar/sonarrestapi"
)

const (
	flagProject = "project"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate quality gate status",
	Long:  "Validate a branch or pull request status on SonarQube.",
}

var branchCmd = &cobra.Command{
	Use:   "branch [branch name]",
	Short: "Validate branch status",
	Long:  "Validate a branch status on SonarQube.",
	Args:  cobra.MinimumNArgs(1),
	Run:   validateBranch,
}

func init() {
	branchCmd.Flags().StringP(flagProject, "p", "", "SonarQube projects key")
	_ = branchCmd.MarkFlagRequired(flagProject)

	validateCmd.AddCommand(branchCmd)
}

func validateBranch(cmd *cobra.Command, args []string) {
	branch := args[0]

	project, _ := cmd.Flags().GetString(flagProject)
	if !validateFlag(flagProject, project) {
		return
	}

	pFlags := getPersistentFlagsFromCmd(cmd)
	if pFlags == nil {
		return
	}

	api := sonarrestapi.NewApi(pFlags.Server, pFlags.Token, pFlags.Timeout)
	err := api.ValidateBranch(project, branch)
	if err != nil {
		log.Fatal(err)
	}
}
