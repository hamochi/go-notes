{
"CreatedAt": "2022-04-28",
"UpdatedAt": "2022-04-29",
"Title": "IO - Custom reader - Go notes"
}
---
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

<<custom_reader.go>>
