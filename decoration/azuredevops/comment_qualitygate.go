package azuredevops

import (
	"encoding/json"
	"fmt"
	"sonarci/decoration/azuredevops/models"
	"sonarci/decoration/template"
	"sonarci/sonar"
)

const routeCommentPullRequest = "%s/_apis/git/repositories/%s/pullRequests/%s/threads?api-version=6.0"

func (decorator *PullRequestDecorator) CommentQualityGate(qualityGate sonar.QualityGate) error {
	tplModel := models.ParseTemplateModel(qualityGate)
	report, err := decorator.Engine.ProcessTemplate(template.ReportTemplate, tplModel)
	if err != nil {
		return err
	}

	commentModel := models.ParseCommentModel(qualityGate, report)
	body, _ := json.Marshal(commentModel)

	endpoint := fmt.Sprintf(routeCommentPullRequest, formatPath(decorator.Project), formatPath(decorator.Repository), qualityGate.Source)
	_, chErr := decorator.Connection.Post(endpoint, body, "application/json")
	return <-chErr
}
