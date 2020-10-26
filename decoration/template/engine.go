package template

type Engine interface {
	ProcessTemplate(template string, dataSource interface{}) (string, error)
}
