package blockchair

import (
	"fmt"

	"zarinworld.ir/event/config"
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
		log_handler.LoggerF("%s", err.Error())
		return []map[string]interface{}{}
	}
	trxList, err := httpRequest.ParseBlockchairListResult(responseString, address, network)
	if err != nil {
		log_handler.LoggerF("Problem occured on parsing %sBLOCKCHAIR%s on %s network", log_handler.ColorRed, log_handler.ColorReset, network)
		log_handler.LoggerF("%s", err.Error())
		log_handler.LoggerF("[DEBUG] %s", responseString)
		return []map[string]interface{}{}
	}
	currentBlock := 0
	for _, blockPerNetwork := range db.GetAll(db.BLOCKNUMBER) {
		if blockPerNetwork["id"] == network {
			currentBlock = utils.ToInt(blockPerNetwork[network])
		}
	}

	for _, trx := range trxList {
		trx["confirm"] = false
		switch network {
		case config.ETHEREUM:
			trx["confirmCount"] = currentBlock - utils.ToInt(trx["block_id"])
		case config.BITCOIN:
			trx["confirmCount"] = currentBlock - utils.ToInt(trx["block_id"])
		}
	}
	return trxList
}
