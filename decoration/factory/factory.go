package factory

import (
	"errors"
	"sonarci/decoration"
	"sonarci/decoration/azuredevops"
	"sonarci/decoration/template"
	connFactory "sonarci/net/factory"
	"time"
)

func CreatePullRequestDecorator(decoratorType string, project string, repository string, token string,
	timeout time.Duration, engine template.Engine) (decoration.PullRequestDecorator, error) {

	switch decoratorType {
	case typeAzRepos:
		return createPullRequestAzureDecorator(project, repository, token, timeout, engine), nil
	case typeGitHub:
		return nil, errors.New("GitHub decoration has not yet been implemented =(\nPlease, contribute with the project on https://github.com/odair-pedro/sonarci")
	default:
		return nil, errors.New("Invalid decorator type: " + decoratorType)
	}
}

const (
	typeAzRepos = "azrepos"
	typeGitHub  = "github"
)

func createPullRequestAzureDecorator(project string, repository string, token string, timeout time.Duration,
	engine template.Engine) *azuredevops.PullRequestDecorator {
	conn := connFactory.CreateHttpConnection(azuredevops.Server, token, timeout)
	return azuredevops.NewPullRequestDecorator(conn, engine, project, repository)
}
