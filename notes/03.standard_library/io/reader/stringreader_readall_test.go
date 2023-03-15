package reader

import (
	"fmt"
	"io"
	"testing"
)

func TestReaderFromStringAll(t *testing.T) {
	reader := ReaderFromString("The quick brown fox jumps over the lazy dog")
	fullBuf, error := io.ReadAll(reader)
	if error != nil {
		t.Fatal(error)
	}
	fmt.Println(string(fullBuf))
}
