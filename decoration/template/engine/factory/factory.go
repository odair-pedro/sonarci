package factory

import (
	"sonarci/decoration/template/engine"
	"sonarci/decoration/template/engine/dummy"
)

func CreateDummyTemplateEngine() engine.Engine {
	return dummy.NewEngine()
}
