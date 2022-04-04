package git_version

import (
	"os"
	"strings"

	"zarinworld.ir/event/pkg/log_handler"
)

func ShowLatestVersion() string {
	versionByte, err := os.ReadFile("./VERSION")
	if err != nil {
		log_handler.LoggerF(err.Error())
	}
	versionString := string(versionByte)
	return strings.ReplaceAll(versionString, "\n", "")
}
