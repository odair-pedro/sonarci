package cmd

import (
	"testing"
)

func TestSetVersion(t *testing.T) {
	const version = "1.0"
	SetVersion(version)
	if rootCmd.Version != version {
		t.Errorf("Version %s has not been defined", version)
	}
}
