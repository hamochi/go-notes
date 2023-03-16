
# IO streaming
In Go we can model data as streams. We can stream data from sources with io.Reader and save data to destinations with io.Writer. Sources can be  files, network connections, std out, strings etc.

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