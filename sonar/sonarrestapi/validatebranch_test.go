package sonarrestapi

import (
	"sonarci/net"
	"testing"
)

func Test_restApi_validateBranchStatus_checkError(t *testing.T) {
	type fields struct {
		Connection net.Connection
		Server     string
	}
	type args struct {
		status *branchStatus
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
			args:    args{status: &branchStatus{Measures: nil, Branch: "branch", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-empty",
			fields:  fields{&mockConnection{}, "http://server"},
			args:    args{status: &branchStatus{Measures: nil, Branch: "branch", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-error",
			fields:  fields{&mockConnection{}, "http://server"},
			args:    args{status: &branchStatus{Measures: []branchStatusMeasure{{Value: "ERROR"}, {Value: "OK"}}, Branch: "branch", Project: "project"}},
			wantErr: true,
		},
		{
			name:    "measures-ok",
			fields:  fields{&mockConnection{}, "http://server"},
			args:    args{status: &branchStatus{Measures: []branchStatusMeasure{{Value: "OK"}, {Value: "ERROR"}}, Branch: "branch", Project: "project"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := &restApi{
				Connection: tt.fields.Connection,
				Server:     tt.fields.Server,
			}
			if err := restApi.validateBranchStatus(tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("validateBranchStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_restApi_validateBranchStatus_checkErrorMessage(t *testing.T) {
	type fields struct {
		Connection net.Connection
		Server     string
	}
	type args struct {
		status *branchStatus
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
			args:       args{status: &branchStatus{Measures: nil, Branch: "branch", Project: "project"}},
			wantErrMsg: "Failure on validate quality gate results\nFor more detail, visit: http://server/dashboard?branch=branch&id=project",
		},
		{
			name:       "measures-empty",
			fields:     fields{&mockConnection{}, "http://server"},
			args:       args{status: &branchStatus{Measures: nil, Branch: "branch", Project: "project"}},
			wantErrMsg: "Failure on validate quality gate results\nFor more detail, visit: http://server/dashboard?branch=branch&id=project",
		},
		{
			name:       "measures-error",
			fields:     fields{&mockConnection{}, "http://server"},
			args:       args{status: &branchStatus{Measures: []branchStatusMeasure{{Value: "ERROR"}, {Value: "OK"}}, Branch: "branch-name", Project: "project"}},
			wantErrMsg: "Branch branch-name has not been passed on quality gate\nFor more detail, visit: http://server/dashboard?branch=branch-name&id=project",
		},
		{
			name:       "measures-ok",
			fields:     fields{&mockConnection{}, "http://server"},
			args:       args{status: &branchStatus{Measures: []branchStatusMeasure{{Value: "OK"}, {Value: "ERROR"}}, Branch: "branch", Project: "project"}},
			wantErrMsg: "anything",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := &restApi{
				Connection: tt.fields.Connection,
				Server:     tt.fields.Server,
			}
			if err := restApi.validateBranchStatus(tt.args.status); err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("validateBranchStatus() error = %v, wantErr %v", err, tt.wantErrMsg)
			}
		})
	}
}

type mockConnection struct {
}

func (connection *mockConnection) DoGet(route string) (<-chan []byte, <-chan error) {
	_ = route
	return nil, nil
}
