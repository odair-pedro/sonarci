package factory

import (
	v8 "sonarci/sonar/rest/v8"
	"testing"
)

func Test_CreateSonarApi_CheckReturn(t *testing.T) {
	api := CreateSonarApi(nil)

	switch tp := api.(type) {
	case *v8.RestApi:
		return
	default:
		t.Errorf("Invalid returned type: %T", tp)
	}
}
