package azuredevops

import "sonarci/sonar"

func (decorator *PullRequestDecorator) CommentSummary(pullRequest string, qualityGate sonar.QualityGate) error {
	return nil
}
