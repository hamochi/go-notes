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
