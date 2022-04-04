package event_stream

import (
	"fmt"
	"os"
	"time"

	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/delay"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func EventHandlerModule(stateChannel chan string) {
	cronProxy(CRON_EVERY_5_SECONDS, func() {
		// Get latest block number
		for _, network := range GetNetworkList() {
			UpdateCurrentBlock(utils.ToString(network["network"]))
			delay.SetSyncDelay(1)
		}
	})
	cronProxy(CRON_EVERY_15_SECONDS, func() {
		// Check new transactions
		for _, address := range GetAddressList() {
			newTransactionsList := UpdateNewTransactionOfAddress(utils.ToString(address["network"]), utils.ToString(address["address"]))
			for _, updatedTrx := range newTransactionsList {
				updatedTrx["type"] = "new transaction detected"
				go sendPostWebhook(updatedTrx)
				delay.SetSyncDelay(3)
			}
		}
		// Check status of new transactions and update them
		for _, newItem := range GetNewTransactions() {
			updatedTrx := map[string]interface{}{}
			switch newItem["network"] {
			case config.BITCOIN:
				updatedTrx = CheckConfirmationOfSingleTransaction(utils.ToString(newItem["network"]), utils.ToString(newItem["hash"]))
			case config.ETHEREUM:
				updatedTrx = CheckConfirmationOfSingleTransaction(utils.ToString(newItem["network"]), utils.ToString(newItem["transaction_hash"]))
			}
			if updatedTrx != nil {
				updatedTrx["address"] = newItem["address"]
				updatedTrx["type"] = "confirm transactions"
				go sendPostWebhook(updatedTrx)
				// FIXME: remove from NEW_TRANSACTIONS
				delay.SetSyncDelay(3)
			}
		}
		// Dobule check status of confirm transactions for confirmCount> 1
		for _, newItem := range GetConfirmTransactions() {
			updatedTrx := map[string]interface{}{}
			switch newItem["network"] {
			case config.BITCOIN:
				updatedTrx = CheckConfirmationOfSingleTransaction(utils.ToString(newItem["network"]), utils.ToString(newItem["hash"]))
			case config.ETHEREUM:
				updatedTrx = CheckConfirmationOfSingleTransaction(utils.ToString(newItem["network"]), utils.ToString(newItem["transaction_hash"]))
			}
			// FIXME: updatedTrx["confirmCount"] > 5 ==> remove from TRANSACTIONS
			updatedTrx["address"] = newItem["address"]
			updatedTrx["type"] = "confirm transactions"
			go sendPostWebhook(updatedTrx)
			delay.SetSyncDelay(5)
		}
	})
	cronProxy(CRON_EVERY_30_MINUTES, func() {
		cleanSystem()
	})
	cronProxy(CRON_EVERY_6_HOURS, func() {
		stopAppliction(stateChannel)
	})
}

func cleanSystem() {
	log_handler.LoggerF("Cleaning start")
}

func stopAppliction(st chan string) {
	fmt.Println("")
	fmt.Println("")
	log_handler.LoggerF("The Application are going to stop, now is [%s]\n", time.Now().Format(time.RFC3339))
	log_handler.LoggerF("BYE")
	fmt.Println("")
	fmt.Println("")
	st <- "DONE"
	defer os.Exit(0)
}
