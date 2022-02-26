package azuredevops

import (
	"net/url"
	"sonarci/connection"
	"sonarci/decoration/template"
	"strings"
)

type PullRequestDecorator struct {
	connection.Connection
	template.Engine
	Project    string
	Repository string
}

func NewPullRequestDecorator(connection connection.Connection, engine template.Engine, project string, repository string) *PullRequestDecorator {
	return &PullRequestDecorator{connection, engine, project, repository}
}

func formatPath(path string) string {
	values := strings.Split(path, "/")
	for i, v := range values {
		values[i] = url.PathEscape(v)
	}

	return strings.Join(values, "/")
}
