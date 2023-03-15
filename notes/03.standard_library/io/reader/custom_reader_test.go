package reader

import (
	"fmt"
	"io"
	"testing"
)

func TestCustomReader(t *testing.T) {
	reader := NewUpperCaseReader("The quick brown fox jumps over the lazy dog")
	p := make([]byte, 3)
	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		fmt.Println(string(p), n)
	}
}
