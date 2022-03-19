package main

import "zarinworld.ir/event/module/event_stream"

func main() {
	event_stream.SetNewNetwork("ethereum")
	event_stream.SetNewNetwork("bitcoin")
	event_stream.SetNewAddress("bitcoin", "1a4c1we54v564we31sv1rg4")
	event_stream.SetNewAddress("bitcoin", "002Qsw1v1v5Dw5O5C405LP")
	stateOfApplication := make(chan interface{})
	event_stream.EventHandlerModule(stateOfApplication)
	<-stateOfApplication
}
