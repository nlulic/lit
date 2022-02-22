// Package `objects` contains structs that can be saved to the content-addressable storage
// which are currently trees, blobs and commits
package objects

import (
	"encoding/json"
	"fmt"
)

type Tree struct {
	Name  string
	Blobs []Blob
	Trees []Tree
	Path  string
}

func (t *Tree) AddBlob(blob Blob) {
	t.Blobs = append(t.Blobs, blob)
}

func (t *Tree) AddTree(tree Tree) {
	t.Trees = append(t.Trees, tree)
}

func (t *Tree) IsEmpty() bool {
	return len(t.Blobs) == 0 && len(t.Blobs) == 0
}

// TODO: Remove
func (t *Tree) Print() {
	s, _ := json.MarshalIndent(t, "", "\t")
	fmt.Printf("%s\n", string(s))
}
