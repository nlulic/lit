package objects

import (
	"bufio"
	"bytes"
	"path/filepath"
	"strings"
)

/*
100644 blob 6194df4de654f46c609780a00bfe9a501b020d7e    .litignore
040000 tree 694f0f2fb5db3e2b2815bbdd16562198060d0e4e    cad
100644 blob f2f2ec4d073b0e08211aa09e86e185e4e1cc6204    cmd_branch.go
100644 blob 2ef8c8f26452717b58a666ae3b5030ae4037a39b    cmd_checkout.go
040000 tree 62e7cb1246b1ccd2d843f07dc9b0b581b1a0ad74    logger
040000 tree f9bfc3df0a802d7443e65d87b3c370aa826ef3ed    util
*/

// TreeFromBytes creates a `Tree` from bytes. The bytes represent
// the string value received by the `Value` method of the `Tree`.
// The basePath is required so the paths will match the FS
func TreeFromBytes(in []byte, basePath string, nextTree func(hash string) []byte) (*Tree, error) {

	baseTree := Tree{
		Name: filepath.Base(basePath),
		Path: basePath,
	}

	s := bufio.NewScanner(bytes.NewReader(in))

	for s.Scan() {
		parts := strings.Fields(s.Text())

		var (
			objType = parts[1]
			objHash = parts[2]
			objName = parts[3]
		)

		switch objType {
		case TreeType:
			tree, err := TreeFromBytes(
				nextTree(objHash),
				filepath.Join(basePath, objName),
				nextTree,
			)

			if err != nil {
				return nil, err
			}

			baseTree.AddTree(*tree)
		case BlobType:
			baseTree.AddBlob(Blob{
				Name: objName,
				Path: filepath.Join(basePath, objName),
				Hash: objHash,
			})
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return &baseTree, nil
}
