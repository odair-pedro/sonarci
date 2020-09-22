package sonar

type Project struct {
	Id           string
	Organization string
	Key          string
	Name         string
	Visibility   string
}

type Api interface {
	GetServerVersion() (string, error)
	SearchProjects(projects string) (<-chan Project, error)
	ValidateBranch(project string, branch string) error
}
