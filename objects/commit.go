// Package `objects` contains structs that can be saved to the content-addressable storage
// which are currently trees, blobs and commits
package objects

import "time"

type Commit struct {
	Tree       *Tree
	ParentHash string
	User       *user
	Message    string
	Timestamp  time.Time
}

// Only for completeness - the user will always be anonymous
type user struct {
	username string
	email    string
}
