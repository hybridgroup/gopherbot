package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/hybridgroup/gopherbot"
)

const (
	greenJet = iota
	redJet
	redRotate
	// tiltVisor
	xmasJet
)

const (
	forward = iota
	backward
)

func main() {
	println("hi")

	backpack := gopherbot.NewBackpack()

	left := machine.GPIO{machine.BUTTONA}
	left.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLDOWN})

	right := machine.GPIO{machine.BUTTONB}
	right.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLDOWN})

	mode := redJet

	for {
		if !right.Get() {
			mode++
			if mode > xmasJet {
				mode = greenJet
			}
		}

		if !left.Get() {
			mode--
			if mode < greenJet {
				mode = xmasJet
			}
		}

		switch mode {
		case greenJet:
			backpack.Green()
		case redJet:
			backpack.Red()
		case redRotate:
			backpack.Alternate(color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0x00})
		// case tiltVisor:
		// 	tilt()
		case xmasJet:
			backpack.Xmas()
		}

		time.Sleep(200 * time.Millisecond)
	}
}
