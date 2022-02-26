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

	if err := os.MkdirAll(filepath.Dir(path), 0644); err != nil {
		return "", err
	}

	c, err := compress(withHeader(in, objectType))

	if err != nil {
		return "", err
	}

	if err := ioutil.WriteFile(path, c, 0644); err != nil {
		return "", err
	}

	return hash, nil
}

func (cad *Cad) Read(relativePath string) ([]byte, error) {

	path := filepath.Join(cad.basePath, relativePath)

	compressed, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	u, err := uncompress(compressed)

	return u, err
}

func relativePath(hash string) string {
	return filepath.Join(hash[:2], hash[2:])
}
