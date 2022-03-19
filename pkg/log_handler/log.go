package log_handler

import (
	"fmt"
	"time"
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
	fmt.Printf("[%s] %s \n", time.Now().Format(time.RFC3339), template)
}
