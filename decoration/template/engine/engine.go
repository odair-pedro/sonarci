package engine

import "sonarci/decoration/template"

type Engine interface {
	template.Resource
	ProcessTemplate(template string, dataSource interface{}) (string, error)
}
