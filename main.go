package main

import (
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/module/event_stream"
)

func main() {
	event_stream.SetNewNetwork(config.ETHEREUM)
	// event_stream.SetNewNetwork(config.BITCOIN)
	// event_stream.SetNewAddress(config.BITCOIN, "1a4c1we54v564we31sv1rg4")
	// event_stream.SetNewAddress(config.BITCOIN, "002Qsw1v1v5Dw5O5C405LP")
	event_stream.SetNewAddress(config.ETHEREUM, "0x2bb413fdadbb639584ea33c96a4caa3f5616ca70")
	stateOfApplication := make(chan interface{})
	event_stream.EventHandlerModuleDev(stateOfApplication)
	<-stateOfApplication
}
