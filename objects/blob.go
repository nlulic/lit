// Package `objects` contains structs that can be saved to the content-addressable storage
// which are currently trees, blobs and commits
package objects

import "fmt"

type Blob struct {
	Name string
	Hash string
	Path string
}

func (b Blob) String() string {
	return fmt.Sprintf("%s %s %s    %s\n", permsBlob, BlobType, b.Hash, b.Name)
}

// getName returns the name of the Blob to implement the objects interface
func (b Blob) getName() string {
	return b.Name
}
