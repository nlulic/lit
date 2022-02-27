package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (lit *Lit) Branch(branch string) {

	head, err := lit.GetHead()
	currentRef, _ := lit.GetRef()

	if err != nil {
		if os.IsNotExist(err) {
			lit.logger.Fatal(NotValidObjectName(filepath.Base(currentRef)))
		}
		lit.logger.Fatal(err)
	}

	if branch == "" {
		listRefs(currentRef, lit.logger)
		return
	}

	createRef(filepath.Dir(currentRef), branch, head, lit.logger)
}

// listRefs will list all branches (refs)
func listRefs(currentRef string, logger Logger) {
	refs, err := ioutil.ReadDir(filepath.Dir(currentRef))
	if err != nil {
		logger.Fatal(err)
	}

	for _, ref := range refs {
		if filepath.Base(currentRef) == ref.Name() {
			logger.Info("*", ref.Name())
			continue
		}

		logger.Info(ref.Name())
	}
}

// createRef creates a new branch (refs)
func createRef(refsDir, newRef, head string, logger Logger) {

	newRefPath := filepath.Join(refsDir, newRef)

	if _, err := os.Stat(newRefPath); !os.IsNotExist(err) {
		logger.Fatal(fmt.Sprintf("fatal: A branch named '%s' already exists", newRef))
	}

	mustWriteToFile(newRefPath, head)
}
