package event_stream

import (
	"fmt"

	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/delay"
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
	fmt.Println(getNewTransactions())
	fmt.Println(getconfirmTransactions())
	for _, newItem := range getNewTransactions() {
		var updatedTrx map[string]interface{}
		switch newItem["network"] {
		case config.BITCOIN:
			updatedTrx = checkConfirmationOfSingleTransaction(utils.ToString(newItem["network"]), utils.ToString(newItem["hash"]))
		case config.ETHEREUM:
			updatedTrx = checkConfirmationOfSingleTransaction(utils.ToString(newItem["network"]), utils.ToString(newItem["transaction_hash"]))
		}
		updatedTrx["type"] = "confirm transactions"
		go sendPostWebhook(updatedTrx)
		// FIXME: remove from NEW_TRANSACTIONS
		delay.SetSyncDelay(1)
	}
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
