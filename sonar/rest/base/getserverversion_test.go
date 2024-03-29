package base

import (
	"errors"
	"sonarci/connection"
	"sonarci/testing/mocks"
	"testing"
)

func Test_restApi_GetServerVersion(t *testing.T) {
	mockVersion := &mocks.MockConnection{GetMock: func(route string) (<-chan []byte, <-chan error) {
		chVersion := make(chan []byte, 1)
		chVersion <- []byte("1.0")
		chErr := make(chan error, 1)
		chErr <- nil
		return chVersion, chErr
	}}
	mockError := &mocks.MockConnection{GetMock: func(route string) (<-chan []byte, <-chan error) {
		chError := make(chan error, 1)
		chError <- errors.New("failure")
		return nil, chError
	}}

	type fields struct {
		Connection connection.Connection
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"success", fields{Connection: mockVersion}, "1.0", false},
		{"error", fields{Connection: mockError}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := NewRestApi(tt.fields.Connection)
			got, err := restApi.GetServerVersion()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServerVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetServerVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}
