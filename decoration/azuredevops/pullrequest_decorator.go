package azuredevops

import (
	"net/url"
	"sonarci/connection"
	"sonarci/decoration/template/engine"
	"strings"
)

type PullRequestDecorator struct {
	connection.Connection
	engine.Engine
	Project    string
	Repository string
}

func NewPullRequestDecorator(connection connection.Connection, engine engine.Engine, project string, repository string) *PullRequestDecorator {
	return &PullRequestDecorator{connection, engine, project, repository}
}

func formatPath(path string) string {
	values := strings.Split(path, "/")
	for i, v := range values {
		values[i] = url.PathEscape(v)
	}

	return strings.Join(values, "/")
}
