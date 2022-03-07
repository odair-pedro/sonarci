package decoration

import "sonarci/sonar"

type PullRequestDecorator interface {
	ClearPreviousComments(pullRequest string, tag string) error
	CommentQualityGate(qualityGate sonar.QualityGate, tag string) error
}
