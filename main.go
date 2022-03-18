package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"zarinworld.ir/event/pkg/db"
)

const (
	CRON_EVERY_SECONDS    = "* * * * * *"
	CRON_EVERY_5_SECONDS  = "*/5 * * * * *"
	CRON_EVERY_15_SECONDS = "*/15 * * * * *"
	CRON_EVERY_30_SECONDS = "*/30 * * * * *"
	CRON_EVERY_60_SECONDS = "* * * * *"
)

func updateConfirmTransactions(trxID string) {
	// call blockchair
	//    if blockNumber === -1 => send confirm: false
	//    if blockNumber > -1 =>
	//			call currentBlock number => x
	//			send { confirm: true, confirmCount: x }
	trx := map[string]interface{}{
		"id":          "c4e064b8-a6b2-11ec-a14d-9f5202368044",
		"value":       "0",
		"blockNumber": "-1",
		"trxHash":     "2sd3srjyj2wg1sfn1y3kl13a1f3fh1k543s2g1bs3jhljwj",
		"expireIn":    time.Now().Add(5 * time.Minute),
	}
	db.Store(db.NEW_TRANSACTIONS, trx)
	fmt.Println("Check status of transaction")
}

func checkNewTransactionOfAddress(address string) {
	// call blockchair
	//    if blockNumber === -1 => send confirm: false
	//    if blockNumber > -1 =>
	//			call currentBlock number => x
	//			send { confirm: true, confirmCount: x }
	trx := make(map[string]string)
	trx["id"] = "c4e064b8-a6b2-11ec-a14d-9f5202368044"
	trx["value"] = "0"
	trx["blockNumber"] = "-1"
	trx["trxHash"] = "2sd3srjyj2wg1sfn1y3kl13a1f3fh1k543s2g1bs3jhljwj"
	trx["expireIn"] = time.Now().Add(5 * time.Minute).String()
	db.Store(db.NEW_TRANSACTIONS, trx)
	fmt.Println("Check status of transaction")
}

func checkUndeterminedAuthorities() {
	// call baas/authorities/
	// filter { incoming_value:0 , status: "ACTIVE", expire >= now }
	// authList []
	authority := make(map[string]string)
	authority["id"] = "0e61a584-a6a4-11ec-849e-d32cf9f774d5"
	authority["value"] = "0"
	authority["expireIn"] = "2022-10-01T00:00:00"
	db.Store(db.AUTHORITIES, authority)
}

func getUndeterminedAuthorities() []interface{} {
	return db.GetAll(db.AUTHORITIES)
}

func currentBlock(network string) {
	// call network tatum/blockchair
	// FIXME: change from mock to real number
	blockNumber := map[string]string{
		"dd3a0b82-a6b1-11ec-ada3-9f36a7680e48": "135510",
	}
	db.Store(db.BLOCKNUMBER, blockNumber)
}

func getCurrentBlock(network string) int {
	number := 0
	networkList := db.GetAll(db.BLOCKNUMBER)
	for index, _ := range networkList {
		number = networkList[index].(int)
	}
	return number
}

func setNewNetwork(network string) {
	networkObject := map[string]interface{}{
		"network": network,
	}
	db.Store(db.NETWORK, networkObject)
}

func gettNetworkList() []interface{} {
	return db.GetAll(db.NETWORK)
}

func cleanSystem() {
	fmt.Println("clean started")
}

func cronProxy(cronTimeString string, cb func()) (bool, error) {
	c := cron.New(cron.WithSeconds())
	cronId, err := c.AddFunc(cronTimeString, cb)
	if err != nil {
		return false, err
	}
	fmt.Println(cronId)
	c.Start()
	return true, nil
}

func EventHandlerApplication() {
	cronProxy(CRON_EVERY_SECONDS, func() {
		fmt.Println("check undetermined authrities")
		checkUndeterminedAuthorities()
	})
	cronProxy(CRON_EVERY_5_SECONDS, func() {
		fmt.Println("check current block number")
		fmt.Println("send notification of confirmed > 0 authorities")
		cleanSystem()
	})
	cronProxy(CRON_EVERY_15_SECONDS, func() {
		fmt.Println("check check new authorities from blockchair")
		fmt.Println("check tatum confirmed authorities")
		fmt.Println("clean system and expire datas")
	})
	time.Sleep(10 * 365 * 24 * 3600 * time.Second)
}

func main() {
	setNewNetwork("ethereum")
	setNewNetwork("bitcoin")
	EventHandlerApplication()
}
