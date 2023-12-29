
# Custom Error
In Go you can create a custom error by implementing the Error interface 
```go
type error interface {
	Error() string
}
```
Latly I've been using a custom error type that wraps errors to show from where they are coming from. It works by whenever I want to return an error up the stack, I wrap it in my custom error type. Here is a simple example of how it can be used.

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

