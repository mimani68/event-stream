package event_stream

import "zarinworld.ir/event/pkg/log_handler"

func sendPostWebhook(payload map[string]interface{}) {
	log_handler.LoggerF("Message sent ", payload["type"].(string))
}
