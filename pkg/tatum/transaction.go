package tatum

import (
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/http_proxy"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func GetTrxDetails(network string, trxID string) map[string]interface{} {
	// trx := make(map[string]interface{})
	// trx["id"] = uuid.New()
	// trx["value"] = "0.0025"
	// trx["blockNumber"] = "15605442"
	// trx["trxHash"] = "2sd0000jyj2wg1sfn1y3kl13a1f3fh1k50000j"
	// trx["expireIn"] = time.Now().Add(5 * time.Minute).String()
	// return trx

	// curl --request GET \
	//   --url https://api-eu1.tatum.io/v3/bitcoin/transaction/{hash} \ \
	//   --header 'x-api-key: REPLACE_KEY_VALUE'
	url := ""
	switch network {
	case config.BITCOIN:
		url = "https://api-eu1.tatum.io/v3/bitcoin/transaction/" + utils.ToString(trxID)
	case config.ETHEREUM:
		url = "https://api-eu1.tatum.io/v3/ethereum/transaction/" + utils.ToString(trxID)
	}
	// httpRequest := blockchair.BlockchairHttpValidation{}
	header := map[string]string{"x-api-key": config.Tatum_token}
	responseString, err := http_proxy.Get(url, header)
	if err != nil {
		log_handler.LoggerF("Problem in GetTrxDetails of %s in %s network", trxID, network)
		// return 0
	}
	a := TatumHttpValidation{}
	result, _ := a.ParseTatumResult(responseString, network)
	return result
}
