package config

import (
	"os"
	"strconv"
)

var (
	Envirnoment      = os.Getenv("ENV")
	WebhookAddress   = os.Getenv("CLIENT_END_POINT")
	Tatum_token      = os.Getenv("TATUM_API_TOKEN")
	Confirm_Count, _ = strconv.Atoi(os.Getenv("CONFIRM_COUNT"))
)
