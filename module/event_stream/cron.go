package event_stream

import (
	"github.com/robfig/cron/v3"
)

const (
	CRON_EVERY_2_SECONDS  = "*/2 * * * * *"
	CRON_EVERY_5_SECONDS  = "*/5 * * * * *"
	CRON_EVERY_10_SECONDS = "*/10 * * * * *"
	CRON_EVERY_15_SECONDS = "*/15 * * * * *"
	CRON_EVERY_20_SECONDS = "*/20 * * * * *"
	CRON_EVERY_30_SECONDS = "*/30 * * * * *"
	CRON_EVERY_30_MINUTES = "*/30 * * * *"
	CRON_EVERY_ONE_MINUTE = "* * * * *"
	CRON_AT_4_OCLOCK      = "0 4 * * *"
	CRON_EVERY_6_HOURS    = "0 6,12,18,23 * * *"
)

func cronProxy(cronTimeString string, cb func()) (bool, error) {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc(cronTimeString, cb)
	if err != nil {
		return false, err
	}
	// log_handler.LoggerF("Cron id is %s%s%s", string(log_handler.ColorRed), fmt.Sprint(cronId), string(log_handler.ColorReset))
	c.Start()
	return true, nil
}
