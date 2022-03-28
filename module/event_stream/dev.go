package event_stream

import (
	"math/rand"
	"time"
)

func EventHandlerModuleDev(stateChannel chan string) {

	// for _, address := range GetAddressList() {
	// 	newTransactionsList := updateNewTransactionOfAddress(utils.ToString(address["network"]), utils.ToString(address["address"]))
	// 	for _, updatedTrx := range newTransactionsList {
	// 		updatedTrx["type"] = "New transaction detected"
	// 		go sendPostWebhook(updatedTrx)
	// 		delay.SetSyncDelay(2)
	// 	}
	// }

	// for _, network := range GetNetworkList() {
	// 	// Get latest block number
	// 	updateCurrentBlock(utils.ToString(network["network"]))
	// 	delay.SetSyncDelay(1)
	// }
	// // Check new transactions
	// for _, address := range GetAddressList() {
	// 	newTransactionsList := updateNewTransactionOfAddress(utils.ToString(address["network"]), utils.ToString(address["address"]))
	// 	for _, updatedTrx := range newTransactionsList {
	// 		updatedTrx["type"] = "new transaction detected"
	// 		go sendPostWebhook(updatedTrx)
	// 		delay.SetSyncDelay(2)
	// 	}
	// }
	// fmt.Println(getNewTransactions())
	// fmt.Println(getconfirmTransactions())
	// for _, newItem := range getNewTransactions() {
	// 	var updatedTrx map[string]interface{}
	// 	switch newItem["network"] {
	// 	case config.BITCOIN:
	// 		updatedTrx = checkConfirmationOfSingleTransaction(utils.ToString(newItem["network"]), utils.ToString(newItem["hash"]))
	// 	case config.ETHEREUM:
	// 		updatedTrx = checkConfirmationOfSingleTransaction(utils.ToString(newItem["network"]), utils.ToString(newItem["transaction_hash"]))
	// 	}
	// 	updatedTrx["type"] = "confirm transactions"
	// 	go sendPostWebhook(updatedTrx)
	// 	// FIXME: remove from NEW_TRANSACTIONS
	// 	delay.SetSyncDelay(1)
	// }

	// updatedTrx := checkConfirmationOfSingleTransaction(config.ETHEREUM, "0x13c28d5e3a0b7a21a4b516e7d1b4f9b22f6cadeeecc93bb5b490cd99ce6f3f2b")
	// fmt.Println(updatedTrx)

	eventList := []map[string]interface{}{
		{
			"balance_change": rand.Intn(3),
			"block_id":       rand.Float64(),
			"trxId":          "e7e027e80d036b4faa2ec5a8e2d8ae584df9e3c566c407483add1edcdc06080f",
			"time":           time.Now().Format(time.RFC3339),
		},
		{
			"balance_change": rand.Intn(3),
			"block_id":       rand.Float64(),
			"transaction":    "e7e027e80d036b4faa2ec5a8e2d8ae584df9e3c566c407483add1edcdc06080f",
			"time":           time.Now().Format(time.RFC3339),
		},
		{
			"balance_change": rand.Intn(3),
			"block_id":       rand.Float64(),
			"hash":           "e7e027e80d036b4faa2ec5a8e2d8ae584df9e3c566c407483add1edcdc06080f",
			"time":           time.Now().Format(time.RFC3339),
		},
	}
	for _, v := range eventList {
		StoreEvent(v, true, nil)
	}
	p := eventList[0]
	p["type"] = "DONE"
	sendPostWebhook(p)
	sendPostWebhook(map[string]interface{}{
		"balance_change": rand.Intn(3),
		"block_id":       rand.Float64(),
		"hash":           "aloo",
		"confirm":        false,
		"confirmCount":   0,
		"type":           "OK",
		"time":           time.Now().Format(time.RFC3339),
	})
	time.Sleep(5 * time.Second)
	sendPostWebhook(map[string]interface{}{
		"balance_change": rand.Intn(3),
		"block_id":       rand.Float64(),
		"hash":           "aloo",
		"confirm":        true,
		"confirmCount":   1,
		"type":           "OK",
		"time":           time.Now().Format(time.RFC3339),
	})
	time.Sleep(5 * time.Second)
	sendPostWebhook(map[string]interface{}{
		"balance_change": rand.Intn(3),
		"block_id":       rand.Float64(),
		"hash":           "aloo",
		"confirm":        true,
		"confirmCount":   2,
		"type":           "OK",
		"time":           time.Now().Format(time.RFC3339),
	})
}
