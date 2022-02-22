// Package `objects` contains structs that can be saved to the content-addressable storage
// which are currently trees, blobs and commits
package objects

type Blob struct {
	Name string
	Hash string
	Path string
}
