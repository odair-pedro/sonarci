package factory

import (
	"sonarci/decoration/template/dummy"
	"testing"
)

func Test_CreateDummyTemplateEngine_CheckReturn(t *testing.T) {
	api := CreateDummyTemplateEngine()

	switch tp := api.(type) {
	case *dummy.Engine:
		return
	default:
		t.Errorf("Invalid returned type: %T", tp)
	}
}
