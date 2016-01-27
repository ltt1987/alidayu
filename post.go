package alidayu

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func DoPost(m map[string]string) (success bool, response string) {
	if AppKey == "" || AppSecret == "" {
		return false, "AppKey or AppSecret is requierd!"
	}

	body, size := getRequestBody(m)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", URL, body)
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
