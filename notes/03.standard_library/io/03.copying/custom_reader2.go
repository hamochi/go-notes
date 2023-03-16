package copying

import (
	"bytes"
	"io"
)

type UpperCaseReader struct {
	src io.Reader
}

func NewUpperCaseReader(src io.Reader) *UpperCaseReader {
	return &UpperCaseReader{
		src: src,
	}
}

func (u *UpperCaseReader) Read(p []byte) (int, error) {
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
