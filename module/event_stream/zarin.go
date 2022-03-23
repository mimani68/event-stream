package event_stream

import (
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/blockchair"
	"zarinworld.ir/event/pkg/db"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/tatum"
	"zarinworld.ir/event/pkg/utils"
)

func getconfirmTransactions() []map[string]interface{} {
	return db.GetAll(db.TRANSACTIONS)
}

func checkConfirmationOfSingleTransaction(network string, trxID string) map[string]interface{} {
	log_handler.LoggerF("Update trx %s that hash confirm > 0 on %s network", trxID, network)
	trx := tatum.GetTrxDetails(network, trxID)
	currentBlock := getCurrentBlock(network)
	trx["confirmCount"] = utils.ToString(math.Abs(float64(currentBlock - utils.ToInt(trx["blockNumber"]))))
	trx["confirm"] = true
	db.Store(db.TRANSACTIONS, trx)
	return trx
}

func getNewTransactionsOfAddress() []map[string]interface{} {
	return db.GetAll(db.NEW_TRANSACTIONS)
}

func updateNewTransactionOfAddress(network string, address string) []map[string]interface{} {
	log_handler.LoggerF("Checking new trx of address %s in network %s", address, network)
	tmp := []map[string]interface{}{}
	for _, transaction := range blockchair.GetAddressHistory(network, address) {
		if transaction["block_id"] == -1 {
			tmp = append(tmp, transaction)
			db.Store(db.NEW_TRANSACTIONS, transaction)
		} else {
			tmp = append(tmp, transaction)
			db.Store(db.TRANSACTIONS, transaction)
		}
	}
	return tmp
}

func GetUndeterminedAuthorities() []map[string]interface{} {
	return db.GetAll(db.AUTHORITIES)
}

func updateCurrentBlock(network string) {
	number := tatum.GetCurrentBlock(network)
	blockNumberObject := map[string]interface{}{
		network: number,
	}
	msg := fmt.Sprintf("Current block number of %s is %d and db update", network, number)
	log_handler.LoggerF(msg)
	db.Store(db.BLOCKNUMBER, blockNumberObject)
}

func getCurrentBlock(network string) int {
	number := 0
	networkList := db.GetAll(db.BLOCKNUMBER)
	for _, net := range networkList {
		number = utils.ToInt(net[network])
	}
	return number
}

func SetNewNetwork(network string) {
	log_handler.LoggerF("Network %s added", network)
	networkObject := map[string]interface{}{
		"network": network,
	}
	db.Store(db.NETWORK, networkObject)
}

func GetNetworkList() []map[string]interface{} {
	return db.GetAll(db.NETWORK)
}

func SetNewAddress(network string, address string) {
	log_handler.LoggerF("Address %s in %s network added", address, network)
	addressObject := map[string]interface{}{
		"network": network,
		"address": address,
	}
	db.Store(db.ADDRESS, addressObject)
}

func GetAddressList() []map[string]interface{} {
	return db.GetAll(db.ADDRESS)
}

func StoreEvent(payload map[string]interface{}, sendingStatus bool, failureObject error) {
	event := map[string]interface{}{
		"id":             uuid.New(),
		"type":           utils.ToString(payload["type"]),
		"payload":        payload,
		"url":            config.WebhookAddress,
		"time":           time.Now().Format(time.RFC3339),
		"sendingStatus":  sendingStatus,
		"failureDetails": failureObject,
	}
	db.Store(db.EVENTS, event)
}
