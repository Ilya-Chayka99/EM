package Task_7

import "time"

var timeNow = time.Now

func discount() int {
	if timeNow().Hour() < 11 {
		return 10
	}
	return 2
}
