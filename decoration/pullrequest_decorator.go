package decoration

import "sonarci/sonar"

type PullRequestDecorator interface {
	ClearPreviousComments(pullRequest string) error
	CommentQualityGate(qualityGate sonar.QualityGate) error
}
