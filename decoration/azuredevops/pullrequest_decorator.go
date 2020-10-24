package azuredevops

import (
	"sonarci/decoration/template"
	"sonarci/net"
)

type PullRequestDecorator struct {
	net.Connection
	template.Engine
}

func NewPullRequestDecorator(connection net.Connection, engine template.Engine) *PullRequestDecorator {
	return &PullRequestDecorator{connection, engine}
}
