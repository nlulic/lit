package main

import (
	"fmt"
	"path/filepath"
	"time"

	"gitlab.com/nlulic/lit/cad"
	"gitlab.com/nlulic/lit/objects"
)

func (lit *Lit) Log() {
	head, err := lit.GetHead()
	ref, _ := lit.GetRef()

	if err != nil {
		if err == RepositoryNotInitialized {
			lit.logger.Fatal(err)
		}

		lit.logger.Fatal(fmt.Sprintf("fatal: current branch '%s' does not have any commits", filepath.Base(ref)))
	}

	objectsDir := lit.objectsDir()
	db := cad.New(objectsDir)

	commit := mustFetchCommit(head, db)
	commits := []*objects.Commit{commit}

	for {
		if commit.ParentHash == "" {
			break
		}

		commit = mustFetchCommit(commit.ParentHash, db)
		commits = append(commits, commit)
	}

	for _, commit := range commits {
		lit.logger.Info(
			fmt.Sprintf("commit %s (HEAD -> %s)\n", commit.Hash(), filepath.Base(ref)) +
				fmt.Sprintf("Author: %s <%s>\n", commit.User().Username(), commit.User().Email()) +
				fmt.Sprintf("Date:   %s\n\n", commit.Timestamp.Format(time.RFC1123)) +
				fmt.Sprintf("\t%s\n", commit.Message),
		)
	}
}

// mustFetchCommit will fetch the commit from the cad based on the hash or panic
func mustFetchCommit(hash string, db *cad.Cad) *objects.Commit {
	value, _, err := db.Read(hash)
	if err != nil {
		panic(err)
	}

	commit, err := objects.CommitFromBytes(value)

	if err != nil {
		panic(err)
	}

	return commit
}
