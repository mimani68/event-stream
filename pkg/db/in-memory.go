package db

var undeterminedAuthorities []interface{}
var currentBlockNumberList []interface{}
var networkList []interface{}
var addresList []interface{}
var newTransactions []interface{}
var transactions []interface{}

const (
	AUTHORITIES      = "authority"
	NETWORK          = "network"
	BLOCKNUMBER      = "currentBlocNumber"
	ADDRESS          = "address"
	TRANSACTIONS     = "trx"
	NEW_TRANSACTIONS = "new trx"
)

func Store(dbName string, param interface{}) bool {
	statOfStoring := false
	paramAfterCasting := param.(map[string]string)
	db := dbSelector(dbName)
	if len(*db) == 0 {
		*db = append(*db, param)
		statOfStoring = true
	}
	for index, value := range *db {
		valueAfterCasting := value.(map[string]string)
		if valueAfterCasting["id"] == paramAfterCasting["id"] {
			(*db)[index] = param
		} else {
			*db = append(*db, param)
		}
	}
	return statOfStoring
}

func GetAll(dbName string) []interface{} {
	result := []interface{}{}
	switch dbName {
	case AUTHORITIES:
		result = undeterminedAuthorities
	}
	return result
}

func dbSelector(dbName string) *[]interface{} {
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
