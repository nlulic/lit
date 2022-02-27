package main

import (
	"os"
	"path/filepath"

	"gitlab.com/nlulic/lit/cad"
)

// TODO: should also be able to pass a commit hash
// TODO: -b param
func (lit *Lit) Checkout(ref string) {

	if ref == "" {
		lit.logger.Fatal("fatal: no ref specified")
	}

	currentRef, err := lit.GetRef()

	if err != nil {
		lit.logger.Fatal(err)
	}

	// check if ref exists in refs/heads
	if _, err := os.Stat(filepath.Join(filepath.Dir(currentRef), ref)); os.IsNotExist(err) {
		lit.logger.Fatal("fatal: ref '" + ref + "' doesn't exist")
	}

	objectsDir := lit.objectsDir()
	db := cad.New(objectsDir)

	head, _ := lit.GetHead()
	rootDir, _ := lit.Root()

	// check if working tree is dirty
	headSnapshot := snapshotHead(head, rootDir, db)
	workingTree := lit.Snapshot(rootDir)

	if headSnapshot.Hash() != workingTree.Hash() {
		lit.logger.Fatal("fatal: changes in current branch would be overwritten")
	}

	// Update HEAD
	lit.logger.Debug("set ref to", ref)
	if err := lit.SetRef(ref); err != nil {
		panic(err)
	}

	// fetch the new HEAD after the ref has been updated
	head, err = lit.GetHead()
	if err != nil {
		panic(err)
	}

	// Snapshot the new ref
	headSnapshot = snapshotHead(head, rootDir, db)
	lit.Sync(headSnapshot, workingTree)

	lit.logger.Info("Switched to branch '" + ref + "'")
}
