package event_stream

import (
	"fmt"
	"time"
)

func EventHandlerModule(stateChannel chan interface{}) {
	go stopAppliction(stateChannel)
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
}

func cleanSystem() {
	fmt.Println("clean started")
}

func stopAppliction(stateChannel chan interface{}) {
	time.Sleep(3 * time.Second)
	// time.Sleep(1 * 24 * 3600 * time.Second)
	defer fmt.Println("bye")
	stateChannel <- "done"
}
