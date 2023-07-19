// This is a stub package that takes the place of a custom made http solution
package httpclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/go-retryablehttp"
	// "github.com/monaco-io/request"
)

func formatError(msg string, err error) error {
	return fmt.Errorf("httpclient error: could not %s (%v)", msg, err)
}
func GetWithSearchString[responseStruct any](url string, searchString string) (*responseStruct, error) {
	response := new(responseStruct)
	client := retryablehttp.NewClient()
	client.RetryMax = 3
	fullUrl := fmt.Sprintf("%s?search=%s", url, searchString)
	if res, err := retryablehttp.Get(fullUrl); err != nil {
		return response, formatError("complete request", err)
	} else {
		defer res.Body.Close()
		if body, err := ioutil.ReadAll(res.Body); err != nil {
			return response, formatError("read body", err)
		} else {
			if err = json.Unmarshal(body, &response); err != nil {
				return response, formatError("unmarshal body", err)
			}
		}
	}
	return response, nil
}
func Get[responseStruct any](url string) (*responseStruct, error) {
	response := new(responseStruct)
	client := retryablehttp.NewClient()
	client.RetryMax = 3
	if res, err := retryablehttp.Get(url); err != nil {
		return response, formatError("complete request", err)
	} else {
		defer res.Body.Close()
		if body, err := ioutil.ReadAll(res.Body); err != nil {
			return response, formatError("read body", err)
		} else {
			if err = json.Unmarshal(body, &response); err != nil {
				return response, formatError("unmarshal body", err)
			}
		}
	}
	return response, nil
}
