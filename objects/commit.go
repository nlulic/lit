// Package `objects` contains structs that can be saved to the content-addressable storage
// which are currently trees, blobs and commits
package objects

import (
	"fmt"
	"time"

	"gitlab.com/nlulic/lit/cad"
)

type Commit struct {
	TreeHash   string
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

func (u *user) Username() string {
	return u.username
}

func (u *user) Email() string {
	return u.email
}

func NewCommit(msg string, tree *Tree, parentHash string) *Commit {
	return &Commit{
		TreeHash:   tree.Hash(),
		ParentHash: parentHash,
		Message:    msg,
		Timestamp:  time.Now(),
		user: &user{
			username: "anonymous",
			email:    "anonymous@anonymous",
		},
	}
}

func (c *Commit) Value() (value string) {

	value += fmt.Sprintf("tree %s\n", c.TreeHash)

	if c.ParentHash != "" {
		value += fmt.Sprintf("parent %s\n", c.ParentHash)
	}

	value += fmt.Sprintf("author %s <%s> %d\n", c.user.username, c.user.email, c.Timestamp.UTC().Unix())
	value += fmt.Sprintf("committer %s <%s> %d\n", c.user.username, c.user.email, c.Timestamp.UTC().Unix())

	value += "\n" + c.Message
	return
}

func (c *Commit) User() *user {
	return c.user
}

func (c *Commit) Hash() string {
	b := []byte(c.Value())
	return cad.Hash(b, CommitType)
}
