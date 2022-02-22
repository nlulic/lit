package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const REF_PREFIX = "ref:"

// SetRef writes the head reference in the HEAD file
// `ref` can be a branch name or a commit hash
func (lit *Lit) SetRef(ref string) error {

	rootDir, err := lit.LitDir()

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(
		filepath.Join(rootDir, lit.HeadPath),
		[]byte(fmt.Sprintf("%s refs/heads/%s\n", REF_PREFIX, ref)),
		0644,
	)

	return err
}

// GetRef reads the current HEAD reference
// from the HEAD file
func (lit *Lit) GetRef() (string, error) {

	rootDir, err := lit.LitDir()

	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadFile(
		filepath.Join(rootDir, lit.HeadPath),
	)

	if err != nil {
		return "", err
	}

	ref := string(b)[len(REF_PREFIX)+1:]

	return strings.TrimSpace(ref), nil
}
