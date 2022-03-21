package event_stream

import (
	"fmt"

	"zarinworld.ir/event/pkg/blockchair"
	"zarinworld.ir/event/pkg/db"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/tatum"
	"zarinworld.ir/event/pkg/utils"
	"zarinworld.ir/event/pkg/zwbaas"
)

func getConfirmedTransactions() []map[string]interface{} {
	return db.GetAll(db.TRANSACTIONS)
}

func checkConfirmationOfSingleTransaction(network string, trxID string) map[string]interface{} {
	// log_handler.LoggerF("Update confirmed > 0 authorities")
	// // call blockchair
	// //    if blockNumber === -1 => send confirm: false
	// //    if blockNumber > -1 =>
	// //			call currentBlock number => x
	// //			send { confirm: true, confirmCount: x }
	// trx := map[string]interface{}{
	// 	"id":          "c4e064b8-a6b2-11ec-a14d-9f5202368044",
	// 	"value":       "0",
	// 	"blockNumber": "-1",
	// 	"trxHash":     "2sd3srjyj2wg1sfn1y3kl13a1f3fh1k543s2g1bs3jhljwj",
	// 	"expireIn":    time.Now().Add(5 * time.Minute),
	// }

	// trxList := get()
	// for _, item := range trxList {
	// 	if item["blockNumber"]
	// }
	// db.Store(db.NEW_TRANSACTIONS, trx)
	// fmt.Println("Check status of transaction")
	// return trx

	trx := tatum.GetTrxDetails(network, trxID)
	currentBlock := getCurrentBlock(network)
	trx["confirmedCount"] = currentBlock
	trx["confirmed"] = true
	db.Store(db.TRANSACTIONS, trx)
	return trx
}

func getNewTransactionsOfAddress() []map[string]interface{} {
	// log_handler.LoggerF("Get new trx of address %s", address)
	return db.GetAll(db.NEW_TRANSACTIONS)
}

func updateNewTransactionOfAddress(network string, address string) []map[string]interface{} {
	log_handler.LoggerF("Checking new trx of address %s in network %s", address, network)
	tmp := []map[string]interface{}{}
	for _, transaction := range blockchair.GetAddressHistory(network, address) {
		if transaction["block_id"] == -1 {
			tmp = append(tmp, transaction)
			db.Store(db.NEW_TRANSACTIONS, transaction)
		}
	}
	return tmp
}

func GetUndeterminedAuthorities() []map[string]interface{} {
	return db.GetAll(db.AUTHORITIES)
}

func updateUndeterminedAuthorities(address string) {
	log_handler.LoggerF("Checking all Undetermined authorities form Zarin BAAS")
	for _, authority := range zwbaas.GetAuthorities(address) {
		db.Store(db.AUTHORITIES, authority)
	}
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
