package log_handler

import (
	"fmt"
	"time"
)

const (
	ColorBlack  = string("\u001b[30m")
	ColorRed    = string("\u001b[31m")
	ColorGreen  = string("\u001b[32m")
	ColorYellow = string("\u001b[33m")
	ColorBlue   = string("\u001b[34m")
	ColorReset  = string("\u001b[0m")
)

func LoggerF(template string, params ...string) {
	switch len(params) {
	case 1:
		template = fmt.Sprintf(template, params[0])
	case 2:
		template = fmt.Sprintf(template, params[0], params[1])
	case 3:
		template = fmt.Sprintf(template, params[0], params[1], params[2])
	case 4:
		template = fmt.Sprintf(template, params[0], params[1], params[2], params[3])
	default:
		template = template
	}
	fmt.Printf("%s[%s]%s %s \n", ColorBlue, time.Now().Format(time.RFC3339), ColorReset, template)
}
