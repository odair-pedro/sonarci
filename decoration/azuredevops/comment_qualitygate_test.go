package azuredevops

import "testing"

func Test_formatPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test-1", args: args{path: "path test"}, want: "path%20test"},
		{name: "test-2", args: args{path: "path-test"}, want: "path-test"},
		{name: "test-3", args: args{path: "test/path-test"}, want: "test/path-test"},
		{name: "test-4", args: args{path: "test/path test"}, want: "test/path%20test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatPath(tt.args.path); got != tt.want {
				t.Errorf("formatPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
