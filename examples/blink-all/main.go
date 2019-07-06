package main

import (
	"time"

	"github.com/hybridgroup/gopherbot"
)

func main() {
	antenna := gopherbot.Antenna()
	visor := gopherbot.Visor()
	backpack := gopherbot.Backpack()

	for {
		antenna.On()
		visor.Blue()
		backpack.Blue()
		time.Sleep(500 * time.Millisecond)

		antenna.Off()
		visor.Clear()
		backpack.Clear()
		time.Sleep(500 * time.Millisecond)
	}
}
