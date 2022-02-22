package main

import (
	"errors"
	"os"
	"path/filepath"
)

type Lit struct {
	*LitConfig
}

type LitConfig struct {
	DefaultBranchName string
	HeadPath          string
	IgnoreFileName    string
	RootDir           string
}

var (
	RepositoryNotInitialized = errors.New("fatal: not a lit repository (or any of the parent directories)")
)

func NewLit() *Lit {

	return &Lit{
		&LitConfig{
			DefaultBranchName: "master",
			HeadPath:          "HEAD",
			IgnoreFileName:    ".litignore",
			RootDir:           ".lit",
		},
	}
}

// Root finds the closest parent of the lit directory (.lit)
// from the CWD
func (lit *Lit) Root() (string, error) {

	litDir, err := lit.LitDir()

	if err != nil {
		return "", err
	}

	return filepath.Dir(litDir), nil
}

// LitDir finds the closest lit directory (.lit) from the CWD
func (lit *Lit) LitDir() (string, error) {

	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	for {
		if lit.isInitializedIn(cwd) {
			return filepath.Join(cwd, lit.RootDir), nil
		}

		prev := filepath.Dir(cwd)

		if cwd == prev {
			return "", RepositoryNotInitialized
		}

		cwd = prev
	}
}

func (lit *Lit) isInitializedIn(path string) bool {
	_, err := os.Stat(filepath.Join(path, lit.RootDir))

	return !os.IsNotExist(err)
}
