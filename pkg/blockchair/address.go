package blockchair

import (
	"time"

	"github.com/google/uuid"
)

func GetAddressHistory(address string) []map[string]interface{} {
	// call blockchair
	//    if blockNumber === -1 => send confirm: false
	//    if blockNumber > -1 =>
	//			call currentBlock number => x
	//			send { confirm: true, confirmCount: x }
	trx := make(map[string]interface{})
	trx["id"] = uuid.New()
	trx["value"] = "0.0025"
	trx["blockNumber"] = "-1"
	trx["trxHash"] = "2sd0000jyj2wg1sfn1y3kl13a1f3fh1k50000j"
	trx["expireIn"] = time.Now().Add(5 * time.Minute).String()
	list := []map[string]interface{}{}
	list = append(list, trx)
	return list
}
