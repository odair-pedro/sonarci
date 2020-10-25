package base

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

func (connection *mockConnection) Request(endpoint string) (<-chan []byte, <-chan error) {
	return connection.request(endpoint)
}

func Test_NewRestApi(t *testing.T) {
	if got := NewRestApi(&mockConnection{}); got == nil {
		t.Errorf("NewRestApi() return nil")
	}
}

func Test_escapeValue(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "feature/update-pipeline", args: args{value: "feature/update-pipeline-build"}, want: "feature%2Fupdate-pipeline-build"},
		{name: "branch_name@123/test%123#-$$$-#098", args: args{value: "branch_name@123/test%123#-$$$-#098"}, want: "branch_name%40123%2Ftest%25123%23-%24%24%24-%23098"},
		{name: "master", args: args{value: "master"}, want: "master"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapeValue(tt.args.value); got != tt.want {
				t.Errorf("escapeValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
