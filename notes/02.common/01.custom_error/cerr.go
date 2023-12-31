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
