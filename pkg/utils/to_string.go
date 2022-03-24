package utils

import (
	"fmt"
	"strconv"
)

func ToString(value interface{}) string {
	result := ""
	if fmt.Sprintf("%T", value) == "bool" {
		result = strconv.FormatBool(value.(bool))
	} else {
		result = fmt.Sprintf("%s", value)
	}
	return result
}
