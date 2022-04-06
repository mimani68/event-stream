package main

import (
	"strings"

	"zarinworld.ir/event/config"
	"zarinworld.ir/event/module/event_stream"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func main() {
	log_handler.LoggerF("%s%s%s", log_handler.ColorBlue, "=====================", log_handler.ColorReset)
	log_handler.LoggerF("Envirnoment: %s%s%s", log_handler.ColorYellow, utils.ToString(config.Envirnoment), log_handler.ColorReset)
	log_handler.LoggerF("Mock: %s%s%s", log_handler.ColorYellow, utils.ToString(config.Offline), log_handler.ColorReset)
	log_handler.LoggerF("Log: %s%s%s", log_handler.ColorYellow, strings.ToLower(config.Log_level), log_handler.ColorReset)
	log_handler.LoggerF("Log folder: %s%s%s", log_handler.ColorYellow, strings.ToLower(config.LogFilePath), log_handler.ColorReset)
	log_handler.LoggerF("%s%s%s", log_handler.ColorBlue, "=====================", log_handler.ColorReset)
	event_stream.SetNewNetwork(config.ETHEREUM)
	event_stream.SetNewNetwork(config.BITCOIN)
	for _, address := range config.AddressList {
		event_stream.SetNewAddress(utils.ToString(address["network"]), utils.ToString(address["address"]))
	}
	stateOfApplication := make(chan string)
	// event_stream.EventHandlerModuleDev(stateOfApplication)
	event_stream.EventHandlerModule(stateOfApplication)
	<-stateOfApplication
}
