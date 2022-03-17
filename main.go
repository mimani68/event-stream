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
	m := make(map[string]string)
	m["sampl auth id"] = "auth goes here"
	db.UndeterminedAuthorities = append(db.UndeterminedAuthorities, m)
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

func cronJob() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/30 * * * * *", func() {
		fmt.Println("Delayed tasks, like check blockNumber and ...")
		currentBlock("bitcoin")
		checkUndeterminedAuthorities()
	})
	c.AddFunc("* * * * * *", func() {
		fmt.Println("Instance task")
	})
	c.Start()
}

func main() {
	// network := "bitcoin"

	// cronJob()
	// time.Sleep(10 * 365 * 24 * 3600 * time.Second)

	checkUndeterminedAuthorities()
	a := getUndeterminedAuthorities()
	for _, value := range a {
		fmt.Println(value)
	}
}
