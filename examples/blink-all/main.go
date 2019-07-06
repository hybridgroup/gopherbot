package main

import (
	"time"

	"github.com/hybridgroup/gopherbot"
)

func main() {
	antenna := gopherbot.Antenna()
	visor := gopherbot.Visor()
	backpack := gopherbot.Backpack()

	slider := gopherbot.Slider()

	for {
		if slider.IsOn() {
			antenna.On()
			visor.Blue()
			backpack.Blue()
			time.Sleep(500 * time.Millisecond)

			antenna.Off()
			visor.Off()
			backpack.Off()
			time.Sleep(500 * time.Millisecond)
		}
	}
}
