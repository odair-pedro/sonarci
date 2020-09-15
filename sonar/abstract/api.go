package abstract

type Project struct {
	Id           string
	Organization string
	Key          string
	Name         string
	Visibility   string
}

type Api interface {
	SearchProjects(projects string) (<-chan Project, error)
	GetServerVersion() (string, error)
}
