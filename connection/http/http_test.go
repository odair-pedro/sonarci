package http

import (
	"encoding/base64"
	"errors"
	"fmt"
	"testing"
	"time"
)

func Test_NewConnection(t *testing.T) {
	if got := NewConnection("host", "token", time.Second); got == nil {
		t.Errorf("NewConnection() return nil")
	}
}

func Test_Connection_GetHostServer(t *testing.T) {
	const hostServer = "host-server"

	connection := &Connection{
		HostServer: hostServer,
	}

	got := connection.GetHostServer()
	if got != hostServer {
		t.Errorf("GetHostServer() = %v, want %v", got, hostServer)
	}
}

func Test_encodeToken(t *testing.T) {
	const token = "123456789"
	want := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", token)))
	if got := encodeToken(token); got != want {
		t.Errorf("getAuthentication() = %v, want %v", got, want)
	}
}

func Test_isStatusSuccess(t *testing.T) {
	type args struct {
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ok", args{statusCode: 200}, true},
		{"created", args{statusCode: 201}, true},
		{"accepted", args{statusCode: 202}, true},
		{"partialInformation", args{statusCode: 203}, true},
		{"noResponse", args{statusCode: 204}, true},
		{"badRequest", args{statusCode: 400}, false},
		{"unauthorized", args{statusCode: 401}, false},
		{"internalServerError", args{statusCode: 500}, false},
		{"badGateway", args{statusCode: 502}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStatusSuccess(tt.args.statusCode); got != tt.want {
				t.Errorf("isStatusSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseUrl(t *testing.T) {
	type args struct {
		host     string
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{host: "http://host", endpoint: "test"}, want: "http://host/test"},
		{name: "test2", args: args{host: "http://host/", endpoint: "test"}, want: "http://host/test"},
		{name: "test3", args: args{host: "http://host", endpoint: "/test"}, want: "http://host/test"},
		{name: "test4", args: args{host: "http://host/", endpoint: "/test"}, want: "http://host/test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseUrl(tt.args.host, tt.args.endpoint); got != tt.want {
				t.Errorf("parseUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockCloser struct {
	closed bool
	close  func() error
}

func (closer *mockCloser) Close() error {
	closer.closed = true
	return closer.close()
}

func Test_closeResource(t *testing.T) {
	closer := &mockCloser{close: func() error {
		return nil
	}}

	closeResource(closer)
	if !closer.closed {
		t.Errorf("closeResource() has not closed resource")
	}
}

func Test_closeResource_CheckPanic(t *testing.T) {
	closer := &mockCloser{close: func() error {
		return errors.New("test-error")
	}}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("None panic has been recovered")
		}
	}()

	closeResource(closer)
}
