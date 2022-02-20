package main

import (
	"log"
	"path/filepath"
)

func (lit *Lit) Init() {

	root, err := filepath.Abs(lit.RootDir)

	if err != nil {
		log.Fatal(err)
	}

	if lit.isInitializedIn(filepath.Dir(root)) {
		log.Printf("Repository already initialized in %s\n", root)
		return
	}

	mustMakeDirs(root)

	err = lit.SetRef(lit.DefaultBranchName)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Initialized empty Lit repository in %s\n", root)
}
