package event_stream

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

const (
	CRON_EVERY_SECONDS    = "* * * * * *"
	CRON_EVERY_5_SECONDS  = "*/5 * * * * *"
	CRON_EVERY_15_SECONDS = "*/15 * * * * *"
	CRON_EVERY_30_SECONDS = "*/30 * * * * *"
	CRON_EVERY_60_SECONDS = "* * * * *"
)

func cronProxy(cronTimeString string, cb func()) (bool, error) {
	c := cron.New(cron.WithSeconds())
	cronId, err := c.AddFunc(cronTimeString, cb)
	if err != nil {
		return false, err
	}
	fmt.Println(cronId)
	c.Start()
	return true, nil
}
