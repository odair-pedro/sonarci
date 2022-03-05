package azuredevops

import (
	"encoding/json"
	"fmt"
	"sonarci/decoration/azuredevops/models"
	"sonarci/sonar"
)

const routeCommentPullRequest = "%s/_apis/git/repositories/%s/pullRequests/%s/threads?api-version=6.0"

func (decorator *PullRequestDecorator) CommentQualityGate(qualityGate sonar.QualityGate, tag string) error {
	model := models.ParseTemplateModel(qualityGate, tag)
	template := decorator.Engine.GetQualityReportTemplate(tag != "")

	report, err := decorator.Engine.ProcessTemplate(template, model)
	if err != nil {
		return err
	}

	commentModel := models.ParseCommentModel(qualityGate, report)
	body, _ := json.Marshal(commentModel)

	endpoint := fmt.Sprintf(routeCommentPullRequest, formatPath(decorator.Project), formatPath(decorator.Repository), qualityGate.Source)
	_, chErr := decorator.Connection.Post(endpoint, body, "application/json")
	return <-chErr
}
