package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"sonarci/connection/http"
	"sonarci/decoration"
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
	decorator, err := decorationFactory.CreatePullRequestDecorator(decoratorType, project, repository, engine,
		func(server string) decoration.Connection {
			return http.NewConnection(server, token, timeout)
		})
	if err != nil {
		log.Printf(err.Error())
		return
	}

	err = decorator.CommentQualityGate(qualityGate)
	if err != nil {
		log.Print("Failure on pull request decoration: ", err.Error())
	}
}
