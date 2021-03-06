package tatum

import (
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/http_proxy"
	"zarinworld.ir/event/pkg/in_memory_db"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func GetCurrentBlock(network string) int {
	if config.Offline {
		return mockCurrentBlock(network)
	}
	// time.Sleep(2 * time.Second)
	result := 0
	url := ""
	switch network {
	case config.BITCOIN:
		// curl --request GET \
		//   --url https://api-eu1.tatum.io/v3/bitcoin/info \
		//   --header "x-api-key: $API_TOKEN"
		url = "https://api-eu1.tatum.io/v3/bitcoin/info"
	case config.ETHEREUM:
		url = "https://api-eu1.tatum.io/v3/ethereum/block/current"
	}
	header := map[string]string{"x-api-key": config.TatumToken}
	var responseString string
	var err error
	key := in_memory_db.KeyGenerator("tatum", "network", network)
	responseString, err = in_memory_db.Get(key)
	if err != nil {
		responseString, err = http_proxy.Get(url, header)
		in_memory_db.Set(key, responseString, 60)
		log_handler.LoggerF("[DEBUG][LIVE] network %s info", network)
	} else {
		log_handler.LoggerF("[DEBUG][CACHE] network %s info", network)
	}

	if reachRateLimitOfTatum(responseString) {
		log_handler.LoggerF("%sTATUM%s rate limit", log_handler.ColorRed, log_handler.ColorReset)
		return 0
	}
	if err != nil {
		log_handler.LoggerF("%sTATUM%s didn't response on %s network", log_handler.ColorRed, log_handler.ColorReset, network)
		return 0
	}
	tv := TatumHttpValidation{}
	switch network {
	case config.BITCOIN:
		result = tv.bitcoinCurrentBlockResponseParser(responseString)
	case config.ETHEREUM:
		result = utils.ToInt(responseString)
	}
	return result
}

func mockCurrentBlock(network string) int {
	result := 0
	switch network {
	case config.BITCOIN:
		result = 729140
	case config.ETHEREUM:
		result = 14463422
	}
	return result
}
