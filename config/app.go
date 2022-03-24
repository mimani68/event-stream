package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envirnoment, WebhookAddress, TatumToken string
var ConfirmCount int

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Env file was not exist.")
		os.Exit(1)
	}
	fmt.Println("Load init in config")

	Envirnoment = os.Getenv("ENV")
	WebhookAddress = os.Getenv("CLIENT_END_POINT")
	TatumToken = os.Getenv("TATUM_API_TOKEN")
	ConfirmCount, _ = strconv.Atoi(os.Getenv("CONFIRM_COUNT"))
}
