package goutils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type HttpRequestParams struct {
	HttpRequestType string
	Url             string
	Data            interface{}
	Headers         map[string]string
}

func HttpGetRequest(url string) []byte {
	resp, err := http.Get(url)
	CheckForErrors(ErrorParams{Err: err, CallerNum: 1})

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckForErrors(ErrorParams{Err: err, CallerNum: 1})

	return body
}

func HttpCreateRequest(p HttpRequestParams) (int, bytes.Buffer) {
	var req *http.Request
	var statusCode int
	var dataBytes, bodyBytes bytes.Buffer

	switch v := p.Data.(type) {
	case string:
		dataBytes = *bytes.NewBufferString(v)
	case []byte:
		dataBytes = *bytes.NewBuffer(v)
	}

	req, _ = http.NewRequest(p.HttpRequestType, p.Url, &dataBytes)

	for k, v := range p.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	CheckForErrors(ErrorParams{Err: err, CallerNum: 1})

	switch resp.StatusCode {
	// etcd server is on redirect
	case http.StatusTemporaryRedirect:
		u, err := resp.Location()

		if err != nil {
			CheckForErrors(ErrorParams{Err: err, CallerNum: 1})
		} else {
			p.Url = u.String()
			HttpCreateRequest(p)
		}
	default:
		statusCode = resp.StatusCode

		body, err := ioutil.ReadAll(resp.Body)
		CheckForErrors(ErrorParams{Err: err, CallerNum: 1})
		bodyBytes = *bytes.NewBuffer(body)
	}
	defer resp.Body.Close()
	return statusCode, bodyBytes
}
