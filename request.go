package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func post(url string, input interface{}, output interface{}) error {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonDataFromHttp, &output)
}
