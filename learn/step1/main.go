package main

import (
	"time"

	"github.com/hybridgroup/gopherbot"
)

func main() {
	left := gopherbot.LeftButton()
	led := gopherbot.StatusLED()

	for {
		if left.Pushed() {
			led.On()
		} else {
			led.Off()
		}

		time.Sleep(100 * time.Millisecond)
	}
}
