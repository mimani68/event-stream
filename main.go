package main

import (
	"zarinworld.ir/event/module/event_stream"
	"zarinworld.ir/event/pkg/blockchair"
)

func main() {
	event_stream.SetNewNetwork("ethereum")
	// event_stream.SetNewNetwork("bitcoin")
	// event_stream.SetNewAddress("bitcoin", "1a4c1we54v564we31sv1rg4")
	// event_stream.SetNewAddress("bitcoin", "002Qsw1v1v5Dw5O5C405LP")
	event_stream.SetNewAddress("ethereum", "0x2bb413fdadbb639584ea33c96a4caa3f5616ca70")
	// stateOfApplication := make(chan interface{})
	// event_stream.EventHandlerModule(stateOfApplication)
	// <-stateOfApplication

	//
	// D E V
	//
	blockchair.GetAddressHistory("ethereum", "0x2bb413fdadbb639584ea33c96a4caa3f5616ca70")
}
