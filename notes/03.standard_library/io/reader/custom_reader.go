package reader

import (
	"io"
	"strings"
)

type UpperCaseReader struct {
	src string
	cur int
}

func NewUpperCaseReader(src string) *UpperCaseReader {
	return &UpperCaseReader{
		src: src,
	}
}

func upperCase(b byte) byte {
	// let's use built-in converter and return the first byte (since we don't expect more than on byte)
	return strings.ToUpper(string(b))[0]
}

func (u *UpperCaseReader) Read(p []byte) (int, error) {
	if u.cur >= len(u.src) {
		return 0, io.EOF
	}

	buf := make([]byte, len(p))
	n := 0
	for {
		if n >= len(p) || u.cur >= len(u.src) {
			break
		}
		buf[n] = upperCase(u.src[u.cur])
		u.cur++
		n++
	}

	copy(p, buf)
	return n, nil
}
