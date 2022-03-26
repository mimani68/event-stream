package blockchain_utils

import (
	"math"

	"zarinworld.ir/event/config"
	"zarinworld.ir/event/pkg/db"
	"zarinworld.ir/event/pkg/tatum"
	"zarinworld.ir/event/pkg/utils"
)

func ConfirmNumber(network string, trxId string) int {
	currentBlock := 0
	trx := tatum.GetTrxDetails(network, trxId)
	result := float64(0)
	for _, blockPerNetwork := range db.GetAll(db.BLOCKNUMBER) {
		if blockPerNetwork["id"] == network {
			currentBlock = utils.ToInt(blockPerNetwork[network])
		}
	}
	if currentBlock <= 0 {
		return 0
	}
	switch network {
	case config.ETHEREUM:
		result = math.Abs(float64(currentBlock - utils.ToInt(trx["block_id"])))
	case config.BITCOIN:
		result = math.Abs(float64(currentBlock - utils.ToInt(trx["blockNumber"])))
	}
	if result == 0 {
		result = 1
	}
	return int(result)
}