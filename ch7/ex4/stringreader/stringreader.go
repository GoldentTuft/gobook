// Package stringreader は文字列からio.Readerを返す
// 完全に写経になってしまった
// https://github.com/ray-g/gopl/blob/master/ch07/ex7.04/stringreader.go
package stringreader

import (
	"io"
)

// StringReader は
type StringReader struct {
	s string
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

// NewReader は文字列からio.Readerを返す
func NewReader(s string) io.Reader {
	return &StringReader{s}
}
