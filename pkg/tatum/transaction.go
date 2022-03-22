package tatum

import (
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/http_proxy"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func GetTrxDetails(network string, trxID string) map[string]interface{} {
	// curl --request GET \
	//   --url https://api-eu1.tatum.io/v3/bitcoin/transaction/{hash} \
	//   --header 'x-api-key: REPLACE_KEY_VALUE'
	url := ""
	switch network {
	case config.BITCOIN:
		url = "https://api-eu1.tatum.io/v3/bitcoin/transaction/" + utils.ToString(trxID)
	case config.ETHEREUM:
		url = "https://api-eu1.tatum.io/v3/ethereum/transaction/" + utils.ToString(trxID)
	}
	header := map[string]string{"x-api-key": config.Tatum_token}
	responseString, err := http_proxy.Get(url, header)
	if err != nil {
		log_handler.LoggerF("%sTATUM%s didn't response on %s network", log_handler.ColorRed, log_handler.ColorReset, network)
	}
	validator := TatumHttpValidation{}
	result, _ := validator.ParseTatumResult(responseString, network)
	return result
}
