package tatum

import (
	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/http_proxy"
	"zarinworld.ir/event/pkg/log_handler"
	"zarinworld.ir/event/pkg/utils"
)

func GetTrxDetails(network string, trxID string) map[string]interface{} {
	url := ""
	switch network {
	case config.BITCOIN:
		// curl --request GET \
		//   --url https://api-eu1.tatum.io/v3/bitcoin/transaction/{hash} \
		//   --header 'x-api-key: REPLACE_KEY_VALUE'
		url = "https://api-eu1.tatum.io/v3/bitcoin/transaction/" + utils.ToString(trxID)
	case config.ETHEREUM:
		url = "https://api-eu1.tatum.io/v3/ethereum/transaction/" + utils.ToString(trxID)
	}
	header := map[string]string{"x-api-key": config.TatumToken}
	var responseString string
	var err error
	if !config.MOCK {
		responseString, err = http_proxy.Get(url, header)
		if reachRateLimitOfTatum(responseString) {
			log_handler.LoggerF("%sTATUM%s rate limit", log_handler.ColorRed, log_handler.ColorReset)
			return map[string]interface{}{}
		}
	} else {
		responseString = mockStringTrxDetails(network)
	}

	if err != nil {
		log_handler.LoggerF("%sTATUM%s didn't response on %s network", log_handler.ColorRed, log_handler.ColorReset, network)
	}
	validator := TatumHttpValidation{}
	result, _ := validator.ParseTatumResult(responseString, network)
	return result
}

func mockStringTrxDetails(network string) string {
	result := ""
	switch network {
	case config.ETHEREUM:
		result = `
		{
			"blockHash": "0x013fb044f367dd43cea14cff72883f1aeeaed94292aa007a614a5cb3331bedcf",
			"blockNumber": 13704493,
			"contractAddress": null,
			"cumulativeGasUsed": 2692070,
			"effectiveGasPrice": "0x2b5a1016c7",
			"from": "0xae45a8240147e6179ec7c9f92c5a18f9a97b3fca",
			"gas": 21000,
			"gasPrice": "186194597575",
			"gasUsed": 21000,
			"hash": "0xb6cdcb77ed36537f8eee50ecc0e9b5ca0e784d7765aa28f212901947473af70f",
			"logs": [],
			"logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			"nonce": 764153,
			"status": true,
			"to": "0x7f1fae878382e968d1abda7192b060e22cef2342",
			"transactionHash": "0xb6cdcb77ed36537f8eee50ecc0e9b5ca0e784d7765aa28f212901947473af70f",
			"transactionIndex": 47,
			"type": "0x0",
			"value": "8367399020422925"
		}
	`
	case config.BITCOIN:
		result = `
		{
			"blockNumber": 729141,
			"fee": 358,
			"hash": "c1d1ae771c10446ac7fff430ea3366ed14401d6bcd1a20d487b16703426fc7a2",
			"hex": "0200000000010247f27e9f73aee21ee710dd2766c36a749c6caf1f7ae47a88c63e8de9b9bed43f0000000000fdffffffb3a32db3c99c7122de9327bd20105fda3279352a68088ad968ce1bf0b81279a60100000000fdffffff01c7c900000000000017a914b0b18823e2576b89c7548ed434efc3a1481ccbe58702473044022070c4204eaccf1c7b38fc3ce2bd7cb7ddd5269da3de495b9467113d2d702e676802207ff08730501a0ce67484b64fe7d9340342c0706951c4e5bd01659ba396be0d190121037b6bf0510f79f92f543ed29a9a68eebd6cee800a76cdf00a5cc538a6950b764a02473044022019a7226bbe3aa829b0a8612a368ffcf1c71a71bdcc4c177de48fd3e98548524b02205fa2424d324d5c08a3fad9052432d38f3e1aaf482f6a2473201fe931233318130121037fda4c73e683a688d751387b4060d6e4ccbc7e71b326c2b4d8853dfdf138d2142d200b00",
			"index": 3649,
			"inputs": [
			  {
				"prevout": {
				  "hash": "3fd4beb9e98d3ec6887ae47a1faf6c9c746ac36627dd10e71ee2ae739f7ef247",
				  "index": 0
				},
				"sequence": 4294967293,
				"script": "",
				"coin": {
				  "version": 2,
				  "height": 698104,
				  "value": 40513,
				  "script": "0014ad19dda2a4e99ce58bc9f823ba2f52355105f07e",
				  "address": "bc1q45vamg4yaxwwtz7flq3m5t6jx4gstur7d5f2sj",
				  "coinbase": false
				}
			  },
			  {
				"prevout": {
				  "hash": "a67912b8f01bce68d98a08682a357932da5f1020bd2793de22719cc9b32da3b3",
				  "index": 1
				},
				"sequence": 4294967293,
				"script": "",
				"coin": {
				  "version": 2,
				  "height": 723469,
				  "value": 11500,
				  "script": "00148162445d3dc4a8d7bbb94da5a0efd2e95b8191b5",
				  "address": "bc1qs93yghfacj5d0waefkj6pm7ja9dcryd4n94y5n",
				  "coinbase": false
				}
			  }
			],
			"locktime": 729133,
			"outputs": [
			  {
				"value": 51655,
				"script": "a914b0b18823e2576b89c7548ed434efc3a1481ccbe587",
				"address": "3HoHXRyerKEoX3B9w9cc1CaQzNoAjrrvHs"
			  }
			],
			"size": 340,
			"time": 1648319277,
			"version": 2,
			"vsize": 178,
			"weight": 712,
			"witnessHash": "fda50df3ffd7f6483892efc7b7c7c31c436f389f47e2bd634c92964e6e5a4811"
		  }
		`
	}
	return result
}
