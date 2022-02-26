package main

import (
	"os"
	"path/filepath"
)

func (lit *Lit) Init() {

	root, err := filepath.Abs(lit.config.RootDir)

	if err != nil {
		panic(err)
	}

	if lit.isInitializedIn(filepath.Dir(root)) {
		lit.logger.Info("Repository already initialized in", root)
		return
	}

	if err := os.MkdirAll(root, 0644); err != nil {
		panic(err)
	}

	err = lit.SetRef(lit.config.DefaultBranchName)

	if err != nil {
		panic(err)
	}

	lit.logger.Info("Initialized empty Lit repository in", root)
}
