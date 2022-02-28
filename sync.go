package main

import (
	"os"

	"gitlab.com/nlulic/lit/cad"
	. "gitlab.com/nlulic/lit/objects"
)

// Sync applies the snapshot to the working directory
func (lit *Lit) Sync(snapshot, workingTree Tree) error {

	if snapshot.Hash() == workingTree.Hash() {
		lit.logger.Debug("nothing changed in", snapshot.Name)
		return nil
	}

	objectsDir := lit.objectsDir()
	db := cad.New(objectsDir)

	// synchronize blobs
	baseBlobs := blobMap(workingTree.Blobs)

	for _, blob := range snapshot.Blobs {
		if b, ok := baseBlobs[blob.Path]; ok {
			delete(baseBlobs, b.Path)

			if b.Hash == blob.Hash {
				lit.logger.Debug(blob.Name, "did not change. Skipping...")
				continue
			}

			// update
			lit.logger.Debug(blob.Name, "changed. Updating...")
			mustCreateFile(blob, db)
			continue
		}

		// doesn't exist in the working directory
		lit.logger.Debug(blob.Name, "doesn't exist anymore. Creating...")
		mustCreateFile(blob, db)
	}

	// leftover blobs that do not exist in the HEAD snapshot will be deleted
	for _, blob := range baseBlobs {
		lit.logger.Debug(blob.Name, "has been removed. Deleting...")
		mustDeleteFile(blob)
	}

	// synchronize trees
	baseTrees := treeMap(workingTree.Trees)

	for _, tree := range snapshot.Trees {
		if baseTree, ok := baseTrees[tree.Path]; ok {
			delete(baseTrees, tree.Path)

			if baseTree.Hash() == tree.Hash() {
				lit.logger.Debug(tree.Name, "did not change. Skipping...")
				continue
			}

			// update
			lit.logger.Debug(tree.Name, "changed. Updating...")
			lit.Sync(tree, baseTree)
			continue
		}

		// doesn't exist in the working directory
		lit.logger.Debug(tree.Name, "doesn't exist anymore. Creating...")
		lit.Sync(tree, Tree{})
	}

	// leftover trees that do not exist in the HEAD snapshot will be deleted
	for _, tree := range baseTrees {
		lit.logger.Debug(tree.Name, "has been removed. Deleting...")
		mustDeleteDir(tree)
	}

	return nil
}

func mustCreateFile(blob Blob, db *cad.Cad) {

	b, _, err := db.Read(blob.Hash)
	if err != nil {
		panic(err)
	}

	mustWriteToFile(blob.Path, string(b))
}

func mustDeleteFile(blob Blob) {
	if err := os.Remove(blob.Path); err != nil {
		panic(err)
	}
}

func mustDeleteDir(tree Tree) {
	if err := os.Remove(tree.Path); err != nil {
		panic(err)
	}
}

func blobMap(blobs []Blob) map[string]Blob {
	m := make(map[string]Blob)
	for _, blob := range blobs {
		m[blob.Path] = blob
	}
	return m
}

func treeMap(trees []Tree) map[string]Tree {
	m := make(map[string]Tree)
	for _, tree := range trees {
		m[tree.Path] = tree
	}
	return m
}
