package db

// var UndeterminedAuthorities []map[string]string
// var CurrentBlockNumberList []map[string]int
// var NetworkList []string

var UndeterminedAuthorities []interface{}
var CurrentBlockNumberList []interface{}
var NetworkList []string

const (
	AUTHORITY   = "AUTHORITY"
	NETWORK     = "NETWORK"
	BLOCKNUMBER = "BLOCKNUMBER"
)

func Store(dbName string, param interface{}) bool {
	// func Store(dbName string, map[string]string) bool {
	// switch dbName {
	// case AUTHORITY:
	// 	if len(UndeterminedAuthorities) == 0 {
	// 		UndeterminedAuthorities = append(UndeterminedAuthorities, param)
	// 	}
	// 	for _, v := range UndeterminedAuthorities {
	// 		if v[title] == "" {
	// 			UndeterminedAuthorities = append(UndeterminedAuthorities, param)
	// 		} else if v[title] != "" && v[title] != param[title] {
	// 			v[title] = param[title]
	// 		}
	// 	}
	// }
	switch dbName {
	case AUTHORITY:
		if len(UndeterminedAuthorities) == 0 {
			UndeterminedAuthorities = append(UndeterminedAuthorities, param)
		}
		// for index, _ := range UndeterminedAuthorities {
		// 	if UndeterminedAuthorities[index] == "" {
		// 		UndeterminedAuthorities = append(UndeterminedAuthorities, param)
		// 	} else if UndeterminedAuthorities[index] != "" && UndeterminedAuthorities[index] != param[title] {
		// 		UndeterminedAuthorities[index] = param[title]
		// 	}
		// }
	}
	return true
}

func GetAll(dbName string) []interface{} {
	result := []interface{}{}
	switch dbName {
	case AUTHORITY:
		for index, _ := range UndeterminedAuthorities {
			result = append(result, UndeterminedAuthorities[index])
		}
	}
	return result
}
