package cmd

import (
	"testing"
)

func Test_NewValidateCmd_CheckReturn(t *testing.T) {
	cmd := NewValidateCmd()
	if cmd == nil {
		t.Errorf("NewValidateCmd() = nil")
	}
}
