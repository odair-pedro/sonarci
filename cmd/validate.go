package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	decorationFactory "sonarci/decoration/factory"
	templateFactory "sonarci/decoration/template/factory"
	"sonarci/sonar"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	flagProject      = "project"
	flagProjectShort = "p"
	flagProjectUsage = "SonarQube projects key"
)

const (
	flagDecorate      = "decorate"
	flagDecorateShort = "d"
	flagDecorateUsage = "Decorate a pull request with quality gate results"
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
	pullRequestCmd.Flags().BoolP(flagDecorate, flagDecorateShort, false, flagDecorateUsage)
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

	decorate, _ := cmd.Flags().GetBool(flagDecorate)
	if decorate {
		decoratePullRequest(qualityGate, pFlags.Timeout)
	}

	checkQualityGate(qualityGate)
}

func decoratePullRequest(qualityGate sonar.QualityGate, timeout time.Duration) {
	const (
		decoratorTypeEnv = "SONARCI_DECORATION_TYPE"
		projectEnv       = "SONARCI_DECORATION_PROJECT"
		repositoryEnv    = "SONARCI_DECORATION_REPOSITORY"
		tokenEnv         = "SONARCI_DECORATION_TOKEN"
	)

	decoratorType := os.Getenv(decoratorTypeEnv)
	if decoratorType == "" {
		log.Printf("Failed decoration, decorator type has not been found")
		return
	}

	project := os.Getenv(projectEnv)
	if project == "" {
		log.Printf("Failed decoration, project information has not been found")
		return
	}

	repository := os.Getenv(repositoryEnv)
	if repository == "" {
		log.Printf("Failed decoration, repository information has not been found")
		return
	}

	token := os.Getenv(tokenEnv)
	if token == "" {
		log.Printf("Failed decoration, token information has not been found")
		return
	}

	engine := templateFactory.CreateDummyTemplateEngine()
	decorator, err := decorationFactory.CreatePullRequestDecorator(decoratorType, project, repository, token, timeout, engine)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	err = decorator.CommentQualityGate(qualityGate)
	if err != nil {
		log.Printf(err.Error())
	}
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
