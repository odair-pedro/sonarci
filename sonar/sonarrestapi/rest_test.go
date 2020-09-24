package sonarrestapi

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"
)

type mockConnection struct {
	doGet func(route string) (<-chan []byte, <-chan error)
}

func (connection *mockConnection) DoGet(route string) (<-chan []byte, <-chan error) {
	return connection.doGet(route)
}

func TestNewApi(t *testing.T) {
	const server = "server"
	const token = "token"
	timeout := time.Duration(1)

	if got := NewApi(server, token, timeout); got == nil {
		t.Errorf("getAuthentication() = nil")
	}
}

func Test_getAuthentication(t *testing.T) {
	const token = "123456789"
	want := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", token)))
	if got := getAuthentication(token); got != want {
		t.Errorf("getAuthentication() = %v, want %v", got, want)
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
