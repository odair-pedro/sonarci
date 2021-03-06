package factory

import (
	"errors"
	"sonarci/decoration"
	"sonarci/decoration/azuredevops"
	"sonarci/decoration/template"
)

func CreatePullRequestDecorator(decoratorType string, project string, repository string,
	engine template.Engine, connectionFactory func(server string) decoration.Connection) (decoration.PullRequestDecorator, error) {

	switch decoratorType {
	case typeAzRepos:
		return createPullRequestAzureDecorator(project, repository, engine, connectionFactory), nil
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

func createPullRequestAzureDecorator(project string, repository string, engine template.Engine,
	connectionFactory func(server string) decoration.Connection) *azuredevops.PullRequestDecorator {
	return azuredevops.NewPullRequestDecorator(connectionFactory(azuredevops.Server), engine, project, repository)
}
