package factory

import (
	"sonarci/net/http"
	"testing"
	"time"
)

func Test_CreateHttpConnection_CheckReturn(t *testing.T) {
	api := CreateHttpConnection("server", "token", time.Second)

	switch tp := api.(type) {
	case *http.Connection:
		return
	default:
		t.Errorf("Invalid returned type: %T", tp)
	}
}
