package main

import (
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/module/event_stream"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func main() {
	log_handler.LoggerF("Envirnoment: %s%s%s", log_handler.ColorYellow, utils.ToString(config.Envirnoment), log_handler.ColorReset)
	log_handler.LoggerF("MOCK: %s%s%s", log_handler.ColorYellow, utils.ToString(config.OFFLINE), log_handler.ColorReset)
	// event_stream.SetNewNetwork(config.ETHEREUM)
	event_stream.SetNewNetwork(config.BITCOIN)
	for _, address := range config.AddressList {
		event_stream.SetNewAddress(utils.ToString(address["network"]), utils.ToString(address["address"]))
	}
	stateOfApplication := make(chan string)
	// event_stream.EventHandlerModuleDev(stateOfApplication)
	event_stream.EventHandlerModule(stateOfApplication)
	<-stateOfApplication
}
