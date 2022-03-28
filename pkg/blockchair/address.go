package blockchair

import (
	"fmt"
	"time"

	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/blockchain_utils"
	"zarinworld.ir/event/pkg/http_proxy"
	"zarinworld.ir/event/pkg/log_handler"
)

func GetAddressHistory(network string, address string) []map[string]interface{} {
	if config.MOCK {
		return mockAddressHistory(network)
	}
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

var mockStart = time.Now()

func mockAddressHistory(network string) []map[string]interface{} {
	result := []map[string]interface{}{}
	switch network {
	case config.BITCOIN:
		result = []map[string]interface{}{
			{
				"balance_change": 521,
				"block_id":       float64(729130),
				"hash":           "e7e027e80d036b4faa2ec5a8e2d8ae584df9e3c566c407483add1edcdc06080f",
				"time":           "2022-03-26 16:06:54",
			},
			// {
			// 	"balance_change": -297,
			// 	"block_id":       float64(729126),
			// 	"hash":           "f606d8aaa00235cc8922227f4293cf80d0b93a3242f589b6d08e3965ad6fff96",
			// 	"time":           "2022-03-26 15:24:48",
			// },
			// {
			// 	"balance_change": 472,
			// 	"block_id":       float64(729126),
			// 	"hash":           "ebf2cc0448c9bf25593be90b43240289a035d1f299614b8cfae60cb8e4debe59",
			// 	"time":           "2022-03-26 15:24:48",
			// },
		}
		if time.Now().After(mockStart.Add(10 * time.Second)) {
			mockStart = time.Now()
			result = append(result, map[string]interface{}{
				"balance_change": 98,
				"block_id":       float64(-1),
				"hash":           "xxx",
				"time":           time.Now().Format(time.RFC3339),
			})
		}
		if time.Now().After(mockStart.Add(20 * time.Second)) {
			mockStart = time.Now()
			result = append(result, map[string]interface{}{
				"balance_change": 100,
				"block_id":       float64(-1),
				"hash":           "xxxz",
				"time":           time.Now().Format(time.RFC3339),
			})
		}
	}
	return result
}
