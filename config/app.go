package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envirnoment, WebhookAddress, TatumToken string
var ConfirmCount int
var Simulate_new_request bool

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Env file was not exist.")
		os.Exit(1)
	}

	Envirnoment = os.Getenv("ENV")
	WebhookAddress = os.Getenv("CLIENT_END_POINT")
	TatumToken = os.Getenv("TATUM_API_TOKEN")
	ConfirmCount, _ = strconv.Atoi(os.Getenv("CONFIRM_COUNT"))
	Simulate_new_request = os.Getenv("SIMULATE_NEW_REQ") == "true"
}
