package main

import (
	"errors"
	"os"
	"path/filepath"
)

type Lit struct {
	config *LitConfig
	logger Logger
}

type LitConfig struct {
	DefaultBranchName string
	HeadPath          string
	IgnoreFileName    string
	RootDir           string
	ObjectsDir        string
}

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
}

var (
	RepositoryNotInitialized = errors.New("fatal: not a lit repository (or any of the parent directories)")
	NotValidObjectName       = func(name string) error {
		return errors.New("fatal: Not a valid object name: " + "'" + name + "'.")
	}
)

func NewLit(logger Logger) *Lit {

	return &Lit{
		config: &LitConfig{
			DefaultBranchName: "master",
			HeadPath:          "HEAD",
			IgnoreFileName:    ".litignore",
			RootDir:           ".lit",
			ObjectsDir:        "objects",
		},
		logger: logger,
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
			return filepath.Join(cwd, lit.config.RootDir), nil
		}

		prev := filepath.Dir(cwd)

		if cwd == prev {
			return "", RepositoryNotInitialized
		}

		cwd = prev
	}
}

func (lit *Lit) isInitializedIn(path string) bool {
	_, err := os.Stat(filepath.Join(path, lit.config.RootDir))

	return !os.IsNotExist(err)
}

func (lit *Lit) objectsDir() string {

	litDir, err := lit.LitDir()

	if err != nil {
		lit.logger.Fatal(err)
	}

	return filepath.Join(litDir, lit.config.ObjectsDir)
}
