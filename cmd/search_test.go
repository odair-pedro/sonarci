package cmd

import "testing"

func Test_padRight(t *testing.T) {
	type args struct {
		str    string
		suffix string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test-0", args{"test", "0", 10}, "test000000"},
		{"test-01", args{"test", "01", 10}, "test010101"},
		{"test-010", args{"test", "010", 10}, "test010010"},
		{"test-01010101", args{"test", "01010101", 10}, "test010101"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := padRight(tt.args.str, tt.args.suffix, tt.args.length); got != tt.want {
				t.Errorf("padRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewSearchCmd_CheckReturn(t *testing.T) {
	cmd := NewSearchCmd()
	if cmd == nil {
		t.Errorf("NewSearchCmd() = nil")
	}
}
