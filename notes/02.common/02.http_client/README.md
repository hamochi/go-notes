
# HTTP Client
Here is a full example of sending a POST request, with headers, cookies and checking the response.

```go
package client

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func POST() ([]byte, error) {
	// We create a http client with timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// The body if the request is POST, we can skip this if request
	// method does not require a body, GET for example
	reqBody := strings.NewReader(`{"key": "value", "key2": "value2"}`)

	// Create request, we can use nil for body if we use GET
	req, err := http.NewRequest("POST", "http://localhost:9000", reqBody)
	if err != nil {
		return nil, err
	}

	// Add some header
	req.Header.Add("Content-Type", "application/json")

	// Add some cookie if you want, cookie is actually just another header
	// But we don't need to do string manipulations with this method
	cookie := &http.Cookie{
		Name:  "someKey",
		Value: "some value",
	}
	req.AddCookie(cookie)

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Make sure we close body
	defer res.Body.Close()

	// res.Body is an io.ReadeCloser, here we read all data at once
	// and store it as []byte
	bodyBuffer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Get status code from result and do some control
	statusCode := res.StatusCode
	if statusCode != http.StatusOK {
		return nil, errors.New("service did not return status 200")
	}

	return bodyBuffer, nil
}
```

## Post form data

```go
package client

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

func PostFormData() error {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	data := url.Values{}
	data.Set("fieldName1", "value1")
	data.Set("fieldName2", "value2")

	res, err := client.PostForm("https://www.somehwere.com/submit", data)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	statusCode := res.StatusCode
	if statusCode != http.StatusOK {
		return errors.New("unexpected status code")
	}

	return nil
}
```