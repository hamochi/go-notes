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
