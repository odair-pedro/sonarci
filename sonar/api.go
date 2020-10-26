package sonar

type Api interface {
	GetServerVersion() (string, error)
	GetBranchQualityGate(project string, branch string) (QualityGate, error)
	GetPullRequestQualityGate(project string, pullRequest string) (QualityGate, error)
	SearchProjects(projects string) (<-chan Project, error)
	ValidateBranch(project string, branch string) error
	ValidatePullRequest(project string, pullRequest string) error
}
