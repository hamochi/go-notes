package reader

import (
	"fmt"
	"io"
	"os"
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

func main() {
	reader := NewUpperCaseReader("The quick brown fox jumps over the lazy dog")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(string(p[:n]))
	}
}
