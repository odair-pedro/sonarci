package azuredevops

import "sonarci/sonar"

const (
	statusActive = "active"
	statusClosed = "closed"
)

const commentType = "system"

type commentWrapperModel struct {
	Status     string         `json:"status"`
	Properties propertyModel  `json:"properties"`
	Comments   []commentModel `json:"comments"`
}

type propertyModel struct {
	GeneratedBySonarCI bool
}

type commentModel struct {
	Content     string `json:"content"`
	CommentType string `json:"commentType"`
}

func parseCommentModel(qualityGate sonar.QualityGate, report string) commentWrapperModel {
	var status string
	if qualityGate.HasPassed() {
		status = statusClosed
	} else {
		status = statusActive
	}

	return commentWrapperModel{
		Status:     status,
		Properties: propertyModel{GeneratedBySonarCI: true},
		Comments: []commentModel{
			{CommentType: commentType, Content: report},
		}}
}
