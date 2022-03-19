package event_stream

import (
	"bytes"
	"encoding/json"
	"net/http"

	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/log_handler"
)

func sendPostWebhook(payload map[string]interface{}) (bool, error) {
	postBody, _ := json.Marshal(payload)
	requestBody := bytes.NewBuffer(postBody)
	_, err := http.Post(config.WebhookAddress, "application/json", requestBody)
	if err != nil {
		log_handler.LoggerF("Error in sending webhook to %s%s%s", log_handler.ColorRed, config.WebhookAddress, log_handler.ColorReset)
		return false, err
	}
	log_handler.LoggerF("Message sent with type: %s%s%s", log_handler.ColorGreen, payload["type"].(string), log_handler.ColorReset)
	return true, nil
}
