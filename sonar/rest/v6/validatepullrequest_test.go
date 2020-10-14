package v6

import (
	"encoding/json"
	"errors"
	"sonarci/net"
	"testing"
)

func Test_restApi_validatePullRequestStatus_checkError(t *testing.T) {
	type fields struct {
		Connection net.Connection
		Server     string
	}
	type args struct {
		status pullRequestStatus
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "measures-nil",
			fields:  fields{&mockConnection{}, "http://server"},
			args:    args{status: pullRequestStatus{Measures: nil, PullRequest: "pullRequest", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-empty",
			fields:  fields{&mockConnection{}, "http://server"},
			args:    args{status: pullRequestStatus{Measures: nil, PullRequest: "pullRequest", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-error",
			fields:  fields{&mockConnection{}, "http://server"},
			args:    args{status: pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "ERROR"}, {Value: "OK"}}, PullRequest: "pullRequest", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-ok",
			fields:  fields{&mockConnection{}, "http://server"},
			args:    args{status: pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "OK"}, {Value: "ERROR"}}, PullRequest: "pullRequest", Project: "project"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := &RestApi{
				Connection: tt.fields.Connection,
				Server:     tt.fields.Server,
			}
			if err := restApi.validatePullRequestStatus(tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("validatepullRequestStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_restApi_validatepullRequestStatus_checkErrorMessage(t *testing.T) {
	type fields struct {
		Connection net.Connection
		Server     string
	}
	type args struct {
		status pullRequestStatus
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErrMsg string
	}{
		{
			name:       "measures-nil",
			fields:     fields{&mockConnection{}, "http://server"},
			args:       args{status: pullRequestStatus{Measures: nil, PullRequest: "pullRequest", Project: "project"}},
			wantErrMsg: "Failure on validate quality gate results\nFor more detail, visit: http://server/dashboard?id=project&pullRequest=pullRequest",
		},
		{
			name:       "measures-empty",
			fields:     fields{&mockConnection{}, "http://server"},
			args:       args{status: pullRequestStatus{Measures: nil, PullRequest: "pullRequest", Project: "project"}},
			wantErrMsg: "Failure on validate quality gate results\nFor more detail, visit: http://server/dashboard?id=project&pullRequest=pullRequest",
		},
		{
			name:       "measures-error",
			fields:     fields{&mockConnection{}, "http://server"},
			args:       args{status: pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "ERROR"}, {Value: "OK"}}, PullRequest: "pullRequest", Project: "project"}},
			wantErrMsg: "PullRequest pullRequest has not been passed on quality gate\nFor more detail, visit: http://server/dashboard?id=project&pullRequest=pullRequest",
		},
		{
			name:       "measures-ok",
			fields:     fields{&mockConnection{}, "http://server"},
			args:       args{status: pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "OK"}, {Value: "ERROR"}}, PullRequest: "pullRequest", Project: "project"}},
			wantErrMsg: "anything",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := &RestApi{
				Connection: tt.fields.Connection,
				Server:     tt.fields.Server,
			}
			if err := restApi.validatePullRequestStatus(tt.args.status); err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("validatepullRequestStatus() error = %v, wantErr %v", err, tt.wantErrMsg)
			}
		})
	}
}

func Test_restApi_ValidatePullRequest(t *testing.T) {
	mockOk := &mockConnection{doGet: func(route string) (<-chan []byte, <-chan error) {
		bStatus := pullRequestStatusWrapper{Component: pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "OK"}}, PullRequest: "pullRequest", Project: "project"}}
		buff, _ := json.Marshal(bStatus)

		chOk := make(chan []byte, 1)
		chOk <- buff

		chErr := make(chan error, 1)
		chErr <- nil
		return chOk, chErr
	}}
	mockError := &mockConnection{doGet: func(route string) (<-chan []byte, <-chan error) {
		chError := make(chan error, 1)
		chError <- errors.New("failure")
		return nil, chError
	}}
	mockErrorStatus := &mockConnection{doGet: func(route string) (<-chan []byte, <-chan error) {
		bStatus := &pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "ERROR"}}, PullRequest: "pullRequest", Project: "project"}
		buff, _ := json.Marshal(bStatus)

		chOk := make(chan []byte, 1)
		chOk <- buff

		chErr := make(chan error, 1)
		chErr <- nil
		return chOk, chErr
	}}
	mockInvalidJson := &mockConnection{doGet: func(route string) (<-chan []byte, <-chan error) {
		chOk := make(chan []byte, 1)
		chOk <- []byte{}

		chEr := make(chan error, 1)
		chEr <- nil
		return chOk, chEr
	}}

	type fields struct {
		Connection net.Connection
		Server     string
	}
	type args struct {
		project     string
		pullRequest string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "ok", fields: fields{mockOk, "http://server"}, args: args{"project", "pullRequest"}, wantErr: false},
		{name: "error", fields: fields{mockError, "http://server"}, args: args{"project", "pullRequest"}, wantErr: true},
		{name: "error-status", fields: fields{mockErrorStatus, "http://server"}, args: args{"project", "pullRequest"}, wantErr: true},
		{name: "invalid-json", fields: fields{mockInvalidJson, "http://server"}, args: args{"project", "pullRequest"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := &RestApi{
				Connection: tt.fields.Connection,
				Server:     tt.fields.Server,
			}
			if err := restApi.ValidatePullRequest(tt.args.project, tt.args.pullRequest); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePullRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRestApi_GetRouteForValidatePullRequest(t *testing.T) {
	const want = "/api/measures/component?componentKey=%s&pullRequest=%s&metricKeys=alert_status"

	restApi := &RestApi{}
	got := restApi.GetRouteForValidatePullRequest()

	if got != want {
		t.Errorf("GetRouteForValidatePullRequest() got %s, want %s", got, want)
	}
}
