package v8

import (
	"testing"
	"time"
)

func Test_NewRestApi(t *testing.T) {
	const server = "server"
	const token = "token"
	timeout := time.Duration(1)

	if got := NewRestApi(server, token, timeout); got == nil {
		t.Errorf("NewRestApi() return nil")
	}
}
