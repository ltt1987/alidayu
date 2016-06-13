package alidayu

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var UseHTTP = false

func DoPost(m map[string]string) (success bool, response string) {
	if AppKey == "" || AppSecret == "" {
		return false, "AppKey or AppSecret is requierd!"
	}

	body, size := getRequestBody(m)
	client := &http.Client{}
	var req *http.Request
	var err error
	if !UseHTTP {
		req, err = http.NewRequest("POST", URL_HTTPS, body)
	} else {
		req, err = http.NewRequest("POST", URL_HTTP, body)
	}
	if err != nil {
		return false, err.Error()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.ContentLength = size

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		response = err.Error()
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	response = string(data)
	if strings.Contains(response, "success") {
		return true, response
	} else {
		return false, response
	}
}
