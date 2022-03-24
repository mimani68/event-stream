package event_stream

import (
	"fmt"
	"time"

	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/delay"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func EventHandlerModuleDev(stateChannel chan string) {
	for _, network := range GetNetworkList() {
		// Get latest block number
		updateCurrentBlock(utils.ToString(network["network"]))
		delay.SetSyncDelay(1)
	}
	// Check new transactions
	for _, address := range GetAddressList() {
		newTransactionsList := updateNewTransactionOfAddress(utils.ToString(address["network"]), utils.ToString(address["address"]))
		for _, updatedTrx := range newTransactionsList {
			updatedTrx["type"] = "new transaction detected"
			go sendPostWebhook(updatedTrx)
			delay.SetSyncDelay(2)
		}
	}
	// for _, newItem := range getNewTransactionsOfAddress() {
	// 	var updatedTrx map[string]interface{}
	// 	switch network["network"] {
	// 	case config.BITCOIN:
	// 		updatedTrx = checkConfirmationOfSingleTransaction(utils.ToString(network["network"]), utils.ToString(newItem["trxHash"]))
	// 	case config.ETHEREUM:
	// 		updatedTrx = checkConfirmationOfSingleTransaction(utils.ToString(network["network"]), utils.ToString(newItem["transaction_hash"]))
	// 	}
	// 	updatedTrx["type"] = "confirm transactions"
	// 	go sendPostWebhook(updatedTrx)
	// 	// FIXME: remove from NEW_TRANSACTIONS
	// 	delay.SetSyncDelay(1)
	// }
	fmt.Println(getNewTransactions())
	fmt.Println(getconfirmTransactions())
	// updatedTrx := checkConfirmationOfSingleTransaction(config.ETHEREUM, "0x13c28d5e3a0b7a21a4b516e7d1b4f9b22f6cadeeecc93bb5b490cd99ce6f3f2b")
	// fmt.Println(updatedTrx)

	// StoreEvent(map[string]interface{}{
	// 	"type":         "sample",
	// 	"confirmCount": 2,
	// 	"time":         time.Now().Format(time.RFC3339),
	// }, true, nil)
	// StoreEvent(map[string]interface{}{
	// 	"type":         "sample",
	// 	"confirmCount": 3,
	// 	"time":         time.Now().Format(time.RFC3339),
	// }, true, nil)
	// sendPostWebhook(map[string]interface{}{
	// 	"type":         "sample",
	// 	"confirmCount": 4,
	// 	"time":         time.Now().Format(time.RFC3339),
	// })
	// sendPostWebhook(map[string]interface{}{
	// 	"type":         "sample",
	// 	"confirmCount": 4,
	// 	"time":         time.Now().Format(time.RFC3339),
	// })
	// sendPostWebhook(map[string]interface{}{
	// 	"type":         "sample",
	// 	"confirmCount": 4,
	// 	"time":         time.Now().Format(time.RFC3339),
	// })
	// sendPostWebhook(map[string]interface{}{
	// 	"type":         "sample",
	// 	"confirmCount": 4,
	// 	"time":         time.Now().Format(time.RFC3339),
	// })
	// a := db.GetAll(db.EVENTS)
	// fmt.Println(a)
}

func EventHandlerModule(stateChannel chan string) {
	cronProxy(CRON_EVERY_5_SECONDS, func() {
		// Get latest block number
		for _, network := range GetNetworkList() {
			updateCurrentBlock(utils.ToString(network["network"]))
			delay.SetSyncDelay(1)
		}
	})
	cronProxy(CRON_EVERY_10_SECONDS, func() {
		for _, network := range GetNetworkList() {
			// Check new transactions
			for _, address := range GetAddressList() {
				newTransactionsList := updateNewTransactionOfAddress(utils.ToString(network["network"]), address["address"].(string))
				for _, updatedTrx := range newTransactionsList {
					updatedTrx["type"] = "new transaction detected"
					go sendPostWebhook(updatedTrx)
					delay.SetSyncDelay(2)
				}
			}
		}
	})
	cronProxy(CRON_EVERY_15_SECONDS, func() {
		for _, network := range GetNetworkList() {
			// Check status of new transactions and update them
			for _, newItem := range getNewTransactions() {
				var updatedTrx map[string]interface{}
				switch network["network"] {
				case config.BITCOIN:
					updatedTrx = checkConfirmationOfSingleTransaction(utils.ToString(network["network"]), utils.ToString(newItem["trxHash"]))
				case config.ETHEREUM:
					updatedTrx = checkConfirmationOfSingleTransaction(utils.ToString(network["network"]), utils.ToString(newItem["transaction_hash"]))
				}
				updatedTrx["type"] = "confirm transactions"
				go sendPostWebhook(updatedTrx)
				// FIXME: remove from NEW_TRANSACTIONS
				delay.SetSyncDelay(1)
			}
			// Dobule check status of confirm transactions for confirmCount> 1
			for _, newItem := range getconfirmTransactions() {
				updatedTrx := checkConfirmationOfSingleTransaction(network["network"].(string), newItem["trxHash"].(string))
				// FIXME: updatedTrx["confirmCount"] > 5 ==> remove from TRANSACTIONS
				updatedTrx["type"] = "confirm transactions"
				go sendPostWebhook(updatedTrx)
				delay.SetSyncDelay(1)
			}
		}
	})
	cronProxy(CRON_EVERY_30_MINUTES, func() {
		cleanSystem()
	})
	cronProxy(CRON_AT_4_OCLOCK, func() {
		stopAppliction(stateChannel)
	})
}

func cleanSystem() {
	log_handler.LoggerF("Cleaning start")
}

func stopAppliction(st chan string) {
	st <- "DONE"
	defer fmt.Printf("The application stoped at [%s]\n", time.Now().Format(time.RFC3339))
}
