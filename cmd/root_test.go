package cmd

import (
	"testing"
)

func Test_NewRootCmd_CheckReturn(t *testing.T) {
	cmd := NewRootCmd()
	if cmd == nil {
		t.Errorf("NewRootCmd() = nil")
	}
}
