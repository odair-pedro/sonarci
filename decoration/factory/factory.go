package factory

import (
	"sonarci/decoration"
	"sonarci/decoration/azuredevops"
	"sonarci/decoration/template"
	"sonarci/net"
)

func CreatePullRequestDecorator(connection net.Connection, engine template.Engine) decoration.PullRequestDecorator {
	return azuredevops.NewPullRequestDecorator(connection, engine)
}
