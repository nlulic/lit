// Package `objects` contains structs that can be saved to the content-addressable storage
// which are currently trees, blobs and commits
package objects

import (
	"fmt"
	"time"

	"gitlab.com/nlulic/lit/cad"
)

type Commit struct {
	Tree       *Tree
	ParentHash string
	Message    string
	Timestamp  time.Time
	user       *user
}

// Only for completeness - the user will always be anonymous
type user struct {
	username string
	email    string
}

func NewCommit(msg string, tree *Tree, parentHash string) *Commit {
	return &Commit{
		Tree:       tree,
		ParentHash: parentHash,
		Message:    msg,
		Timestamp:  time.Now(),
		user: &user{
			username: "anonymous",
			email:    "anonymous",
		},
	}
}

func (c *Commit) Value() (value string) {

	value += fmt.Sprintf("tree %s\n", c.Tree.Hash())

	if c.ParentHash != "" {
		value += fmt.Sprintf("parent %s\n", c.ParentHash)
	}

	value += fmt.Sprintf("author %s <%s> %d\n", c.user.username, c.user.email, c.Timestamp.UTC().Unix())
	value += fmt.Sprintf("committer %s <%s> %d\n", c.user.username, c.user.email, c.Timestamp.UTC().Unix())

	value += "\n" + c.Message
	return
}

func (c *Commit) Hash() string {
	b := []byte(c.Value())
	return cad.Hash(b, CommitType)
}
