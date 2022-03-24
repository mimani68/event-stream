package main

import (
	"fmt"

	"zarinworld.ir/event/config"
	"zarinworld.ir/event/module/event_stream"
	"zarinworld.ir/event/pkg/utils"
)

func main() {
	fmt.Printf("Simulation mode: %t\n", config.Simulate_new_request)
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
