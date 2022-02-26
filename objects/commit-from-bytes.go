package objects

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
	tree 17133c7b0083abebdadc1e73b15a2a3f2707819d
	parent e9dc2953af87550d37a064a998b7875d33b97df4
	author anonymous <anonymous> 1645906671
	committer anonymous <anonymous> 1645906671

	default message
*/

// CommitFromBytes creates a commit from bytes. The bytes
// represent the string value received by the `Value` method
// TODO: TEST
func CommitFromBytes(in []byte) (*Commit, error) {

	// lines
	const (
		treeKw      = "tree"
		parentKw    = "parent"
		authorKw    = "author"
		committerKw = "committer"
	)

	s := bufio.NewScanner(bytes.NewReader(in))

	var (
		tree      string
		parent    string
		user      *user
		timestamp time.Time
	)

	var line int
	var skippedEmpty bool
	var message string
	for s.Scan() {
		line++
		txt := s.Text()

		if strings.HasPrefix(txt, treeKw) && isMeta(line) {
			tree = strings.TrimSpace(txt[len(treeKw):])
			continue
		} else if strings.HasPrefix(txt, parentKw) && isMeta(line) {
			parent = strings.TrimSpace(txt[len(parentKw):])
			continue
		} else if strings.HasPrefix(txt, authorKw) && isMeta(line) {
			user, timestamp = userAndTimeFromMeta(txt[len(authorKw):])
			continue
		} else if strings.HasPrefix(txt, committerKw) && isMeta(line) {
			// skip because author and committer will same since no cherry-picking is implemented
			continue
		} else if strings.TrimSpace(txt) == "" && isMeta(line) {
			skippedEmpty = true // Skip empty line only once
			continue
		}

		// skip empty line if not already skipped
		if !skippedEmpty && strings.TrimSpace(txt) == "" {
			skippedEmpty = true
			continue
		}

		message += txt
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return &Commit{
		TreeHash:   tree,
		ParentHash: parent,
		Message:    message,
		Timestamp:  timestamp,
		user:       user,
	}, nil
}

// isMeta returns true if the current line is a metadata line.
// This is tedious because the parent is an optional value. If the
// parent is present the 'committer' will be on line 4 - otherweise
// the line 4 will be empty. Everything after that (which is not an empty line)
// will be the message. We check if we are int he first 4 lines when parsing
// the `Commit` to a void a message being misinterpreted as a meta property. This
// could happen if the message starts with any of the meta keywords (tree, parent...)
func isMeta(line int) bool {
	const maxMetadataLines = 4
	return line <= maxMetadataLines
}

// userFromMeta returns a `user` struct from a meta string
// author anonymous <anonymous@anonymous> 1645906671
func userAndTimeFromMeta(line string) (*user, time.Time) {
	parts := strings.Fields(line)

	emailPrefixAndSuffixRegexp := regexp.MustCompile(`<|>`)

	u := user{
		username: parts[0],
		email:    emailPrefixAndSuffixRegexp.ReplaceAllString(parts[1], ""),
	}

	timestamp := parts[2]

	// should never fail...
	timestampInt, _ := strconv.ParseInt(timestamp, 10, 64)

	return &u, time.Unix(timestampInt, 0)
}
