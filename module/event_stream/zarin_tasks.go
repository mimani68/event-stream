package event_stream

import (
	"fmt"
	"time"

	"zarinworld.ir/event/pkg/blockchair"
	"zarinworld.ir/event/pkg/db"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/tatum"
)

func updateConfirmTransactions(trxID string) map[string]interface{} {
	fmt.Println("send notification of confirmed > 0 authorities")
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
	// db.Store(db.NEW_TRANSACTIONS, trx)
	// fmt.Println("Check status of transaction")
	// return trx
	trx := tatum.GetTrxDetails(trxID)
	db.Store(db.TRANSACTIONS, trx)
	return trx
}

func updateNewTransactionOfAddress(address string) []map[string]interface{} {
	log_handler.LoggerF("Checking new trx of address %s", address)
	tmp := []map[string]interface{}{}
	for _, transaction := range blockchair.GetAddressHistory(address) {
		// if transaction["confirm"] == -1
		tmp = append(tmp, transaction)
		db.Store(db.NEW_TRANSACTIONS, transaction)
	}
	return tmp
}

// func updateUndeterminedAuthorities(address string) {
// 	log_handler.LoggerF("Checking all Undetermined authorities form Zarin BAAS")
// 	for _, authority := range zwbaas.GetAuthorities(address) {
// 		db.Store(db.AUTHORITIES, authority)
// 	}
// }

func GetUndeterminedAuthorities() []map[string]interface{} {
	return db.GetAll(db.AUTHORITIES)
}

func updateConfirmdTransactions(newTransactionObject map[string]interface{}) {
	db.Store(db.TRANSACTIONS, newTransactionObject)
}

func getConfirmdTransactions() []map[string]interface{} {
	return db.GetAll(db.TRANSACTIONS)
}

func updateCurrentBlock(network string) {
	// call network tatum/blockchair
	// FIXME: change from mock to real number
	fakeNumber := time.Now().Unix() / 100
	blockNumber := map[string]interface{}{
		network: fakeNumber + 1,
	}
	db.Store(db.BLOCKNUMBER, blockNumber)
	msg := fmt.Sprintf("Current block number of %s is %d and db update", network, blockNumber[network])
	log_handler.LoggerF(msg)
}

func getCurrentBlock(network string) int {
	number := 0
	networkList := db.GetAll(db.BLOCKNUMBER)
	for _, net := range networkList {
		number = net[network].(int)
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
