{
"CreatedAt": "2022-04-28",
"UpdatedAt": "2023-03-15",
"Title": "IO - Reading - Go notes"
}
---
# IO streaming
In Go we can model data as streams. We can stream data from sources with io.Reader and save data to destinations with io.Writer. Sources can be  files, network connections, std out, strings etc.

1. [Reader](#reader)
  1. [Reading from a reader](#reading-from-a-reader)
  2. [Reading all at once](#reading-all-at-once)
  3. [Creating a custom reader](#creating-a-custom-reader)
2. [Writer](#writer)
  1. [Writing to a Writer](#writing-to-a-writer)
  2. [Writing a string to a writer](#writing-a-string-to-a-writer)
  3. [Writing with fmt](#writing-with-fmt)
3. [Chaining from Reader to Writer](#chaining-from-reader-to-writer)
  1. [Using ReadFrom](#using-readfrom)
  2. [Using WriteTo](#using-writeto)
  3. [Using io.Copy](#using-iocopy)


## Reader
A reader reads data from the source, loads it in a buffer and returns how many bytes it has written to the buffer. This is usually done in a loop until the reader returns an EOT error.

[Data source] -> [io.Reader] -> [Transfer buffer []byte]

The reader interface in the io package.
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### Reading from a Reader
In order to read from a Reader we just call the Read() method in a loop and load chunks of data in a byte array until we get an EOT error. Let's look at a simple example where we convert a string to a Reader.

```bash 
stringreader.go
```
<<stringreader.go>>

```bash 
stringreader_test.go
```
<<stringreader_test.go>>

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

### Reading all at once
If we want to all the data at once, without looping we can use the the ReadAll function from the io package.

```go
// io.ReadAll()
func ReadAll(r Reader)([][byte] (error)
```

```bash
stringreader_readall_test.go
```
<<stringreader_readall_test.go>>

Output:
```bash
The quick brown fox jumps over the lazy dog
```

### Creating a custom reader 
The easiet way of creating a custom reader is to base it on an existing reader. We create a struct that embeds a reader, and add a Read() method to the struct, where we read from the source and transform the output. Here is an example of a custom reader that coverts the output a text to upper case.
```bash
custom_reader2.go
```
<<custom_reader2.go>>

Let's test it
```bash
custom_reader_test.go
```
<<custom_reader2_test.go>>

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

## Writer
Let's look at the writer interface:

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

Any type that implements the above Write method is a Writer. A few examples of writers are:
- bytes.Buffer
- os.File
- os.Stdout (which is of type os.File)
- http.Header
- http.ResponseWriter

#### Writing to a Writer
So a simple example of writing to a write could look like:

```go
package main

import (
	"os"
)

func main() {
	os.Stdout.Write([]byte("hello\n"))
	os.Stdout.Write([]byte("world"))
}
```

Output:
```bash
hello
world
```

#### Writing a string to a writer
In a lot of cases you might want to write a string to a Writer. You can of course cast the string to []byte but there is a convenient function in the io package called WriteString:

```go
func WriteString(w Writer, s string) (n int, err error)
```

The above example would be:

```go
package main

import (
	"os"
	"io"
)

func main() {
	io.WriteString(os.Stdout, "hello\n")
	io.WriteString(os.Stdout, "world")
}
```

Output:
```bash
hello
world
```

### Writing with fmt
You can also some of the functions from the fmt package that uses the a io.Writer as an argument to write to it.

```go
// fmt package

func Fprint(w io.Writer, a ...any) (n int, err error)
func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
func Fprintln(w io.Writer, a ...any) (n int, err error)
```

Above example could look like:
```go
package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Fprintf(os.Stdout, "%s\n%s", "hello", "world")
}
```

Read more about it in the [fmt section](#).

## Chaining from Reader to Writer 
It's important to remember when we read from a Reader, that data is no longer available. We are basically consuming a stream, one chunk at the time. 

There are two other interfaces we need to take a look at.
```go
type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}

type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}
```

There are two ways of passing data from a Reader to a Writer. Either the source (implementing the Reader) implements the WriterTo interface, or the destination (implementing the Writer) implements the ReadFrom interface.

### Using ReadFrom
 ```go
package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")
	_, err := os.Stdout.ReadFrom(r)
	if err != nil {
		log.Fatal(err)
	}
}
```

Output:
```bash
some io.Reader stream to be read
```

### Using WriteTo
```go
package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	_, err := r.WriteTo(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	
}
```

Output:
```bash
some io.Reader stream to be read
```

### Using io.Copy
io.Copy is a convenient function that will either use the ReadFrom or the WriteTo function. 

```go
// io package

func Copy(dst Writer, src Reader) (written int64, err error)
```

The official description:

```
If src implements the WriterTo interface, the copy is implemented by calling 
src.WriteTo(dst). Otherwise, if dst implements the ReaderFrom interface, 
the copy is implemented by calling dst.ReadFrom(src).
```

Lets look at the same as above but with io.Copy:
```go
package main

import (
	"log"
	"os"
	"strings"
	"io"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
	
}
```

Output:
```bash
some io.Reader stream to be read
```


