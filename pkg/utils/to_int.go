package utils

import (
	"fmt"
	"strconv"
	"strings"

	"zarinworld.ir/event/pkg/log_handler"
)

func ToInt(value interface{}) int {
	typeOfValue := fmt.Sprintf("%T", value)
	if typeOfValue == "int" {
		return value.(int)
	} else if typeOfValue == "string" {
		tmp := fmt.Sprintf("%s", value)
		res, err := strconv.Atoi(tmp)
		if err != nil {
			log_handler.LoggerF("Problem in converting string to int")
			return 0
		}
		return res
	} else if typeOfValue == "float64" {
		tmp := fmt.Sprintf("%f", value)
		digit := strings.Split(tmp, ".")
		res, err := strconv.Atoi(digit[0])
		if err != nil {
			log_handler.LoggerF("Problem in converting float64 to int")
			return 0
		}
		return res
	}
	return 0
}
