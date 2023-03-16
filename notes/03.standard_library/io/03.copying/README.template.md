{
"CreatedAt": "2023-03-25",
"UpdatedAt": "2023-03-15",
"Title": "IO - Writing - Go notes"
}
---
# IO streaming
In Go we can model data as streams. We can stream data from sources with io.Reader and save data to destinations with io.Writer. Sources can be  files, network connections, std out, strings etc.

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
io.Copy is a convenient function that will either use the ReadFrom or Write to function. 

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

## Chaining Reader to Reader

Lets say we want to transform the data we read before we pass it to a writer. We could for example read the content of a stream make it into uppercase and write it to std out.

Read -> Transform -> Read -> Write

We can rewrite our [custom Reader](#) that it accepts a reader as source.

<<custom_reader2.go>>

An test it:
<<custom_reader2_test.go>>

Output:
```bash
THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG-
```