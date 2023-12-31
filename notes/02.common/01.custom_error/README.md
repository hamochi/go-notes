
# Custom Error
In Go you can create a custom error by implementing the Error interface 
```go
type error interface {
	Error() string
}
```
Recently, I've started using a custom error type that wraps errors, providing clear information about their source. Whenever I need to return an error up the stack, I encapsulate it in this custom error type. Below is a simple example demonstrating its application.

```go
package cerr_test

import (
	"fmt"
	cerr "go-notes/notes/01.basics/03.custom_error"
	"testing"
)

var userId = 123
var userName = "John Doe"

func TestErrors(t *testing.T) {
	err := One()
	fmt.Println(err)
	err = Two()
	fmt.Println(err)

	err = Four()
	fmt.Println(err)
}

func One() error {
	var err1 = fmt.Errorf("error 1")

	// Add some context to the error
	return cerr.New(err1, "something went wrong with user", userId)
}

func Two() error {
	err := Three()

	// Simply wrap the error
	return cerr.New(err)
}

func Three() error {
	var err2 = fmt.Errorf("error 2")

	// Add some formatted context to the error
	return cerr.Newf(err2, "something went wrong with user %d %s", userId, userName)
}

func Four() error {
	// Create a new error
	return cerr.NewError("something went wrong")
}
```

### Output
```bash
cerr_test.go->One:26 something went wrong with user 123 - error 1
cerr_test.go->Two:32 cerr_test.go->Three:37 something went wrong with user 123 John Doe - error 2
cerr_test.go->Four:41 something went wrong
```


## Implementation

```go
package cerr

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type CustomError struct {
	message  string
	fileName string
	line     int
	err      error
}

func New(err error, message ...any) *CustomError {
	fileName, line := getCallerInfo()
	finalMessage := strings.Builder{}

	for i := range message {
		if i > 0 {
			finalMessage.WriteString(" ")
		}
		finalMessage.WriteString(fmt.Sprint(message[i]))
	}

	return &CustomError{
		message:  finalMessage.String(),
		fileName: fileName,
		line:     line,
		err:      err,
	}
}

func Newf(err error, format string, a ...any) *CustomError {
	fileName, line := getCallerInfo()
	message := fmt.Sprintf(format, a...)
	return &CustomError{
		message:  message,
		fileName: fileName,
		line:     line,
		err:      err,
	}
}

func NewError(message string) *CustomError {
	err := fmt.Errorf(message)
	fileName, line := getCallerInfo()
	return &CustomError{
		message:  "",
		fileName: fileName,
		line:     line,
		err:      err,
	}
}

func (e *CustomError) Error() string {
	if e.message == "" {
		return fmt.Sprintf("%s:%d %s", e.fileName, e.line, e.err.Error())
	}
	return fmt.Sprintf("%s:%d %s - %s", e.fileName, e.line, e.message, e.err.Error())
}

func (e *CustomError) Unwrap() error {
	return e.err
}

func getCallerInfo() (file string, line int) {
	pc, file, line, ok := runtime.Caller(2) // use 2 as argument to retrieve information about the calling function
	if !ok {
		return "unknown", 0
	}

	caller := runtime.FuncForPC(pc)
	if caller == nil {
		return file, line
	}
	fullFileName, line := caller.FileLine(pc)
	fileName := filepath.Base(fullFileName) // Get just the filename
	fullFunctionName := caller.Name()
	functionNameParts := strings.Split(fullFunctionName, ".")
	functionName := functionNameParts[len(functionNameParts)-1]

	file = fmt.Sprintf("%v->%v", fileName, functionName)

	return file, line
}
```