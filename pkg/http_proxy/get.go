package http_proxy

import (
	"io/ioutil"
	"net/http"

	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func Get(url string, header map[string]string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log_handler.LoggerF("Send \"GET\" request to %s is unable", url)
		return "", err
	}
	if header["x-api-token"] != "" {
		req.Header.Add("x-api-token", utils.ToString(header["x-api-token"]))
	}
	response, err := client.Do(req)
	// responseByte, err := http.Get(url)
	if err != nil {
		log_handler.LoggerF("Send \"GET\" request to %s is unable", url)
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log_handler.LoggerF(err.Error())
		return "", err
	}

	return string(body), err
}
