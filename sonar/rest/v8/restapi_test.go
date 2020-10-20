package v8

import (
	"testing"
)

type mockConnection struct {
	hostServer string
	doGet      func(route string) (<-chan []byte, <-chan error)
}

func (connection *mockConnection) GetHostServer() string {
	return connection.hostServer
}

func (connection *mockConnection) DoGet(route string) (<-chan []byte, <-chan error) {
	return connection.doGet(route)
}

func Test_NewRestApi(t *testing.T) {
	if got := NewRestApi(&mockConnection{}); got == nil {
		t.Errorf("NewRestApi() return nil")
	}
}
