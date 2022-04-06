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
	cronProxy(CRON_EVERY_10_SECONDS, func() {
		// Get latest block number
		for _, network := range GetNetworkList() {
			UpdateCurrentBlock(utils.ToString(network["network"]))
			delay.SetSyncDelay(5)
		}
	})

	// ensuring all network data loaded
	delay.SetSyncDelay(30)

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
	})

	// Ensuring all new TRX downloaded
	delay.SetSyncDelay(10)

	cronProxy(CRON_EVERY_30_SECONDS, func() {

		dn := func(newItem map[string]interface{}) {
			updatedTrx := map[string]interface{}{}
			switch newItem["network"] {
			case config.BITCOIN:
				updatedTrx = CheckConfirmationOfSingleTransaction(utils.ToString(newItem["network"]), utils.ToString(newItem["hash"]))
			case config.ETHEREUM:
				updatedTrx = CheckConfirmationOfSingleTransaction(utils.ToString(newItem["network"]), utils.ToString(newItem["transaction_hash"]))
			}
			if updatedTrx != nil {
				updatedTrx["address"] = newItem["address"]
				updatedTrx["type"] = "confirm_transaction"
				go sendPostWebhook(updatedTrx)
				// FIXME: remove from NEW_TRANSACTIONS
				delay.SetSyncDelay(2)
			}
		}

		// Dobule check status of confirm transactions for confirmCount> 1
		for _, newItem := range GetConfirmTransactions() {
			dn(newItem)
			delay.SetSyncDelay(15)
		}

		// Check status of new transactions and update them
		for _, newItem := range GetNewTransactions() {
			dn(newItem)
			delay.SetSyncDelay(10)
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
	err := os.Remove(config.LogFilePath)
	if err != nil {
		log_handler.LoggerF("Unable to remove log file " + config.LogFilePath)
		log_handler.LoggerF(err.Error())
	}
	log_handler.LoggerF("Cleaning finished")
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
