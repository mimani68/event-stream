package utils

import (
	"fmt"
	"strconv"

	"zarinworld.ir/event/pkg/log_handler"
)

func ToInt(value interface{}) int {
	res, err := strconv.Atoi(fmt.Sprint(value))
	if err != nil {
		log_handler.LoggerF("Problem in converting string to int")
		return 0
	}
	return res
}
