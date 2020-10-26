package azuredevops

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sonarci/decoration/template"
	"sonarci/sonar"
	"strings"
)

const Server = "https://dev.azure.com"
const routeCommentPullRequest = "%s/_apis/git/repositories/%s/pullRequests/%s/threads?api-version=6.0"

func (decorator *PullRequestDecorator) CommentQualityGate(qualityGate sonar.QualityGate) error {
	tplModel := parseTemplateModel(qualityGate)
	report, err := decorator.Engine.ProcessTemplate(template.ReportTemplate, tplModel)
	if err != nil {
		return err
	}

	commentModel := parseCommentModel(qualityGate, report)
	body, _ := json.Marshal(commentModel)

	endpoint := fmt.Sprintf(routeCommentPullRequest, formatPath(decorator.Project), formatPath(decorator.Repository), qualityGate.Source)
	_, chErr := decorator.Connection.Send(endpoint, body, "application/json")
	return <-chErr
}

func formatPath(path string) string {
	values := strings.Split(path, "/")
	for i, v := range values {
		values[i] = url.PathEscape(v)
	}

	return strings.Join(values, "/")
}
