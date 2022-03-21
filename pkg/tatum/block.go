package tatum

import (
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/http_request"
	"zarinworld.ir/event/pkg/utils"
)

func GetCurrentBlock(network string) int {
	// curl --request GET \
	//   --url https://api-eu1.tatum.io/v3/bitcoin/info \
	//   --header 'x-api-key: REPLACE_KEY_VALUE'
	url := ""
	switch network {
	case config.BITCOIN:
		url = "https://api-eu1.tatum.io/v3/%s/info"
	case config.ETHEREUM:
		url = "https://api-eu1.tatum.io/v3/ethereum/block/current"
	}
	httpRequest := http_request.Http{}
	header := map[string]string{"x-api-key": config.Tatum_token}
	responseString, err := httpRequest.Get(url, header)
	if err != nil {
		return 0
	}
	return utils.ToInt(responseString)
}
