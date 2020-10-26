package v8

import (
	"testing"
)

type mockConnection struct {
	hostServer string
	request    func(route string) (<-chan []byte, <-chan error)
}

func (connection *mockConnection) GetHostServer() string {
	return connection.hostServer
}

func (connection *mockConnection) Request(route string) (<-chan []byte, <-chan error) {
	return connection.request(route)
}

func Test_NewRestApi(t *testing.T) {
	if got := NewRestApi(&mockConnection{}); got == nil {
		t.Errorf("NewRestApi() return nil")
	}
}
