package main

import (
	"image/color"
	"time"

	"github.com/hybridgroup/gopherbot"
)

const (
	greenJet = iota
	redJet
	redRotate
)

const (
	forward = iota
	backward
)

var (
	visor *gopherbot.VisorDevice
)

func main() {
	left := gopherbot.LeftButton()
	right := gopherbot.RightButton()
	slider := gopherbot.Slider()

	led := gopherbot.StatusLED()
	backpack := gopherbot.Backpack()
	antenna := gopherbot.Antenna()

	visor = gopherbot.Visor()

	mode := greenJet

	go antenna.Blink()

	for {
		statusOn := false

		if left.Pushed() {
			statusOn = true
			mode--
			if mode < greenJet {
				mode = redRotate
			}
		}

		if right.Pushed() {
			statusOn = true
			mode++
			if mode > redRotate {
				mode = greenJet
			}
		}

		if statusOn {
			led.On()
		} else {
			led.Off()
		}

		switch mode {
		case greenJet:
			visor.Green()
		case redJet:
			visor.Red()
		case redRotate:
			visor.Cylon()
		}

		if slider.IsOn() {
			switch mode {
			case greenJet:
				backpack.Green()
			case redJet:
				backpack.Red()
			case redRotate:
				backpack.Alternate(color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0x00})
			}
		} else {
			backpack.Off()
		}

		time.Sleep(100 * time.Millisecond)
	}
}
