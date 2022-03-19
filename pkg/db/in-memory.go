package db

var undeterminedAuthorities []map[string]interface{}
var currentBlockNumberList []map[string]interface{}
var networkList []map[string]interface{}
var addresList []map[string]interface{}
var newTransactions []map[string]interface{}
var transactions []map[string]interface{}

const (
	AUTHORITIES      = "authority"
	NETWORK          = "network"
	BLOCKNUMBER      = "currentBlocNumber"
	ADDRESS          = "address"
	TRANSACTIONS     = "trx"
	NEW_TRANSACTIONS = "new trx"
)

func Store(dbName string, param map[string]interface{}) bool {
	statOfStoring := false
	db := dbSelector(dbName)
	if len(*db) == 0 {
		*db = append(*db, param)
		statOfStoring = true
	}
	for index, value := range *db {
		if value["id"] == param["id"] {
			(*db)[index] = param
		} else {
			*db = append(*db, param)
		}
	}
	return statOfStoring
}

func GetAll(dbName string) []map[string]interface{} {
	result := []map[string]interface{}{}
	switch dbName {
	case AUTHORITIES:
		result = undeterminedAuthorities
	}
	return result
}

func dbSelector(dbName string) *[]map[string]interface{} {
	dbPointer := &undeterminedAuthorities
	switch dbName {
	case AUTHORITIES:
		dbPointer = &undeterminedAuthorities
	case NETWORK:
		dbPointer = &networkList
	case BLOCKNUMBER:
		dbPointer = &currentBlockNumberList
	case ADDRESS:
		dbPointer = &addresList
	case TRANSACTIONS:
		dbPointer = &transactions
	case NEW_TRANSACTIONS:
		dbPointer = &newTransactions
	}
	return dbPointer
}
