package objects

import "fmt"

type object interface {
	getName() string
	fmt.Stringer
}

const (
	TreeType   = "tree"
	BlobType   = "blob"
	CommitType = "commit"
)

const (
	permsBlob = "100644"
	permsTree = "040000"
)
