package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var Envirnoment, WebhookAddress, TatumToken, LOG_FILE_PATH string
var ConfirmCount int
var AgeOfOldMessage time.Duration
var FAKE_FIRST_TRX_NEW, OFFLINE bool
var confErr error

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Env file was not exist.")
		os.Exit(1)
	}

	Envirnoment = os.Getenv("ENV")
	WebhookAddress = os.Getenv("CLIENT_END_POINT")
	TatumToken = os.Getenv("TATUM_API_TOKEN")
	ConfirmCount, confErr = strconv.Atoi(os.Getenv("CONFIRM_COUNT"))
	if confErr != nil {
		ConfirmCount = 5
	}
	LOG_FILE_PATH = "./logs/daily.log"
	OFFLINE = os.Getenv("MOCK") == "true"
	AgeOfOldMessage = 3 * time.Hour
	FAKE_FIRST_TRX_NEW = os.Getenv("FAKE_FIRST_TRX_NEW") == "true"
}
