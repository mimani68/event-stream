package log_handler

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"zarinworld.ir/event/config"
)

var lock sync.Mutex

const (
	ColorBlack  = string("\u001b[30m")
	ColorRed    = string("\u001b[31m")
	ColorGreen  = string("\u001b[32m")
	ColorYellow = string("\u001b[33m")
	ColorBlue   = string("\u001b[34m")
	ColorReset  = string("\u001b[0m")
)

func LoggerF(template string, params ...string) {
	isDebug, _ := regexp.MatchString(`(?i)DEBUG`, template)
	if isDebug && config.Log_level != "debug" {
		return
	}

	switch len(params) {
	case 1:
		template = fmt.Sprintf(template, params[0])
	case 2:
		template = fmt.Sprintf(template, params[0], params[1])
	case 3:
		template = fmt.Sprintf(template, params[0], params[1], params[2])
	case 4:
		template = fmt.Sprintf(template, params[0], params[1], params[2], params[3])
	case 5:
		template = fmt.Sprintf(template, params[0], params[1], params[2], params[3], params[4])
	case 6:
		template = fmt.Sprintf(template, params[0], params[1], params[2], params[3], params[4], params[5])
	case 7:
		template = fmt.Sprintf(template, params[0], params[1], params[2], params[3], params[4], params[5], params[6])
	case 8:
		template = fmt.Sprintf(template, params[0], params[1], params[2], params[3], params[4], params[5], params[6], params[7])
	}
	fmt.Printf("%s[%s]%s %s \n", ColorBlue, time.Now().Format(time.RFC3339), ColorReset, template)
	storeFile(config.LogFilePath, template)
}

var Marshal = func(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	// b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return "nil", err
	}
	return string(b), nil
	// return bytes.NewReader(b), nil
}

func storeFile(path string, content interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	payload := map[string]interface{}{
		"time":    time.Now().Format(time.RFC3339),
		"message": content,
		"isAlert": false,
	}
	fileString, err := Marshal(payload)
	if err != nil {
		return err
	}

	if _, err = f.WriteString(clearColor(fileString) + "\n"); err != nil {
		return err
	}
	return nil
}

func clearColor(text string) string {
	color := []string{"\\u001b[30m", "\\u001b[31m", "\\u001b[32m", "\\u001b[33m", "\\u001b[34m", "\\u001b[0m"}
	for _, c := range color {
		text = strings.Replace(text, c, "", 5)
	}
	text = strings.Replace(text, "\\u003e", "more", 5)
	return text
}
