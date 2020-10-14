package cmd

import (
	"testing"
)

func Test_NewServerVersionCmd_CheckReturn(t *testing.T) {
	cmd := NewServerVersionCmd()
	if cmd == nil {
		t.Errorf("NewServerVersionCmd() = nil")
	}
}
