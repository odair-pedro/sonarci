package azuredevops

import "testing"

func Test_convertStatus(t *testing.T) {
	type args struct {
		status string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test-ok", args: args{status: "ok"}, want: "SUCCESS"},
		{name: "test-OK", args: args{status: "OK"}, want: "SUCCESS"},
		{name: "test-error", args: args{status: "error"}, want: "FAILED"},
		{name: "test-ERROR", args: args{status: "ERROR"}, want: "FAILED"},
		{name: "test-other", args: args{status: "other"}, want: "OTHER"},
		{name: "test-OTHER", args: args{status: "OTHER"}, want: "OTHER"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStatus(tt.args.status); got != tt.want {
				t.Errorf("convertStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertStatusColor(t *testing.T) {
	type args struct {
		status string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test-ok", args: args{status: "ok"}, want: "green"},
		{name: "test-OK", args: args{status: "OK"}, want: "green"},
		{name: "test-error", args: args{status: "error"}, want: "red"},
		{name: "test-ERROR", args: args{status: "ERROR"}, want: "red"},
		{name: "test-other", args: args{status: "other"}, want: "yellow"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStatusColor(tt.args.status); got != tt.want {
				t.Errorf("convertStatusColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
