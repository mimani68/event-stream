package tatum

import (
	"time"
)

func Ping() map[string]interface{} {
	trx := make(map[string]interface{})
	trx["msg"] = "Pong"
	trx["amount"] = time.Now().Format(time.RFC3339)
	return trx
}
