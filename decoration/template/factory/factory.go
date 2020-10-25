package factory

import (
	"sonarci/decoration/template"
	"sonarci/decoration/template/dummy"
)

func CreateDummyTemplateEngine() template.Engine {
	return dummy.NewEngine()
}
