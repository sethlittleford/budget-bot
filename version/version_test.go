package version

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

// TestVersionSuccess must be run with go test -ldflags (see Makefile for details)
func TestVersionSuccess(t *testing.T) {
	// get the current commit hash
	hashOut, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		t.Error(err)
	}
	commitHash := removeNewline(string(hashOut))

	// get the current branch name
	branchOut, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		t.Error(err)
	}
	branch := removeNewline(string(branchOut))

	got, err := Version()
	if err != nil {
		t.Fatalf("Version() test failed with error = %v", err)
	}
	want := fmt.Sprintf("budget-bot_%s_%s\nCommitHash: %s\nBranch: %s\n", runtime.GOOS, runtime.GOARCH, commitHash, branch)
	if got != want {
		t.Fatalf("Version() = %v, want %v", got, want)
	}
}

func removeNewline(s string) string {
	return strings.TrimSuffix(s, "\n")
}
