package event_stream

import (
	"fmt"
	"time"

	"zarinworld.ir/event/pkg/log_handler"
)

func EventHandlerModule(stateChannel chan interface{}) {
	go stopAppliction(stateChannel)
	cronProxy(CRON_EVERY_2_SECONDS, func() {
		// Get latest block number
		for _, network := range GetNetworkList() {
			updateCurrentBlock(network["network"].(string))
		}
	})
	cronProxy(CRON_EVERY_5_SECONDS, func() {
		// Check new transactions
		for _, address := range GetAddressList() {
			newTransactionsList := updateNewTransactionOfAddress(address["address"].(string))
			for _, updatedTrx := range newTransactionsList {
				updatedTrx["type"] = "new transaction detected"
				updatedTrx["confirmed"] = false
				sendPostWebhook(updatedTrx)
			}
		}
	})
	cronProxy(CRON_EVERY_10_SECONDS, func() {
		// Check status of new transactions and update them
		for _, newItem := range getNewTransactionsOfAddress() {
			updatedTrx := checkConfirmationOfSingleTransaction(newItem["trxHash"].(string))
			updatedTrx["type"] = "confirmed transactions"
			updatedTrx["confirmed"] = true
			updatedTrx["confirmedCount"] = 1
			sendPostWebhook(updatedTrx)
		}
		// Dobule check status of confirmed transactions for confirmedCount> 1
		for _, newItem := range getConfirmedTransactions() {
			updatedTrx := checkConfirmationOfSingleTransaction(newItem["trxHash"].(string))
			updatedTrx["type"] = "confirmed transactions"
			updatedTrx["confirmed"] = true
			if updatedTrx["confirmedCount"] == "" {
				updatedTrx["confirmedCount"] = 0
			} else {
				updatedTrx["confirmedCount"] = updatedTrx["confirmedCount"].(int) + 1
			}
			sendPostWebhook(updatedTrx)
		}
		cleanSystem()
	})
}

func cleanSystem() {
	log_handler.LoggerF("Cleaning start")
}

func stopAppliction(stateChannel chan interface{}) {
	time.Sleep(7 * 24 * 3600 * time.Second)
	defer fmt.Printf("The application stoped at [%s]\n", time.Now().Format(time.RFC3339))
	stateChannel <- "done"
}
