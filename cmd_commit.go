package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gitlab.com/nlulic/lit/cad"
	. "gitlab.com/nlulic/lit/objects"
	"gitlab.com/nlulic/lit/util"
)

func (lit *Lit) Commit(message string) {

	r, err := lit.Root()

	if err != nil {
		lit.logger.Fatal(err)
	}

	snapshot := lit.Snapshot(r)

	objectsDir := lit.objectsDir()
	db := cad.New(objectsDir)

	// used later for STDOUT
	var createdBlobs []Blob

	// write all the trees and files to the objects storage
	for _, tree := range trees(snapshot) {
		_, exists := mustCreate(db, []byte(tree.Value()), TreeType, tree.Hash())
		if !exists {
			lit.logger.Debug(fmt.Sprintf("created %s object %s", TreeType, tree.Hash()))
		}

		for _, blob := range tree.Blobs {
			b, err := ioutil.ReadFile(blob.Path)
			if err != nil {
				panic(err)
			}

			_, exists := mustCreate(db, b, BlobType, blob.Hash)
			if !exists {
				createdBlobs = append(createdBlobs, blob)
				lit.logger.Debug(fmt.Sprintf("created %s object %s", BlobType, blob.Hash))
			}
		}
	}

	ref, err := lit.GetRef()
	if err != nil {
		panic(err)
	}

	// exit if no objects have been added to the cad
	if len(createdBlobs) == 0 {
		lit.logger.Info("On branch", filepath.Base(ref))
		lit.logger.Info("nothing to commit, working tree clean")
		return
	}

	head, err := lit.GetHead()

	if err != nil {
		if os.IsNotExist(err) {
			lit.logger.Debug("HEAD currently doesn't exist")
		} else {
			panic(err)
		}
	}

	commit := NewCommit(message, &snapshot, head)
	hash, _ := mustCreate(db, []byte(commit.Value()), CommitType, commit.Hash())
	lit.logger.Debug(fmt.Sprintf("created %s object %s", CommitType, commit.Hash()))

	mustWriteToFile(ref, hash)

	// output
	lit.logger.Info(fmt.Sprintf("[%s (commit %s)] %s", filepath.Base(ref), hash, message))
	lit.logger.Info(len(createdBlobs), "files changed:")
	for _, blob := range createdBlobs {
		lit.logger.Info("created/updated", blob.Name)
	}
}

// tree traverses a base tree and returns the base and all of its subtrees
func trees(tree Tree) (trees []Tree) {

	if tree.IsEmpty() {
		return nil
	}

	stack := util.Stack()
	stack.Push(tree)

	for !stack.IsEmpty() {
		next := stack.Pop().(Tree)
		trees = append(trees, next)
		for _, t := range next.Trees {
			stack.Push(t)
		}
	}

	return
}

// mustCreate creates object to the cad. If any errors occur or the snpashot hash
// doesn't match the created hashed object it panics
func mustCreate(db *cad.Cad, b []byte, objectType string, snapshotHash string) (hash string, exists bool) {

	hash, err := db.Write(b, objectType)

	if err != nil {
		// ignore if the object already exists in the cad
		if err == cad.ObjectAlreadyExists {
			return hash, true
		}

		panic(err)
	}

	if hash != snapshotHash {
		panic(fmt.Sprintf("created %s hash %s doesn't match snapshot hash %s", objectType, hash, snapshotHash))
	}

	return hash, false
}

func mustWriteToFile(path, value string) {

	if err := os.MkdirAll(filepath.Dir(path), 0664); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(path, []byte(value), 0664); err != nil {
		panic(err)
	}
}
