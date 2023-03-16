package copying

import (
	"os"
	"strings"
	"testing"
)

func TestCustomReader(t *testing.T) {
	reader := strings.NewReader("The quick brown fox jumps over the lazy dog")
	upperCaseReader := NewUpperCaseReader(reader)
	_, err := os.Stdout.ReadFrom(upperCaseReader)
	if err != nil {
		t.Fatal(err)
	}
}
