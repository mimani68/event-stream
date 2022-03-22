package main

import (
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/module/event_stream"
	"zarinworld.ir/event/pkg/utils"
)

func main() {
	event_stream.SetNewNetwork(config.ETHEREUM)
	// event_stream.SetNewNetwork(config.BITCOIN)
	for _, address := range config.AddressList {
		event_stream.SetNewAddress(utils.ToString(address["network"]), utils.ToString(address["address"]))
	}
	stateOfApplication := make(chan interface{})
	event_stream.EventHandlerModule(stateOfApplication)
	// event_stream.EventHandlerModuleDev(stateOfApplication)
	<-stateOfApplication
}
