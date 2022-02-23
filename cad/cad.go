// Package `cad` provides read and write access to the content-addressable storage
package cad

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Cad struct {
	basePath string
}

func New(basePath string) *Cad {
	return &Cad{basePath}
}

func (cad *Cad) Write(in []byte, relativeOutPath string) error {

	path := filepath.Join(cad.basePath, relativeOutPath)

	_, err := os.Stat(path)

	// ignore if the file already exists
	if !os.IsNotExist(err) {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(path), 0644); err != nil {
		return err
	}

	c, err := compress(in)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, c, 0644); err != nil {
		return err
	}

	return nil
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
