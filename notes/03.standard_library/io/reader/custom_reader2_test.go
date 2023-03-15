package reader

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestCustomReader2(t *testing.T) {
	reader := strings.NewReader("The quick brown fox jumps over the lazy dog")
	upperCaseReader := NewUpperCaseReader2(reader)
	p := make([]byte, 3)
	for {
		n, err := upperCaseReader.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		fmt.Println(string(p), n)
	}
}
