package models

import "sonarci/sonar"

const (
	statusActive = "active"
	statusClosed = "closed"
)

const commentType = "system"

type CommentWrapperModel struct {
	Status     string               `json:"status"`
	Properties CommentPropertyModel `json:"properties"`
	Comments   []CommentModel       `json:"comments"`
}

type CommentPropertyModel struct {
	GeneratedBySonarCI bool
}

type CommentModel struct {
	Content     string `json:"content"`
	CommentType string `json:"commentType"`
}

func ParseCommentModel(qualityGate sonar.QualityGate, report string) CommentWrapperModel {
	var status string
	if qualityGate.HasPassed() {
		status = statusClosed
	} else {
		status = statusActive
	}

	return CommentWrapperModel{
		Status:     status,
		Properties: CommentPropertyModel{GeneratedBySonarCI: true},
		Comments: []CommentModel{
			{CommentType: commentType, Content: report},
		}}
}
