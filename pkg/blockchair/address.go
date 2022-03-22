package blockchair

import (
	"fmt"
	"math"

	"zarinworld.ir/event/pkg/db"
	"zarinworld.ir/event/pkg/http_proxy"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func GetAddressHistory(network string, address string) []map[string]interface{} {
	// https://api.blockchair.com/bitcoin/dashboards/address/bc1qcrudsryq8gcuspdz3ddvzytt8vch6l4ugfzp5y?transaction_details=true
	url := fmt.Sprintf("https://api.blockchair.com/%s/dashboards/address/%s?transaction_details=true", network, address)
	httpRequest := BlockchairHttpValidation{}
	responseString, err := http_proxy.Get(url, nil)
	if err != nil {
		log_handler.LoggerF("%sBLOCKCHAIR%s didn't response on %s network", log_handler.ColorRed, log_handler.ColorReset, network)
		return []map[string]interface{}{}
	}
	trxList, err := httpRequest.ParseBlockchairListResult(responseString, address, network)
	if err != nil {
		log_handler.LoggerF("Problem occured on parsing %sBLOCKCHAIR%s on %s network", log_handler.ColorRed, log_handler.ColorReset, network)
		return []map[string]interface{}{}
	}
	currentBlock := 0
	networkList := db.GetAll(db.BLOCKNUMBER)
	for _, net := range networkList {
		currentBlock = utils.ToInt(net["network"])
	}

	for _, v := range trxList {
		v["confirm"] = false
		v["confirmCount"] = utils.ToString(math.Abs(float64(currentBlock - utils.ToInt(v["blockNumber"]))))
	}
	return trxList
}
