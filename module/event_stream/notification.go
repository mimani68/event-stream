package event_stream

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/db"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func sendPostWebhook(payload map[string]interface{}) (bool, error) {
	preqeust, _ := json.Marshal(payload)
	log_handler.LoggerF("[DEBUG][WEBHOOK] Sending post webhook request with payload %s", string(preqeust))
	// Check confirm more than "config.Confirm_Count"
	if !config.Simulate_new_request && utils.ToInt(payload["confirmCount"]) > config.ConfirmCount {
		return false, nil
	}
	for _, event := range db.GetAll(db.EVENTS) {
		eventPayload := event["payload"].(map[string]interface{})
		timeString, _ := time.Parse(time.RFC3339, event["time"].(string))
		isDuplicatedRequest := utils.ToString(payload["id"]) == utils.ToString(eventPayload["id"]) &&
			utils.ToString(payload["type"]) == utils.ToString(eventPayload["type"]) &&
			utils.ToString(payload["hash"]) == utils.ToString(eventPayload["hash"]) &&
			utils.ToString(payload["trxHash"]) == utils.ToString(eventPayload["trxHash"]) &&
			utils.ToString(payload["transaction_hash"]) == utils.ToString(eventPayload["transaction_hash"]) &&
			utils.ToString(payload["trxId"]) == utils.ToString(eventPayload["trxId"]) &&
			utils.ToString(payload["address"]) == utils.ToString(eventPayload["address"]) &&
			utils.ToString(payload["network"]) == utils.ToString(eventPayload["network"]) &&
			utils.ToString(payload["value"]) == utils.ToString(eventPayload["value"]) &&
			utils.ToString(payload["confirmCount"]) == utils.ToString(eventPayload["confirmCount"])
		// Check duplicated request
		if !config.Simulate_new_request && isDuplicatedRequest {
			return false, nil
		}
		// Check for old events
		if !config.Simulate_new_request && utils.ToInt(timeString.Add(24*time.Hour).Unix()) < int(time.Now().Unix()) {
			return false, nil
		}
		// Check "confirmCount" less than config.Confirm_Count
		if !config.Simulate_new_request && utils.ToInt(eventPayload["confirmCount"]) > config.ConfirmCount {
			return false, nil
		}
	}
	if payload["transaction_hash"] != nil {
		payload["trxId"] = utils.ToString(payload["transaction_hash"])
	} else if payload["hash"] != nil {
		payload["trxId"] = utils.ToString(payload["hash"])
	}
	postBody, _ := json.Marshal(payload)
	requestBody := bytes.NewBuffer(postBody)
	jsonPayload, _ := json.Marshal(payload)
	_, err := http.Post(config.WebhookAddress, "application/json", requestBody)
	if err != nil {
		log_handler.LoggerF("Error in sending webhook to %s%s%s", log_handler.ColorRed, config.WebhookAddress, log_handler.ColorReset)
		StoreEvent(payload, false, err)
		return false, err
	}
	log_handler.LoggerF("[DEBUG][WEBHOOK] webhook sent by payload %s", string(jsonPayload))
	log_handler.LoggerF("Message sent as \"%s%s%s\" status to \"%s\"", log_handler.ColorGreen, payload["type"].(string), log_handler.ColorReset, config.WebhookAddress)
	StoreEvent(payload, true, nil)
	return true, nil
}
