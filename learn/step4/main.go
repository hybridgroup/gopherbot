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
	xmasJet
)

const (
	forward = iota
	backward
)

func main() {
	left := gopherbot.LeftButton()
	right := gopherbot.RightButton()
	slider := gopherbot.Slider()

	led := gopherbot.StatusLED()
	backpack := gopherbot.Backpack()

	mode := greenJet

	for {
		statusOn := false

		if left.Pushed() {
			statusOn = true
			mode--
			if mode < greenJet {
				mode = xmasJet
			}
		}

		if right.Pushed() {
			statusOn = true
			mode++
			if mode > xmasJet {
				mode = greenJet
			}
		}

		if statusOn {
			led.On()
		} else {
			led.Off()
		}

		if slider.IsOn() {
			switch mode {
			case greenJet:
				backpack.Green()
			case redJet:
				backpack.Red()
			case redRotate:
				backpack.Alternate(color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0x00})
			case xmasJet:
				backpack.Xmas()
			}
		} else {
			backpack.Off()
		}

		time.Sleep(100 * time.Millisecond)
	}
}
