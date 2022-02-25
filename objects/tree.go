// Package `objects` contains structs that can be saved to the content-addressable storage
// which are currently trees, blobs and commits
package objects

import (
	"fmt"
	"sort"

	"gitlab.com/nlulic/lit/cad"
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

// Value returns the formatted data as a slice of bytes. The
// result is sorted to make sure it is consistent on every fs
func (t *Tree) Value() string {

	var objects []object

	for _, blob := range t.Blobs {
		objects = append(objects, blob)
	}

	for _, tree := range t.Trees {
		objects = append(objects, tree)
	}

	// sort blobs and trees by name
	sort.Slice(objects, func(left, right int) bool {
		return objects[left].getName() < objects[right].getName()
	})

	var value string
	for _, object := range objects {
		value += object.String()
	}

	return value
}

func (t Tree) String() string {
	return fmt.Sprintf("%s %s %s    %s\n", permsTree, TreeType, t.Hash(), t.Name)
}

func (t *Tree) Hash() string {
	return cad.Hash([]byte(t.Value()), TreeType)
}

// getName returns the name of the Tree to implement the object interface
func (t Tree) getName() string {
	return t.Name
}
