package azuredevops

import (
	"sonarci/decoration/template"
	"sonarci/net"
)

type PullRequestDecorator struct {
	net.Connection
	template.Engine
	Project    string
	Repository string
}

func NewPullRequestDecorator(connection net.Connection, engine template.Engine, project string, repository string) *PullRequestDecorator {
	return &PullRequestDecorator{connection, engine, project, repository}
}
