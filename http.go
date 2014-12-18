package goutils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HttpGetRequest(url string) []byte {
	resp, err := http.Get(url)
	CheckForErrors(ErrorParams{Err: err, CallerNum: 1})

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckForErrors(ErrorParams{Err: err, CallerNum: 1})

	return body
}

func HttpPutRequest(urlStr string, data []byte) *http.Response {
	var req *http.Request

	req, _ = http.NewRequest("PUT", urlStr, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	CheckForErrors(ErrorParams{Err: err, CallerNum: 1})

	defer resp.Body.Close()

	return resp
}

func HttpPutRequestRedirect(urlStr string, data string) {
	var req *http.Request
	req, _ = http.NewRequest("PUT", urlStr, bytes.NewBufferString(data))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	client := &http.Client{}
	resp, err := client.Do(req)
	CheckForErrors(ErrorParams{Err: err, CallerNum: 1})

	if resp.StatusCode == http.StatusTemporaryRedirect {
		u, err := resp.Location()

		if err != nil {
			CheckForErrors(ErrorParams{Err: err, CallerNum: 1})
		} else {
			HttpPutRequestRedirect(u.String(), data)
		}
		resp.Body.Close()
	}
}
