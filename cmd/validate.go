package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"sonarci/sonar/sonarrestapi"
)

const (
	flagProject      = "project"
	flagProjectShort = "p"
	flagProjectUsage = "SonarQube projects key"
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

var pullRequestCmd = &cobra.Command{
	Use:   "pr [pull request id]",
	Short: "Validate pull request status",
	Long:  "Validate a pull request status on SonarQube.",
	Args:  cobra.MinimumNArgs(1),
	Run:   validatePullRequest,
}

func init() {
	branchCmd.Flags().StringP(flagProject, flagProjectShort, "", flagProjectUsage)
	_ = branchCmd.MarkFlagRequired(flagProject)

	pullRequestCmd.Flags().StringP(flagProject, flagProjectShort, "", flagProjectUsage)
	_ = pullRequestCmd.MarkFlagRequired(flagProject)

	validateCmd.AddCommand(branchCmd)
	validateCmd.AddCommand(pullRequestCmd)
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

func validatePullRequest(cmd *cobra.Command, args []string) {
	pr := args[0]

	project, _ := cmd.Flags().GetString(flagProject)
	if !validateFlag(flagProject, project) {
		return
	}

	pFlags := getPersistentFlagsFromCmd(cmd)
	if pFlags == nil {
		return
	}

	api := sonarrestapi.NewApi(pFlags.Server, pFlags.Token, pFlags.Timeout)
	err := api.ValidatePullRequest(project, pr)
	if err != nil {
		log.Fatal(err)
	}
}
