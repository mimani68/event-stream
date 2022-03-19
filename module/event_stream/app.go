package event_stream

import (
	"fmt"
	"time"

	"zarinworld.ir/event/pkg/log_handler"
)

func EventHandlerModule(stateChannel chan interface{}) {
	go stopAppliction(stateChannel)
	// cronProxy(CRON_EVERY_2_SECONDS, func() {
	// 	log_handler.LoggerF("check undetermined authrities")
	// 	for _, address := range GetAddressList() {
	// 		updateUndeterminedAuthorities(address["addres"].(string))
	// 	}
	// })
	cronProxy(CRON_EVERY_10_SECONDS, func() {
		for _, network := range GetNetworkList() {
			updateCurrentBlock(network["network"].(string))
		}
		// Check new transactions
		for _, address := range GetAddressList() {
			newTransactionsList := updateNewTransactionOfAddress(address["address"].(string))
			for _, updatedTrx := range newTransactionsList {
				updatedTrx["type"] = "confirmed transactions"
				sendPostWebhook(updatedTrx)
			}
		}
		// dobule check status of confirm transactions
		for _, item := range getConfirmdTransactions() {
			updatedTrx := updateConfirmTransactions(item["txId"].(string))
			updatedTrx["type"] = "confirmed transactions"
			sendPostWebhook(updatedTrx)
		}
		cleanSystem()
	})
}

func cleanSystem() {
	log_handler.LoggerF("clean started")
}

func stopAppliction(stateChannel chan interface{}) {
	time.Sleep(7 * 24 * 3600 * time.Second)
	defer fmt.Printf("The application stoped at [%s]\n", time.Now().Format(time.RFC3339))
	stateChannel <- "done"
}
