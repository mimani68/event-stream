package blockchair

import (
	"fmt"

	"zarinworld.ir/event/pkg/blockchain_utils"
	"zarinworld.ir/event/pkg/http_proxy"
	"zarinworld.ir/event/pkg/log_handler"
)

func GetAddressHistory(network string, address string) []map[string]interface{} {
	// https://api.blockchair.com/bitcoin/dashboards/address/bc1qcrudsryq8gcuspdz3ddvzytt8vch6l4ugfzp5y?transaction_details=true
	url := fmt.Sprintf("https://api.blockchair.com/%s/dashboards/address/%s?transaction_details=true", network, address)
	httpRequest := BlockchairHttpValidation{}
	responseString, err := http_proxy.Get(url, nil)
	blockchairStatus, _ := httpRequest.blockchairOkResponse(responseString)
	if err != nil || !blockchairStatus {
		log_handler.LoggerF("%sBLOCKCHAIR%s didn't response on %s network", log_handler.ColorRed, log_handler.ColorReset, network)
		log_handler.LoggerF("%s", err.Error())
		return []map[string]interface{}{}
	}
	trxList, err := httpRequest.ParseBlockchairListResult(responseString, address, network)
	if err != nil {
		log_handler.LoggerF("Problem occured on parsing %sBLOCKCHAIR%s on %s network", log_handler.ColorRed, log_handler.ColorReset, network)
		log_handler.LoggerF("%s", err.Error())
		return []map[string]interface{}{}
	}

	for _, trx := range trxList {
		if trx["block_id"] == float64(-1) {
			trx["confirm"] = false
			trx["confirmCount"] = 0
		} else {
			trx["confirm"] = true
			trx["confirmCount"] = blockchain_utils.ConfirmNumber(network, trx["hash"].(string))
		}
	}
	return trxList
}
