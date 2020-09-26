package http

import (
	"encoding/base64"
	"fmt"
	"testing"
)

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
