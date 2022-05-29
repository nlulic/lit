// Package `cad` provides read and write access to the content-addressable storage
package cad

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ObjectAlreadyExists = errors.New("warn: object already exists")
)

type Cad struct {
	basePath string
}

func New(basePath string) *Cad {
	return &Cad{basePath}
}

func (cad *Cad) Write(in []byte, objectType string) (string, error) {

	hash := Hash(in, objectType)
	path := filepath.Join(cad.basePath, relativePath(hash))

	_, err := os.Stat(path)

	// ignore if the file already exists
	if !os.IsNotExist(err) {
		return hash, ObjectAlreadyExists
	}

	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return "", err
	}

	c, err := compress(withHeader(in, objectType))

	if err != nil {
		return "", err
	}

	if err := ioutil.WriteFile(path, c, 0664); err != nil {
		return "", err
	}

	return hash, nil
}

// Read will fetch a stored object by hash and return the objects header and value
func (cad *Cad) Read(hash string) ([]byte, []byte, error) {

	path := filepath.Join(cad.basePath, relativePath(hash))

	compressed, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	u, err := uncompress(compressed)

	if err != nil {
		return nil, nil, err
	}

	// separate header from the content
	var nullByteIndex int
	var null_byte byte

	for i, b := range u {
		if b == null_byte {
			nullByteIndex = i
			break
		}
	}

	header := u[:nullByteIndex]
	value := u[nullByteIndex+1:]

	return value, header, nil
}

func relativePath(hash string) string {
	return filepath.Join(hash[:2], hash[2:])
}
