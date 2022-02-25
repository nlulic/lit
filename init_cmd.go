package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (lit *Lit) Init() {

	root, err := filepath.Abs(lit.config.RootDir)

	if err != nil {
		panic(err)
	}

	if lit.isInitializedIn(filepath.Dir(root)) {
		fmt.Printf("Repository already initialized in %s\n", root)
		return
	}

	if err := os.MkdirAll(root, 0644); err != nil {
		panic(err)
	}

	err = lit.SetRef(lit.config.DefaultBranchName)

	if err != nil {
		panic(err)
	}

	log.Printf("Initialized empty Lit repository in %s\n", root)
}
