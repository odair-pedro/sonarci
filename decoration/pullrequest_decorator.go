package decoration

import "sonarci/sonar"

type PullRequestDecorator interface {
	CommentQualityGate(qualityGate sonar.QualityGate) error
}
