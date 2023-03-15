package reader

import (
	"bytes"
	"io"
	"strings"
)

type UpperCaseReader2 struct {
	src io.Reader
}

func NewUpperCaseReader2(src io.Reader) *UpperCaseReader2 {
	return &UpperCaseReader2{
		src: src,
	}
}

func upperCase2(b byte) byte {
	// let's use built-in converter and return the first byte (since we don't expect more than on byte)
	return strings.ToUpper(string(b))[0]
}

func (u *UpperCaseReader2) Read(p []byte) (int, error) {
	// we create a new buffer to read into, same size as p
	b := make([]byte, len(p))

	// we read from src Reader
	n, err := u.src.Read(b)
	if err != nil {
		return 0, err
	}

	// we make it into upper case
	b = bytes.ToUpper(b)

	// and copy it into p
	copy(p, b)
	return n, nil
}
