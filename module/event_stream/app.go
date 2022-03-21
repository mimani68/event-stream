package event_stream

import (
	"fmt"
	"time"

	"gorm.io/gorm/utils"
	"zarinworld.ir/event/pkg/delay"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func EventHandlerModule(stateChannel chan interface{}) {
	go stopAppliction(stateChannel)
	cronProxy(CRON_EVERY_2_SECONDS, func() {
		// Get latest block number
		for _, network := range GetNetworkList() {
			go updateCurrentBlock(network["network"].(string))
		}
	})
	cronProxy(CRON_EVERY_5_SECONDS, func() {
		for _, network := range GetNetworkList() {
			// Check new transactions
			for _, address := range GetAddressList() {
				newTransactionsList := updateNewTransactionOfAddress(network["network"].(string), address["address"].(string))
				for _, updatedTrx := range newTransactionsList {
					updatedTrx["type"] = "new transaction detected"
					updatedTrx["confirmedCount"] = 0
					updatedTrx["confirmed"] = false
					go sendPostWebhook(updatedTrx)
				}
			}
			delay.SetSyncDelay(2)
		}
	})
	cronProxy(CRON_EVERY_10_SECONDS, func() {
		for _, network := range GetNetworkList() {
			// Check status of new transactions and update them
			for _, newItem := range getNewTransactionsOfAddress() {
				updatedTrx := checkConfirmationOfSingleTransaction(utils.ToString(network["network"]), utils.ToString(newItem["trxHash"]))
				updatedTrx["type"] = "confirmed transactions"
				go sendPostWebhook(updatedTrx)
				// FIXME: remove from NEW_TRANSACTIONS
			}
			delay.SetSyncDelay(2)
			// Dobule check status of confirmed transactions for confirmedCount> 1
			for _, newItem := range getConfirmedTransactions() {
				updatedTrx := checkConfirmationOfSingleTransaction(network["network"].(string), newItem["trxHash"].(string))
				// FIXME: updatedTrx["confirmedCount"] > 5 ==> remove from TRANSACTIONS
				updatedTrx["type"] = "confirmed transactions"
				go sendPostWebhook(updatedTrx)
			}
			delay.SetSyncDelay(2)
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
