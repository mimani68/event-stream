package zwbaas

import (
	"github.com/google/uuid"
)

func GetAuthorities(address string) []map[string]interface{} {
	// call baas/authorities/
	// filter { incoming_value:0 , status: "ACTIVE", expire >= now }
	// authList []
	trx := make(map[string]interface{})
	trx["id"] = uuid.New()
	trx["amount"] = "0.0025"
	trx["incoming_value"] = "0"
	trx["address"] = []string{}
	trx["transactions"] = []string{}
	trx["wallet"] = []string{}
	trx["status"] = "ACTIVE"
	list := []map[string]interface{}{}
	list = append(list, trx)
	return list
}
