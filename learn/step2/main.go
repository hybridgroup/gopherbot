package main

import (
	"time"

	"github.com/hybridgroup/gopherbot"
)

func main() {
	left := gopherbot.LeftButton()
	right := gopherbot.RightButton()

	led := gopherbot.StatusLED()

	for {
		statusOn := false

		if left.Pushed() {
			statusOn = true
		}

		if right.Pushed() {
			statusOn = true
		}

		if statusOn {
			led.On()
		} else {
			led.Off()
		}

		time.Sleep(100 * time.Millisecond)
	}
}
