package db

var undeterminedAuthorities []map[string]interface{}
var currentBlockNumberList []map[string]interface{}
var networkList []map[string]interface{}
var addresList []map[string]interface{}
var newTransactions []map[string]interface{}
var transactions []map[string]interface{}
var events []map[string]interface{}

const (
	AUTHORITIES      = "authority"
	NETWORK          = "network"
	BLOCKNUMBER      = "currentBlocNumber"
	ADDRESS          = "address"
	TRANSACTIONS     = "trx"
	NEW_TRANSACTIONS = "new trx"
	EVENTS           = "events"
)

func Store(dbName string, param map[string]interface{}) bool {
	statOfStoring := false
	db := dbSelector(dbName)
	if len(*db) == 0 {
		*db = append(*db, param)
		statOfStoring = true
	} else {
		for index, value := range *db {
			if value["id"] == param["id"] {
				(*db)[index] = param
				statOfStoring = true
			}
		}
		if !statOfStoring {
			*db = append(*db, param)
		}
	}
	return statOfStoring
}

func GetAll(dbName string) []map[string]interface{} {
	dbMemoryAddress := dbSelector(dbName)
	return *dbMemoryAddress
}

// func RemoveObject(dbName string) []map[string]interface{} {
// 	dbMemoryAddress := dbSelector(dbName)
// 	index := 2
// 	return append(*dbMemoryAddress[:index], *dbMemoryAddress[index+01:]...)
// }

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
	case EVENTS:
		dbPointer = &events

	}
	return dbPointer
}
