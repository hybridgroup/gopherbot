package main

import (
	"time"

	"github.com/hybridgroup/gopherbot"
)

func main() {
	antenna := gopherbot.NewAntenna()

	for {
		antenna.On()
		time.Sleep(500 * time.Millisecond)

		antenna.Off()
		time.Sleep(500 * time.Millisecond)
	}
}
