package goutils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type HttpRequestParams struct {
	HttpRequestType string
	Url             string
	Data            string
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

func HttpCreateRequest(p HttpRequestParams) *http.Response {
	var req *http.Request

	var dataBytes = bytes.NewBufferString(p.Data)
	req, _ = http.NewRequest(p.HttpRequestType, p.Url, dataBytes)

	for k, v := range p.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	CheckForErrors(ErrorParams{Err: err, CallerNum: 1})

	defer resp.Body.Close()

	// etcd server is on redirect
	if resp.StatusCode == http.StatusTemporaryRedirect {
		u, err := resp.Location()

		if err != nil {
			CheckForErrors(ErrorParams{Err: err, CallerNum: 1})
		} else {
			p.Url = u.String()
			HttpCreateRequest(p)
		}
	}

	return resp
}
