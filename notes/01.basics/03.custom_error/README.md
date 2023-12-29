
# Custom Error
In Go you can create a custom error by implementing the Error interface 
```go
type error interface {
	Error() string
}
```
Recently, I've started using a custom error type that wraps errors, providing clear information about their source. Whenever I need to return an error up the stack, I encapsulate it in this custom error type. Below is a simple example demonstrating its application.

```go
customerId := 1
customerName := "Frank"

// Simple wrapping and returning
somevalue, err := someFunctionThatCanReturnError()
if err != nil {
  return cerr.New(err)
}

// Adding some additional message
somevalue2, err := someFunction2ThatCanReturnError()
if err != nil {
  return cerr.New(err, "could not get value")
}


// Adding some additional messages
somevalue3, err := someFunction3ThatCanReturnError()
if err != nil {
  return cerr.New(err, "could not get value", customerId, "some other message")
}

// Adding some additional formated messages
somevalue4, err := someFunction4ThatCanReturnError()
if err != nil {
  return cerr.New(err, "could not get value for user %d, %s", customerId, customerName)
}

// Can also create a new error 
if customerId != 1 {
  return cerr.NewError("invalid customerId")
}
```

