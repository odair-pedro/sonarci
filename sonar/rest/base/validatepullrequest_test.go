package base

import (
	"encoding/json"
	"errors"
	"sonarci/net"
	"testing"
)

func Test_restApi_validatePullRequestStatus_checkError(t *testing.T) {
	type fields struct {
		Connection net.Connection
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
			fields:  fields{&mockConnection{hostServer: "http://server"}},
			args:    args{status: pullRequestStatus{Measures: nil, PullRequest: "pullRequest", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-empty",
			fields:  fields{&mockConnection{hostServer: "http://server"}},
			args:    args{status: pullRequestStatus{Measures: nil, PullRequest: "pullRequest", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-error",
			fields:  fields{&mockConnection{hostServer: "http://server"}},
			args:    args{status: pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "ERROR"}, {Value: "OK"}}, PullRequest: "pullRequest", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-ok",
			fields:  fields{&mockConnection{hostServer: "http://server"}},
			args:    args{status: pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "OK"}, {Value: "ERROR"}}, PullRequest: "pullRequest", Project: "project"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := NewRestApi(tt.fields.Connection)
			if err := restApi.validatePullRequestStatus(tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("validatepullRequestStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_restApi_validatePullRequestStatus_checkErrorMessage(t *testing.T) {
	type fields struct {
		Connection net.Connection
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
			fields:     fields{&mockConnection{hostServer: "http://server"}},
			args:       args{status: pullRequestStatus{Measures: nil, PullRequest: "pullRequest", Project: "project"}},
			wantErrMsg: "Failure on validate quality gate results\nFor more detail, visit: http://server/dashboard?id=project&pullRequest=pullRequest",
		},
		{
			name:       "measures-empty",
			fields:     fields{&mockConnection{hostServer: "http://server"}},
			args:       args{status: pullRequestStatus{Measures: nil, PullRequest: "pullRequest", Project: "project"}},
			wantErrMsg: "Failure on validate quality gate results\nFor more detail, visit: http://server/dashboard?id=project&pullRequest=pullRequest",
		},
		{
			name:       "measures-error",
			fields:     fields{&mockConnection{hostServer: "http://server"}},
			args:       args{status: pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "ERROR"}, {Value: "OK"}}, PullRequest: "pullRequest", Project: "project"}},
			wantErrMsg: "PullRequest pullRequest has not been passed on quality gate\nFor more detail, visit: http://server/dashboard?id=project&pullRequest=pullRequest",
		},
		{
			name:       "measures-ok",
			fields:     fields{&mockConnection{hostServer: "http://server"}},
			args:       args{status: pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "OK"}, {Value: "ERROR"}}, PullRequest: "pullRequest", Project: "project"}},
			wantErrMsg: "anything",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := NewRestApi(tt.fields.Connection)
			if err := restApi.validatePullRequestStatus(tt.args.status); err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("validatepullRequestStatus() error = %v, wantErr %v", err, tt.wantErrMsg)
			}
		})
	}
}

func Test_restApi_ValidatePullRequestInternal(t *testing.T) {
	mockOk := &mockConnection{hostServer: "http://server", doGet: func(route string) (<-chan []byte, <-chan error) {
		bStatus := pullRequestStatusWrapper{Component: pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "OK"}}, PullRequest: "pullRequest", Project: "project"}}
		buff, _ := json.Marshal(bStatus)

		chOk := make(chan []byte, 1)
		chOk <- buff

		chErr := make(chan error, 1)
		chErr <- nil
		return chOk, chErr
	}}
	mockError := &mockConnection{hostServer: "http://server", doGet: func(route string) (<-chan []byte, <-chan error) {
		chError := make(chan error, 1)
		chError <- errors.New("failure")
		return nil, chError
	}}
	mockErrorStatus := &mockConnection{hostServer: "http://server", doGet: func(route string) (<-chan []byte, <-chan error) {
		bStatus := &pullRequestStatus{Measures: []pullRequestStatusMeasure{{Value: "ERROR"}}, PullRequest: "pullRequest", Project: "project"}
		buff, _ := json.Marshal(bStatus)

		chOk := make(chan []byte, 1)
		chOk <- buff

		chErr := make(chan error, 1)
		chErr <- nil
		return chOk, chErr
	}}
	mockInvalidJson := &mockConnection{hostServer: "http://server", doGet: func(route string) (<-chan []byte, <-chan error) {
		chOk := make(chan []byte, 1)
		chOk <- []byte{}

		chEr := make(chan error, 1)
		chEr <- nil
		return chOk, chEr
	}}

	type fields struct {
		Connection net.Connection
	}
	type args struct {
		routeApi    string
		project     string
		pullRequest string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "ok", fields: fields{mockOk}, args: args{"api/route", "project", "pullRequest"}, wantErr: false},
		{name: "error", fields: fields{mockError}, args: args{"api/route", "project", "pullRequest"}, wantErr: true},
		{name: "error-status", fields: fields{mockErrorStatus}, args: args{"api/route", "project", "pullRequest"}, wantErr: true},
		{name: "invalid-json", fields: fields{mockInvalidJson}, args: args{"api/route", "project", "pullRequest"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := NewRestApi(tt.fields.Connection)
			if err := restApi.ValidatePullRequestInternal(tt.args.routeApi, tt.args.project, tt.args.pullRequest); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePullRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
