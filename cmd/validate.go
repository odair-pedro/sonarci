package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sonarci/sonar"
	"sort"
	"strconv"
	"strings"
)

const (
	flagProject      = "project"
	flagProjectShort = "p"
	flagProjectUsage = "SonarQube projects key"
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

	branchCmd.Flags().StringP(flagProject, flagProjectShort, "", flagProjectUsage)
	_ = branchCmd.MarkFlagRequired(flagProject)

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

	pullRequestCmd.Flags().StringP(flagProject, flagProjectShort, "", flagProjectUsage)
	_ = pullRequestCmd.MarkFlagRequired(flagProject)

	return pullRequestCmd
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

	api := createSonarApi(pFlags.Server, pFlags.Token, pFlags.Timeout)

	qualityGate, err := api.GetBranchQualityGate(project, branch)
	if err != nil {
		log.Fatal(err)
	}

	checkQualityGate(qualityGate)
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

	api := createSonarApi(pFlags.Server, pFlags.Token, pFlags.Timeout)

	qualityGate, err := api.GetPullRequestQualityGate(project, pr)
	if err != nil {
		log.Fatal(err)
	}

	checkQualityGate(qualityGate)
}

func checkQualityGate(qualityGate sonar.QualityGate) {
	const banner = " ____                            ____ ___ \n" +
		"/ ___|  ___  _ __   __ _ _ __   / ___|_ _|\n" +
		"\\___ \\ / _ \\| '_ \\ / _  | '__| | |    | | \n" +
		" ___) | (_) | | | | (_| | |    | |___ | | \n" +
		"|____/ \\___/|_| |_|\\__,_|_|     \\____|___|\n\n"

	log.Print(banner)
	log.Println(genQualityReport(qualityGate))
	log.Printf("\nSee more details in %s", qualityGate.LinkDetail)

	if !qualityGate.HasPassed() {
		os.Exit(1)
	}
}

func genQualityReport(qualityGate sonar.QualityGate) string {
	const (
		metricColW     = 28
		comparatorColW = 10
		errorColW      = 15
		valueColW      = 12
		statusColW     = 6
	)

	header := "+------------------------------+------------+-----------------+--------------+--------+\n" +
		"| METRIC                       | COMPARATOR | ERROR THRESHOLD | ACTUAL VALUE | STATUS |\n" +
		"+------------------------------+------------+-----------------+--------------+--------+\n"
	footer := "+------------------------------+------------+-----------------+--------------+--------+\n" +
		fmt.Sprintf("|                              |            |                 | QUALITY GATE | %s |\n",
			colorful(qualityGate.Status, padRight(qualityGate.Status, " ", statusColW))) +
		"+------------------------------+------------+-----------------+--------------+--------+"

	keys := make([]string, 0)
	for k := range qualityGate.Conditions {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var rows string
	for _, key := range keys {
		metric := qualityGate.Conditions[key]
		rows += fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
			colorful(metric.Status, padRight(metric.Description, " ", metricColW)),
			colorful(metric.Status, padRight(metric.Comparator, " ", comparatorColW)),
			colorful(metric.Status, padRight(strconv.FormatFloat(float64(metric.ErrorThreshold), 'f', 5, 32), " ", errorColW)),
			colorful(metric.Status, padRight(strconv.FormatFloat(float64(metric.Value), 'f', 5, 32), " ", valueColW)),
			colorful(metric.Status, padRight(metric.Status, " ", statusColW)))
	}

	return header + rows + footer
}

func colorful(status string, value string) string {
	const (
		colorReset = "\033[0m"
		colorRed   = "\033[31m"
		colorGreen = "\033[32m"
	)

	status = strings.ToUpper(strings.Trim(status, " "))
	if status == "OK" {
		return fmt.Sprint(colorGreen, value, colorReset)
	}
	if status == "ERROR" {
		return fmt.Sprint(colorRed, value, colorReset)
	}
	return value
}
