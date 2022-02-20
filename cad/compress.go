package cad

import (
	"bytes"
	"compress/zlib"
)

func compress(in []byte) ([]byte, error) {

	var buff bytes.Buffer

	w := zlib.NewWriter(&buff)

	if _, err := w.Write(in); err != nil {
		return nil, err
	}
	w.Close()

	return buff.Bytes(), nil
}

func uncompress(in []byte) ([]byte, error) {

	r, err := zlib.NewReader(bytes.NewBuffer(in))

	if err != nil {
		return nil, err
	}

	var buff bytes.Buffer

	_, err = buff.ReadFrom(r)
	r.Close()

	if err != nil {
		return nil, err
	}

	return buff.Bytes(), err
}
