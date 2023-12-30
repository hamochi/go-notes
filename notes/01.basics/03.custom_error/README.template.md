{
"CreatedAt": "2023-12-29",
"UpdatedAt": "2023-12-29",
"Title": "Custom error - Go notes"
}
---
# Custom Error
In Go you can create a custom error by implementing the Error interface 
```go
type error interface {
	Error() string
}
```
Recently, I've started using a custom error type that wraps errors, providing clear information about their source. Whenever I need to return an error up the stack, I encapsulate it in this custom error type. Below is a simple example demonstrating its application.

<<cerr_test.go>>

### Output
```bash
cerr_test.go->One:26 something went wrong with user 123 - error 1
cerr_test.go->Two:32 cerr_test.go->Three:37 something went wrong with user 123 John Doe - error 2
cerr_test.go->Four:41 something went wrong
```


## Implementation

<<cerr.go>>