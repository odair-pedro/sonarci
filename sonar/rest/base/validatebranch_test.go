package base

import (
	"encoding/json"
	"errors"
	"sonarci/net"
	"testing"
)

func Test_restApi_validateBranchStatus_checkError(t *testing.T) {
	type fields struct {
		Connection net.Connection
	}
	type args struct {
		status branchStatus
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
			args:    args{status: branchStatus{Measures: nil, Branch: "branch", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-empty",
			fields:  fields{&mockConnection{hostServer: "http://server"}},
			args:    args{status: branchStatus{Measures: nil, Branch: "branch", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-error",
			fields:  fields{&mockConnection{hostServer: "http://server"}},
			args:    args{status: branchStatus{Measures: []branchStatusMeasure{{Value: "ERROR"}, {Value: "OK"}}, Branch: "branch", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-ok",
			fields:  fields{&mockConnection{hostServer: "http://server"}},
			args:    args{status: branchStatus{Measures: []branchStatusMeasure{{Value: "OK"}, {Value: "ERROR"}}, Branch: "branch", Project: "project"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := NewRestApi(tt.fields.Connection)
			if err := restApi.validateBranchStatus(tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("validateBranchStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_restApi_validateBranchStatus_checkErrorMessage(t *testing.T) {
	type fields struct {
		Connection net.Connection
	}
	type args struct {
		status branchStatus
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
			args:       args{status: branchStatus{Measures: nil, Branch: "branch", Project: "project"}},
			wantErrMsg: "Failure on validate quality gate results\nFor more detail, visit: http://server/dashboard?id=project&branch=branch",
		},
		{
			name:       "measures-empty",
			fields:     fields{&mockConnection{hostServer: "http://server"}},
			args:       args{status: branchStatus{Measures: nil, Branch: "branch", Project: "project"}},
			wantErrMsg: "Failure on validate quality gate results\nFor more detail, visit: http://server/dashboard?id=project&branch=branch",
		},
		{
			name:       "measures-error",
			fields:     fields{&mockConnection{hostServer: "http://server"}},
			args:       args{status: branchStatus{Measures: []branchStatusMeasure{{Value: "ERROR"}, {Value: "OK"}}, Branch: "branch-name", Project: "project"}},
			wantErrMsg: "Branch branch-name has not been passed on quality gate\nFor more detail, visit: http://server/dashboard?id=project&branch=branch-name",
		},
		{
			name:       "measures-ok",
			fields:     fields{&mockConnection{hostServer: "http://server"}},
			args:       args{status: branchStatus{Measures: []branchStatusMeasure{{Value: "OK"}, {Value: "ERROR"}}, Branch: "branch", Project: "project"}},
			wantErrMsg: "anything",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := NewRestApi(tt.fields.Connection)
			if err := restApi.validateBranchStatus(tt.args.status); err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("validateBranchStatus() error = %v, wantErr %v", err, tt.wantErrMsg)
			}
		})
	}
}

func Test_restApi_ValidateBranchInternal(t *testing.T) {
	mockOk := &mockConnection{hostServer: "http://server", doGet: func(route string) (<-chan []byte, <-chan error) {
		bStatus := branchStatusWrapper{Component: branchStatus{Measures: []branchStatusMeasure{{Value: "OK"}}, Branch: "branch-name", Project: "project"}}
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
		bStatus := &branchStatus{Measures: []branchStatusMeasure{{Value: "ERROR"}}, Branch: "branch-name", Project: "project"}
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
	}
	type args struct {
		routeApi string
		project  string
		branch   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "ok", fields: fields{mockOk}, args: args{"api/route", "project", "branch-name"}, wantErr: false},
		{name: "error", fields: fields{mockError}, args: args{"api/route", "project", "branch-name"}, wantErr: true},
		{name: "error-status", fields: fields{mockErrorStatus}, args: args{"api/route", "project", "branch-name"}, wantErr: true},
		{name: "invalid-json", fields: fields{mockInvalidJson}, args: args{"api/route", "project", "branch-name"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := NewRestApi(tt.fields.Connection)
			if err := restApi.ValidateBranchInternal(tt.args.routeApi, tt.args.project, tt.args.branch); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBranch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
