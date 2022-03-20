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

	// trx := make(map[string]interface{})
	// trx["id"] = uuid.New()
	// trx["value"] = "0.0025"
	// trx["blockNumber"] = "-1"
	// trx["trxHash"] = "2sd0000jyj2wg1sfn1y3kl13a1f3fh1k50000j"
	// trx["expireIn"] = time.Now().Add(5 * time.Minute).String()

	// https://api.blockchair.com/bitcoin/dashboards/address/bc1qcrudsryq8gcuspdz3ddvzytt8vch6l4ugfzp5y?transaction_details=true
	url := fmt.Sprintf("https://api.blockchair.com/%s/dashboards/address/%s?transaction_details=true", network, address)
	response, err := http_request.Get(url)
	if err != nil {
		return []map[string]interface{}{}
	}
	return response
}
