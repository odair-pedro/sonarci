package factory

import (
	"sonarci/decoration/template"
	"sonarci/decoration/template/dummy"
)

func CreateTemplateEngine() template.Engine {
	return dummy.NewEngine()
}
