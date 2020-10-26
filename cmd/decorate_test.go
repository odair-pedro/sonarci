package cmd

import "testing"

func Test_NewDecorateCmd_CheckReturn(t *testing.T) {
	cmd := NewDecorateCmd()
	if cmd == nil {
		t.Errorf("NewDecorateCmd() = nil")
	}
}
