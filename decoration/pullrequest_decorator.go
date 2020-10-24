package decoration

import "sonarci/sonar"

type PullRequestDecorator interface {
	CommentSummary(pullRequest string, qualityGate sonar.QualityGate) error
}
