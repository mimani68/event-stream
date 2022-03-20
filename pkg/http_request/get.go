package http_request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"zarinworld.ir/event/pkg/log_handler"
)

func Get(url string) ([]map[string]interface{}, error) {
	response := []map[string]interface{}{}
	responseByte, err := http.Get(url)
	if err != nil {
		log_handler.LoggerF("Send \"GET\" request to %s is unable", url)
		return response, err
	}
	body, err := ioutil.ReadAll(responseByte.Body)
	if err != nil {
		log_handler.LoggerF(err.Error())
		return response, err
	}
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		log_handler.LoggerF(err.Error())
		return response, err
	}
	return response, nil
}
