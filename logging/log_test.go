package logging

import (
	"log"
	"testing"
)

func Test_Setup_CheckFlags(t *testing.T) {
	Setup()
	if log.Flags() != 0 {
		t.Error("Setup() not set flas to 0")
	}
}
