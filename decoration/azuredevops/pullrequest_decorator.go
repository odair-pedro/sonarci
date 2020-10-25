package azuredevops

import (
	"sonarci/decoration"
	"sonarci/decoration/template"
)

type PullRequestDecorator struct {
	decoration.Connection
	template.Engine
	Project    string
	Repository string
}

func NewPullRequestDecorator(connection decoration.Connection, engine template.Engine, project string, repository string) *PullRequestDecorator {
	return &PullRequestDecorator{connection, engine, project, repository}
}
