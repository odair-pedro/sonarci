package cmd

import "testing"

func Test_validateFlag(t *testing.T) {
	type args struct {
		flagName  string
		flagValue interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "valid_string", args: args{flagName: "name", flagValue: "value"}, want: true},
		{name: "valid_int", args: args{flagName: "name", flagValue: 1}, want: true},
		{name: "invalid_string", args: args{flagName: "name", flagValue: ""}, want: false},
		{name: "invalid_int", args: args{flagName: "name", flagValue: 0}, want: false},
		{name: "invalid_int", args: args{flagName: "name", flagValue: -1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateFlag(tt.args.flagName, tt.args.flagValue); got != tt.want {
				t.Errorf("validateFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
