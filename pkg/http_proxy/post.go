package http_proxy

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"zarinworld.ir/event/pkg/log_handler"
)

func Post(url string, payload []map[string]interface{}) (string, error) {
	postBody, _ := json.Marshal(payload)
	requestBody := bytes.NewBuffer(postBody)
	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		log_handler.LoggerF("Send \"POST\" request to %s is unable", url)
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log_handler.LoggerF(err.Error())
		return "", err
	}
	return string(body), nil
}
