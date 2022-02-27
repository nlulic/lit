package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gitlab.com/nlulic/lit/cad"
	. "gitlab.com/nlulic/lit/objects"
)

// snapshot creates a Tree of the passed directory
func (lit *Lit) Snapshot(dir string) Tree {
	ignorefiles, err := lit.Ignorefiles()

	if err != nil {
		lit.logger.Fatal(err)
	}

	lit.logger.Debug("creating directory tree snapshot for", dir)
	defer lit.logger.Debug("created recursive directory snapshot for", dir)

	return snapshotDir(dir, ignorefiles)
}

// snapshotDir creates a Tree from a directory and calls itself recursively
func snapshotDir(dir string, ignorefiles []string) Tree {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	tree := Tree{
		Name: filepath.Base(dir),
		Path: dir,
	}

	for _, file := range files {
		path := filepath.Join(dir, file.Name())

		if contains(ignorefiles, path) {
			continue
		}

		if file.IsDir() {
			t := snapshotDir(path, ignorefiles) // recursive call

			if !t.IsEmpty() {
				tree.AddTree(t)
			}

			continue
		}

		h, err := hashFile(path)

		if err != nil {
			panic(err)
		}

		tree.AddBlob(Blob{
			Name: file.Name(),
			Path: path,
			Hash: h,
		})
	}

	return tree
}

// HashFile returns the sha1 hash of a file. The type will always be a "blob"
func hashFile(path string) (string, error) {

	b, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return cad.Hash(b, BlobType), nil
}

// contains returns true if a specific string is present in a string slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// snapshotHead recreates the folder structure for the current head
func snapshotHead(head, basePath string, db *cad.Cad) Tree {
	commit := mustFetchCommit(head, db)
	return mustFetchTree(commit.TreeHash, basePath, db)
}

func mustFetchTree(hash, basePath string, db *cad.Cad) Tree {
	b, _, err := db.Read(hash)
	if err != nil {
		panic(err)
	}

	tree, err := TreeFromBytes(
		b,
		basePath,
		func(db *cad.Cad) func(hash string) []byte {
			return func(hash string) []byte {
				b, _, err := db.Read(hash)
				if err != nil {
					panic(err)
				}

				return b
			}
		}(db),
	)

	if err != nil {
		panic(err)
	}

	return *tree
}
