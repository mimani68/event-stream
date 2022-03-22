package tatum

import (
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/http_proxy"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func GetCurrentBlock(network string) int {
	// curl --request GET \
	//   --url https://api-eu1.tatum.io/v3/bitcoin/info \
	//   --header 'x-api-key: REPLACE_KEY_VALUE'
	url := ""
	switch network {
	case config.BITCOIN:
		url = "https://api-eu1.tatum.io/v3/bitcoin/info"
	case config.ETHEREUM:
		url = "https://api-eu1.tatum.io/v3/ethereum/block/current"
	}
	header := map[string]string{"x-api-key": config.Tatum_token}
	responseString, err := http_proxy.Get(url, header)
	if err != nil {
		log_handler.LoggerF("%sTATUM%s didn't response on %s network", log_handler.ColorRed, log_handler.ColorReset, network)
		return 0
	}
	return utils.ToInt(responseString)
}
