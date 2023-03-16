
# IO streaming
In Go we can model data as streams. We can stream data from sources with io.Reader and save data to destinations with io.Writer. Sources can be  files, network connections, std out, strings etc.

## Reader
A reader reads data from the source, loads it in a buffer and returns how many bytes it has written to the buffer. This is usually done in a loop until the reader returns an EOT error.

[Data source] -> [io.Reader] -> [Transfer buffer []byte]

The reader interface in the io package.
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### Read() - Reading from a Reader
In order to read from a Reader we just call the Read() method in a loop and load chunks of data in a byte array until we get an EOT error. Let's look at a simple example where we convert a string to a Reader.

```bash 
stringreader.go
```
```go
package reading

import (
	"io"
	"strings"
)

func ReaderFromString(text string) io.Reader {
	return strings.NewReader(text)
}
```

```bash 
stringreader_test.go
```
```go
package reading

import (
	"fmt"
	"io"
	"testing"
)

func TestReaderFromString(t *testing.T) {
	reader := ReaderFromString("The quick brown fox jumps over the lazy dog")
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
```

Output:
```bash
The 3
 qu 3
ick 3
 br 3
own 3
 fo 3
x j 3
ump 3
s o 3
ver 3
 th 3
e l 3
azy 3
 do 3
gdo 1

```

### ReadAll() - Reading all at once
If we want to all the data at once, without looping we can use the the ReadAll function from the io package.

```go
// io.ReadAll()
func ReadAll(r Reader)([][byte] (error)
```

```bash
stringreader_readall_test.go
```
```go
package reading

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
```

Output:
```bash
The quick brown fox jumps over the lazy dog
```

### Creating a custom reader from scratch
This is a rare case. We usually don't create custom readers from scratch but instead use an other reader as the source of our reader. 

Lets look at an example where we create a reader that changes a string to upper case.
```bash
custom_reader.go
```
```go
package reading

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
```

Let's test it
```bash
custom_reader_test.go
```
```go
package reading

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
```

Output:
```bash
THE 3
 QU 3
ICK 3
 BR 3
OWN 3
 FO 3
X J 3
UMP 3
S O 3
VER 3
 TH 3
E L 3
AZY 3
 DO 3
G 1
```

