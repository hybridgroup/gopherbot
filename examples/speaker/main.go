package main

import (
	"time"

	"github.com/hybridgroup/gopherbot"
)

func main() {
	speaker := gopherbot.Speaker()

	left := gopherbot.LeftButton()
	right := gopherbot.RightButton()

	for {
		if right.Pushed() {
			speaker.Bleep()
		}

		if left.Pushed() {
			speaker.Bloop()
		}

		time.Sleep(200 * time.Millisecond)
	}
}
