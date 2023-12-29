package reading

import (
	"io"
	"strings"
)

func ReaderFromString(text string) io.Reader {
	return strings.NewReader(text)
}
