package sonarrestapi

//func TestApi_ValidateBranch(t *testing.T) {
//}

//func TestApi_validateBranchStatus(t *testing.T) {
//}

//func TestApi_validateBranchStatus_checkResult(t *testing.T) {
//	type fields struct {
//		Connection *http.Connection
//	}
//	type args struct {
//		status *branchStatus
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			api := &Api{
//				Connection: tt.fields.Connection,
//			}
//			if got := api.validateBranchStatus(tt.args.status, nil); got != tt.want {
//				t.Errorf("validateBranchStatus() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestApi_validateBranchStatus_checkLogs(t *testing.T) {
//	type fields struct {
//		Connection *http.Connection
//	}
//	type args struct {
//		status *branchStatus
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   string
//	}{
//		{
//			name:   "measures-nil",
//			fields: fields{Connection: &http.Connection{Server: "http://server"}},
//			args:   args{status: &branchStatus{Measures: nil, Branch: "branch", Project: "project"}},
//			want:   "Failure on validate quality gate results\nFor more detail, visit: http://server/dashboard?branch=branch&id=project",
//		},
//		{
//			name:   "measures-empty",
//			fields: fields{Connection: &http.Connection{Server: "http://server"}},
//			args:   args{status: &branchStatus{Measures: []branchStatusMeasure{}, Branch: "branch", Project: "project"}},
//			want:   "Failure on validate quality gate results\nFor more detail, visit: http://server/dashboard?branch=branch&id=project\n",
//		},
//		{
//			name:   "measures-error",
//			fields: fields{Connection: &http.Connection{Server: "http://server"}},
//			args:   args{status: &branchStatus{Measures: []branchStatusMeasure{{Value: "ERROR"}, {Value: "OK"}}, Branch: "branch", Project: "project"}},
//			want:   "",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			con := &mockConnection{Connection: tt.fields.Connection}
//
//			api := &Api{
//				Connection: con
//			}
//
//			writer := &mockWriter{}
//			api.validateBranchStatus(tt.args.status, writer)
//			if writer.wroteText != tt.want {
//				t.Errorf("validateBranchStatus generated log: %s, want %s", writer.wroteText, tt.want)
//			}
//		})
//	}
//}
//
//type mockConnection struct {
//	*http.Connection
//}
//
//func (connection *mockConnection) DoGet(route string) (<-chan []byte, <-chan error) {
//	return nil, nil
//}
//
//type mockWriter struct {
//	wroteText string
//}
//
//func (w *mockWriter) Write(p []byte) (n int, err error) {
//	w.wroteText = string(p)
//	return len(p), nil
//}
