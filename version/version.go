package version

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	// set at build time
	commitHash string
	branch     string

	// set at run time
	os           string
	architecture string
)

func init() {
	os = runtime.GOOS
	architecture = runtime.GOARCH
}

type version struct {
	commitHash   string
	branch       string
	os           string
	architecture string
}

func Version() (string, error) {
	v := version{
		commitHash:   commitHash,
		branch:       branch,
		os:           os,
		architecture: architecture,
	}
	if v.validVersion() {
		return v.String(), nil
	}
	return "", errors.New("current version is not valid")
}

func (v version) validVersion() bool {
	if v.commitHash == "" {
		return false
	}
	if v.branch == "" {
		return false
	}
	if v.os == "" {
		return false
	}
	if v.architecture == "" {
		return false
	}
	return true
}

func (v version) String() string {
	return fmt.Sprintf("budget-bot_%s_%s\nCommitHash: %s\nBranch: %s\n", v.os, v.architecture, v.commitHash, v.branch)
}
