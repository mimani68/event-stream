package blockchair

import (
	"fmt"
	"math/rand"
	"regexp"

	"github.com/bxcodec/faker/v3"
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/blockchain_utils"
	"zarinworld.ir/event/pkg/http_proxy"
	"zarinworld.ir/event/pkg/log_handler"
)

func GetAddressHistory(network string, address string) []map[string]interface{} {
	// https://api.blockchair.com/bitcoin/dashboards/address/bc1qcrudsryq8gcuspdz3ddvzytt8vch6l4ugfzp5y?transaction_details=true
	url := fmt.Sprintf("https://api.blockchair.com/%s/dashboards/address/%s?transaction_details=true", network, address)
	httpRequest := BlockchairHttpValidation{}
	var responseString string
	var err error
	if !config.OFFLINE {
		responseString, err = http_proxy.Get(url, nil)
	} else if config.OFFLINE {
		responseString = mockAddressHistory(network, address)
	}
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
	dev := true
	for _, trx := range trxList {
		if config.FAKE_FIRST_TRX_NEW {
			if dev {
				dev = false
				trx["block_id"] = float64(-1)
			}
		}
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

func mockAddressHistory(network string, address string) string {
	result := ""
	switch network {
	case config.BITCOIN:
		result = `
		{
			"context": {
				"api": {
					"documentation": "https://blockchair.com/api/docs",
					"last_major_update": "2021-07-19 00:00:00",
					"next_major_update": null,
					"notice": ":)",
					"version": "2.0.95-ie"
				},
				"cache": {
					"duration": 60,
					"live": true,
					"since": "2022-03-28 07:49:27",
					"time": null,
					"until": "2022-03-28 07:50:27"
				},
				"code": 200,
				"full_time": 0.20636701583862305,
				"limit": "100,100",
				"market_price_usd": 46901,
				"offset": "0,0",
				"render_time": 0.04719114303588867,
				"request_cost": 2,
				"results": 1,
				"servers": "API4,BTC5",
				"source": "D",
				"state": 729354,
				"time": 0.15917587280273438
			},
			"data": {
				"` + address + `": {
					"address": {
						"balance": 1129,
						"balance_usd": 0.52951229,
						"first_seen_receiving": "2022-02-20 15:52:48",
						"first_seen_spending": "2022-02-27 15:36:17",
						"last_seen_receiving": "2022-03-26 16:06:54",
						"last_seen_spending": "2022-03-26 15:24:48",
						"output_count": 32,
						"received": 30089,
						"received_usd": 12.3477,
						"script_hex": "0014c0f8d80c803a31c805a28b5ac1116b3b317d7ebc",
						"scripthash_type": null,
						"spent": 28960,
						"spent_usd": 11.8716,
						"transaction_count": 34,
						"type": "witness_v0_scripthash",
						"unspent_output_count": 2
					},
					"transactions": [
						{
							"balance_change": "` + fmt.Sprintf("%d", rand.Intn(10)) + `",
							"block_id": -1,
							"hash": "` + faker.Password() + `",
							"time": "` + faker.Date() + `"
						}
					],
					"utxo": [
						{
							"block_id": 729130,
							"index": 0,
							"transaction_hash": "e7e027e80d036b4faa2ec5a8e2d8ae584df9e3c566c407483add1edcdc06080f",
							"value": 521
						}
					]
				}
			}
		}							
		`
		// mockStart = time.Now()
		// for i := 0; i < 10; i++ {
		// 	time.Sleep(2 * time.Second)
		// 	if time.Now().After(mockStart.Add(2 * time.Second)) {
		// 		result = `
		// 		{
		// 			"context": {
		// 				"code": 200
		// 			},
		// 			"data": {
		// 				"` + address + `": {
		// 					"address": {
		// 						"balance": 1129,
		// 						"balance_usd": 0.52951229
		// 					},
		// 					"transactions": [
		// 						{
		// 							"balance_change": ` + string(rune(rand.Intn(100))) + `,
		// 							"block_id": -1,
		// 							"hash": ` + string(rune(rand.Intn(100))) + `,
		// 							"time": "2022-03-26 18:06:54"
		// 						}
		// 					]
		// 				}
		// 			}
		// 		}
		// 		`
		// 	}
		// }

	}
	result = regexp.MustCompile(`\r?\n`).ReplaceAllString(result, "")
	result = regexp.MustCompile(`\t`).ReplaceAllString(result, "")
	return result
}
