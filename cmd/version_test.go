package cmd

import (
	"testing"
)

// TestVersionCmd must be run with go test -ldflags (see Makefile for details)
func TestVersionCmd(t *testing.T) {
	rootCmd.SetArgs([]string{"version"})
	if err := versionCmd.Execute(); err != nil {
		t.Fatalf("Version Command failed with error = %v", err)
	}
}