package blockchair

import (
	"fmt"

	"zarinworld.ir/event/pkg/http_request"
)

func GetAddressHistory(network string, address string) []map[string]interface{} {
	// call blockchair
	//    if blockNumber === -1 => send confirm: false
	//    if blockNumber > -1 =>
	//			call currentBlock number => x
	//			send { confirm: true, confirmCount: x }

	// https://api.blockchair.com/bitcoin/dashboards/address/bc1qcrudsryq8gcuspdz3ddvzytt8vch6l4ugfzp5y?transaction_details=true
	url := fmt.Sprintf("https://api.blockchair.com/%s/dashboards/address/%s?transaction_details=true", network, address)
	httpRequest := http_request.Http{}
	responseString, err := httpRequest.Get(url, nil)
	if err != nil {
		return []map[string]interface{}{}
	}
	trxList, err := httpRequest.ParseBlockchairResult(responseString, address, network)
	if err != nil {
		return []map[string]interface{}{}
	}
	return trxList
}
