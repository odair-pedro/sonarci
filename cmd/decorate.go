package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"sonarci/connection"
	"sonarci/connection/http"
	decorationFactory "sonarci/decoration/factory"
	templateFactory "sonarci/decoration/template/factory"
	"sonarci/sonar"
	"time"
)

const (
	flagDecorateProject      = "project"
	flagDecorateProjectShort = "p"
	flagDecorateProjectUsage = "SonarQube projects key"
)

const (
	flagDecoratePullRequest      = "pull-request"
	flagDecoratePullRequestShort = "r"
	flagDecoratePullRequestUsage = "Pull request ID"
)

func NewDecorateCmd() *cobra.Command {
	decorateCmd := &cobra.Command{
		Use:   "decorate",
		Short: "Decorate pull request with the quality gate report",
		Long:  "Decorate pull request with the SonarQube's quality gate report.",
		Run:   decorate,
	}

	decorateCmd.Flags().StringP(flagDecorateProject, flagDecorateProjectShort, "", flagDecorateProjectUsage)
	_ = decorateCmd.MarkFlagRequired(flagDecorateProject)

	decorateCmd.Flags().StringP(flagDecoratePullRequest, flagDecoratePullRequestShort, "", flagDecoratePullRequestUsage)
	_ = decorateCmd.MarkFlagRequired(flagDecoratePullRequest)

	return decorateCmd
}

func decorate(cmd *cobra.Command, _ []string) {
	pr, _ := cmd.Flags().GetString(flagDecoratePullRequest)
	if !validateFlag(flagDecoratePullRequest, pr) {
		return
	}

	project, _ := cmd.Flags().GetString(flagDecorateProject)
	if !validateFlag(flagDecorateProject, project) {
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

	decoratePullRequest(qualityGate, pFlags.Timeout)
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
		log.Print("Failed decoration, decorator type has not been found")
		return
	}

	project := os.Getenv(projectEnv)
	if project == "" {
		log.Print("Failed decoration, project information has not been found")
		return
	}

	repository := os.Getenv(repositoryEnv)
	if repository == "" {
		log.Print("Failed decoration, repository information has not been found")
		return
	}

	token := os.Getenv(tokenEnv)
	if token == "" {
		log.Print("Failed decoration, token information has not been found")
		return
	}

	engine := templateFactory.CreateDummyTemplateEngine()
	decorator, err := decorationFactory.CreatePullRequestDecorator(decoratorType, project, repository, engine,
		func(server string) connection.Connection {
			return http.NewConnection(server, token, timeout)
		})
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = decorator.ClearPreviousComments(qualityGate.Source)
	if err != nil {
		log.Printf("Failue at remove old comments from pull request (%s): %s", qualityGate.Source, err.Error())
	}

	err = decorator.CommentQualityGate(qualityGate)
	if err != nil {
		log.Printf("Failure on pull request decoration: %s", err.Error())
	}
}
