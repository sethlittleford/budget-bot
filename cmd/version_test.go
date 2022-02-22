package cmd

import (
	"testing"
)

func TestVersionCmd(t *testing.T) {
	if err := versionCmd.Execute(); err != nil {
		t.Fatalf("Version Command failed with error = %v", err)
	}
}