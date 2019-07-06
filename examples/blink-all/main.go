package main

import (
	"time"

	"github.com/hybridgroup/gopherbot"
)

func main() {
	led := gopherbot.StatusLED()
	antenna := gopherbot.Antenna()
	visor := gopherbot.Visor()
	backpack := gopherbot.Backpack()

	slider := gopherbot.Slider()

	for {
		if slider.IsOn() {
			led.On()
			antenna.On()
			visor.Blue()
			backpack.Blue()
			time.Sleep(500 * time.Millisecond)

			led.Off()
			antenna.Off()
			visor.Off()
			backpack.Off()
			time.Sleep(500 * time.Millisecond)
		}
	}
}
