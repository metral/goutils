package goutils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
)

func httpGetRequest(url string) []byte {
	resp, err := http.Get(url)
	checkForErrors(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkForErrors(err)

	return body
}

func httpPutRequest(urlStr string, data []byte) *http.Response {
	var req *http.Request

	req, _ = http.NewRequest("PUT", urlStr, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	checkForErrors(err)

	defer resp.Body.Close()

	return resp
}

func httpPutRequestRedirect(urlStr string, data string) {
	var req *http.Request
	req, _ = http.NewRequest("PUT", urlStr, bytes.NewBufferString(data))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	client := &http.Client{}
	resp, err := client.Do(req)
	checkForErrors(err)

	if resp.StatusCode == http.StatusTemporaryRedirect {
		u, err := resp.Location()

		if err != nil {
			checkForErrors(err)
		} else {
			httpPutRequestRedirect(u.String(), data)
		}
		resp.Body.Close()
	}
}
