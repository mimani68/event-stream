package db

var UndeterminedAuthorities []interface{}
var CurrentBlockNumberList []interface{}
var NetworkList []string

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
	switch dbName {
	case AUTHORITIES:
		if len(UndeterminedAuthorities) == 0 {
			UndeterminedAuthorities = append(UndeterminedAuthorities, param)
			statOfStoring = true
		}
		for index, value := range UndeterminedAuthorities {
			valueAfterCasting := value.(map[string]string)
			if valueAfterCasting["id"] == paramAfterCasting["id"] {
				UndeterminedAuthorities[index] = param
			} else {
				UndeterminedAuthorities = append(UndeterminedAuthorities, param)
			}
		}
	}
	return statOfStoring
}

func GetAll(dbName string) []interface{} {
	result := []interface{}{}
	switch dbName {
	case AUTHORITIES:
		result = UndeterminedAuthorities
	}
	return result
}
