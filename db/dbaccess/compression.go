package dbaccess

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
)

func Compress(s []byte) []byte {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(s)
	w.Close()
	return b.Bytes()
}

func Decompress(s []byte) ([]byte, error) {
	b := bytes.NewBuffer(s)
	r, err := zlib.NewReader(b)
	if err != nil {
		return []byte{}, err
	}
	result, err := ioutil.ReadAll(r)
	r.Close()
	if err != nil {
		return []byte{}, err
	}
	return result, nil
}
