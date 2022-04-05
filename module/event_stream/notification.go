package event_stream

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/db"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func sendPostWebhook(payload map[string]interface{}) (bool, error) {
	existInSendedList := false
	if payload["transaction_hash"] != nil {
		payload["trxId"] = utils.ToString(payload["transaction_hash"])
	} else if payload["hash"] != nil {
		payload["trxId"] = utils.ToString(payload["hash"])
	} else if payload["trxHash"] != nil {
		payload["trxId"] = utils.ToString(payload["trxHash"])
	} else if payload["transaction"] != nil {
		payload["trxId"] = utils.ToString(payload["transaction"])
	}

	if payload["address"] == nil || payload["address"] == "" || payload["address"] == "UNKNOWN" {
		log_handler.LoggerF("The message which its \"payload[address]\" is empty, will not send.")
		return false, errors.New("empty Address")
	}

	if payload["trxId"] == nil || payload["trxId"] == "" || payload["trxId"] == "UNKNOWN" {
		log_handler.LoggerF("The message which its \"payload[trxId]\" is empty, will not send.")
		return false, errors.New("empty trxId")
	}

	if payload["confirmCount"] == nil || payload["confirmCount"] == "" || payload["confirmCount"] == "UNKNOWN" {
		log_handler.LoggerF("The message which its \"payload[confirmCount]\" is empty, will not send.")
		return false, errors.New("empty confirmCount")
	}

	if payload["confirm"] == nil || payload["confirm"] == "" || payload["confirm"] == "UNKNOWN" {
		log_handler.LoggerF("The message which its \"payload[confirm]\" is empty, will not send.")
		return false, errors.New("empty confirm")
	}

	if payload["type"] == nil || payload["type"] == "" {
		log_handler.LoggerF("The message which its \"payload[type]\" is empty, will not send.")
		return false, errors.New("empty type")
	}

	preqeust, _ := json.Marshal(payload)
	if payload["trxId"] != nil && payload["type"] != nil {
		log_handler.LoggerF("[DEBUG][WEBHOOK] trxId=%s type=%s", payload["trxId"].(string), payload["type"].(string))
	}
	log_handler.LoggerF("[DEBUG][WEBHOOK] payload %s", string(preqeust))

	if len(db.GetAll(db.EVENTS)) == 0 {
		existInSendedList = false
	}

	// Check event list
	for _, event := range db.GetAll(db.EVENTS) {
		// Check duplicated request
		eventPayload := event["payload"].(map[string]interface{})
		isDuplicatedRequest := utils.ToString(payload["type"]) == utils.ToString(eventPayload["type"]) &&
			utils.ToString(payload["trxId"]) == utils.ToString(eventPayload["trxId"]) &&
			utils.ToString(payload["sendingStatus"]) == utils.ToString(eventPayload["sendingStatus"]) &&
			utils.ToString(payload["address"]) == utils.ToString(eventPayload["address"]) &&
			utils.ToString(payload["network"]) == utils.ToString(eventPayload["network"]) &&
			utils.ToString(payload["value"]) == utils.ToString(eventPayload["value"]) &&
			utils.ToString(payload["confirmCount"]) == utils.ToString(eventPayload["confirmCount"]) &&
			utils.ToString(payload["confirm"]) == utils.ToString(eventPayload["confirm"])
		// Check "confirmCount" less than config.Confirm_Count
		inConfirmRange := utils.ToInt(eventPayload["confirmCount"]) < config.ConfirmCount
		// Check for old events
		timeString, _ := time.Parse(time.RFC3339, event["time"].(string))
		notOldMessage := timeString.Add(config.AgeOfOldMessage).Unix() > time.Now().Unix()
		// sendedBefore := event["sendingStatus"].(bool) || payload["sendingStatus"] != nil || payload["status"] != nil
		if isDuplicatedRequest && notOldMessage && inConfirmRange {
			existInSendedList = true
			break
		}
	}

	// Check confirm more than "config.Confirm_Count"
	if utils.ToInt(payload["confirmCount"]) > config.ConfirmCount {
		return false, nil
	}

	if existInSendedList {
		log_handler.LoggerF("[DEBUG] Sending message address: %s and trxId: %s is unable because duplicated", payload["address"].(string), payload["trxId"].(string))
		return false, nil
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
	log_handler.LoggerF(`Confirmation message sent => network: "%s", address: "%s", trxId: "%s", status: "%s", webhookUrl: "%s"`, utils.ToString(payload["network"]), utils.ToString(payload["address"]), utils.ToString(payload["trxId"]), utils.ToString(payload["type"]), config.WebhookAddress)
	StoreEvent(payload, true, nil)
	return true, nil
}
