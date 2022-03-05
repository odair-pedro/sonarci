package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

const (
	flagValidateProject      = "project"
	flagValidateProjectShort = "p"
	flagValidateProjectUsage = "SonarQube projects key"
)

const (
	flagValidateAndDecorate      = "decorate"
	flagValidateAndDecorateShort = "d"
	flagValidateAndDecorateUsage = "Decorate a pull request with quality gate results"
)

func NewValidateCmd() *cobra.Command {
	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate quality gate status",
		Long:  "Validate a branch or pull request status on SonarQube.",
	}

	validateCmd.AddCommand(newBranchCmd())
	validateCmd.AddCommand(newPullRequestCmd())

	return validateCmd
}

func newBranchCmd() *cobra.Command {
	branchCmd := &cobra.Command{
		Use:   "branch [branch name]",
		Short: "Validate branch status",
		Long:  "Validate a branch status on SonarQube.",
		Args:  cobra.MinimumNArgs(1),
		Run:   validateBranch,
	}

	branchCmd.Flags().StringP(flagValidateProject, flagValidateProjectShort, "", flagValidateProjectUsage)
	_ = branchCmd.MarkFlagRequired(flagValidateProject)

	return branchCmd
}

func newPullRequestCmd() *cobra.Command {
	pullRequestCmd := &cobra.Command{
		Use:   "pr [pull request id]",
		Short: "Validate pull request status",
		Long:  "Validate a pull request status on SonarQube.",
		Args:  cobra.MinimumNArgs(1),
		Run:   validatePullRequest,
	}

	pullRequestCmd.Flags().StringP(flagValidateProject, flagValidateProjectShort, "", flagValidateProjectUsage)
	pullRequestCmd.Flags().BoolP(flagValidateAndDecorate, flagValidateAndDecorateShort, false, flagValidateAndDecorateUsage)
	_ = pullRequestCmd.MarkFlagRequired(flagValidateProject)

	pullRequestCmd.Flags().String(flagDecorateTag, "", flagDecorateTagUsage)

	return pullRequestCmd
}

func validateBranch(cmd *cobra.Command, args []string) {
	branch := args[0]

	project, _ := cmd.Flags().GetString(flagValidateProject)
	if !validateFlag(flagValidateProject, project) {
		return
	}

	pFlags := getPersistentFlagsFromCmd(cmd)
	if pFlags == nil {
		return
	}

	api := createSonarApi(pFlags.Server, pFlags.Token, pFlags.Timeout)

	qualityGate, err := api.GetBranchQualityGate(project, branch)
	if err != nil {
		log.Fatal(err)
	}

	if !checkQualityGate(qualityGate) {
		os.Exit(1)
	}
}

func validatePullRequest(cmd *cobra.Command, args []string) {
	pr := args[0]

	project, _ := cmd.Flags().GetString(flagValidateProject)
	if !validateFlag(flagValidateProject, project) {
		return
	}

	tag, _ := cmd.Flags().GetString(flagDecorateTag)

	pFlags := getPersistentFlagsFromCmd(cmd)
	if pFlags == nil {
		return
	}

	api := createSonarApi(pFlags.Server, pFlags.Token, pFlags.Timeout)

	qualityGate, err := api.GetPullRequestQualityGate(project, pr)
	if err != nil {
		log.Fatal(err)
	}

	decorate, _ := cmd.Flags().GetBool(flagValidateAndDecorate)
	if decorate {
		decoratePullRequest(qualityGate, tag, pFlags.Timeout)
	}

	if !checkQualityGate(qualityGate) {
		os.Exit(1)
	}
}
