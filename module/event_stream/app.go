package event_stream

import (
	"fmt"
	"time"

	"zarinworld.ir/event/pkg/log_handler"
)

func EventHandlerModule(stateChannel chan interface{}) {
	go stopAppliction(stateChannel)
	cronProxy(CRON_EVERY_SECONDS, func() {
		log_handler.LoggerF("check undetermined authrities")
		checkUndeterminedAuthorities()
	})
	cronProxy(CRON_EVERY_5_SECONDS, func() {
		networkList := GetNetworkList()
		for _, item := range networkList {
			// updateCurrentBlock(item["network"])
			fmt.Println(item)
		}
		// updateConfirmTransactions()
		cleanSystem()
	})
	cronProxy(CRON_EVERY_15_SECONDS, func() {
		log_handler.LoggerF("check check new authorities from blockchair")
		log_handler.LoggerF("check tatum confirmed authorities")
		log_handler.LoggerF("clean system and expire datas")
	})
}

func cleanSystem() {
	log_handler.LoggerF("clean started")
}

func stopAppliction(stateChannel chan interface{}) {
	time.Sleep(1 * 24 * 3600 * time.Second)
	defer fmt.Printf("The application stoped at [%s]\n", time.Now().Format(time.RFC3339))
	stateChannel <- "done"
}
