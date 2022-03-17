package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"zarinworld.ir/event/pkg/db"
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
	if len(db.UndeterminedAuthorities) == 0 {
		db.UndeterminedAuthorities = append(db.UndeterminedAuthorities, m)
	}
	for _, v := range db.UndeterminedAuthorities {
		if v[title] == "" {
			db.UndeterminedAuthorities = append(db.UndeterminedAuthorities, m)
		} else if v[title] != "" && v[title] != m[title] {
			v[title] = m[title]
		}
	}
}

func cleanAuthoriries() {
	fmt.Println("clean started")
}

func getUndeterminedAuthorities() []map[string]string {
	return db.UndeterminedAuthorities
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
	for _, value := range db.CurrentBlockNumberList {
		number = value[network]
	}
	return number
}

func setNewNetwork(network string) {
	db.NetworkList = append(db.NetworkList, network)
}

func gettNetworkList() []string {
	return db.NetworkList
}

func eventHandlerApplication() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/30 * * * * *", func() {
		fmt.Println("Delayed tasks, like check blockNumber and ...")
		a := gettNetworkList()
		for index := range a {
			currentBlock(a[index])
		}
		checkUndeterminedAuthorities()
	})
	c.AddFunc("* * * * * *", func() {
		fmt.Println("Instance task")
	})
	c.Start()
}

func main() {

	setNewNetwork("ethereum")
	setNewNetwork("bitcoin")
	setNewNetwork("litecoin")
	setNewNetwork("tron")

	// eventHandlerApplication()
	// time.Sleep(10 * 365 * 24 * 3600 * time.Second)

	checkUndeterminedAuthorities()
	checkUndeterminedAuthorities()
	checkUndeterminedAuthorities()
	checkUndeterminedAuthorities()
	checkUndeterminedAuthorities()
	a := getUndeterminedAuthorities()
	for _, value := range a {
		fmt.Println(value)
	}
}
