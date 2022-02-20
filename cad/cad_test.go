package cad

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRW(t *testing.T) {

	tempDir, err := ioutil.TempDir("", t.Name())

	if err != nil {
		t.Fatalf("TempDir %s: %v", t.Name(), err)
	}

	defer os.RemoveAll(tempDir)

	tests := []struct {
		name string
		path string
		text string
		want []byte
	}{
		{
			name: "Create File",
			path: "foo/bar",
			text: "de07f660-02d5-46be-b79e-020f909acedc",
			want: []byte("de07f660-02d5-46be-b79e-020f909acedc"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cad := NewCad(tempDir)

			err := cad.Write([]byte(tt.text), tt.path)
			if err != nil {
				t.Error(err)
			}

			_, err = os.Stat(filepath.Join(tempDir, tt.path))
			if err != nil {
				t.Error(err)
			}

			b, err := cad.Read(tt.path)
			if err != nil {
				t.Error(err)
			}

			if !bytes.Equal(tt.want, b) {
				t.Errorf("want %q; got %q", string(tt.want), string(b))
			}
		})
	}

}

func TestCompressUncompress(t *testing.T) {

	tests := []struct {
		name string
		text string
		want []byte
	}{
		{
			name: "Compress String",
			text: "9322940b-20a7-4fa1-82fd-c4f98cc0c190",
			want: []byte("9322940b-20a7-4fa1-82fd-c4f98cc0c190"),
		},
		{
			name: "Empty",
			text: "",
			want: []byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := compress([]byte(tt.text))

			if err != nil {
				t.Error(err)
			}

			u, err := uncompress(c)

			if err != nil {
				t.Error(err)
			}

			if !bytes.Equal(tt.want, u) {
				t.Errorf("want %q; got %q", tt.want, string(u))
			}
		})
	}
}
