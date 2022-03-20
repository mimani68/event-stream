package delay

import "time"

func SetSyncDelay(delayNumber int) {
	if delayNumber > 0 {
		time.Sleep(time.Second * time.Duration(delayNumber))
	} else {
		time.Sleep(time.Second * 2)
	}
}
