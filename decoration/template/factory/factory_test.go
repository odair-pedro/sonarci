package factory

import (
	"sonarci/decoration/template/dummy"
	"testing"
)

func Test_CreateTemplateEngine_CheckReturn(t *testing.T) {
	api := CreateTemplateEngine()

	switch tp := api.(type) {
	case *dummy.Engine:
		return
	default:
		t.Errorf("Invalid returned type: %T", tp)
	}
}
