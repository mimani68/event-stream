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

func checkStatusOfTransaction(address string) {
	// call blockchair
	//    if blockNumber === -1 => send confirm: false
	//    if blockNumber > -1 =>
	//			call currentBlock number => x
	//			send { confirm: true, confirmCount: x }
	fmt.Println("Check status of transaction")
}

func checkUndeterminedAuthorities() {
	// call baas/authorities/
	// filter { incoming_value:0 , status: "ACTIVE", expire >= now }
	// authList []
	title := "sampl auth id"
	m := make(map[string]string)
	m[title] = "auth goes here"
	db.Store(db.AUTHORITY, m)
}

func cleanAuthoriries() {
	fmt.Println("clean started")
}

func getUndeterminedAuthorities() []interface{} {
	return db.GetAll(db.AUTHORITY)
}

func currentBlock(network string) {
	// call network tatum/blockchair
	// FIXME: change from mock to real number
	m := make(map[string]int)
	m[network] = int(time.Now().Unix())
	db.CurrentBlockNumberList = append(db.CurrentBlockNumberList, m)
}

func getCurrentBlock(network string) int {
	number := 0
	for index, _ := range db.CurrentBlockNumberList {
		number = db.CurrentBlockNumberList[index].(int)
	}
	return number
}

func setNewNetwork(network string) {
	db.NetworkList = append(db.NetworkList, network)
}

func gettNetworkList() []string {
	return db.NetworkList
}

func EventHandlerApplication() {
	cronProxy(CRON_EVERY_SECONDS, func() {
		fmt.Println("check undetermined authrities")
		checkUndeterminedAuthorities()
	})
	cronProxy(CRON_EVERY_5_SECONDS, func() {
		fmt.Println("check current block number")
		fmt.Println("send notification of confirmed > 0 authorities")
	})
	cronProxy(CRON_EVERY_15_SECONDS, func() {
		fmt.Println("check check new authorities from blockchair")
		fmt.Println("check tatum confirmed authorities")
		fmt.Println("clean system and expire datas")
	})
	time.Sleep(10 * 365 * 24 * 3600 * time.Second)
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

func main() {

	setNewNetwork("ethereum")
	setNewNetwork("bitcoin")

	EventHandlerApplication()

	// checkUndeterminedAuthorities()
	// checkUndeterminedAuthorities()
	// checkUndeterminedAuthorities()
	// checkUndeterminedAuthorities()
	// checkUndeterminedAuthorities()
	// a := getUndeterminedAuthorities()
	// for _, value := range a {
	// 	fmt.Println(value)
	// }
}
