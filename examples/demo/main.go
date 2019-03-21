package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/hybridgroup/gopherbot"
)

const (
	greenVisor = iota
	redVisor
	cylonVisor
	// tiltVisor
	xmasVisor
)

const (
	forward = iota
	backward
)

var (
	antenna *machine.GPIO
)

func main() {
	ant := machine.GPIO{machine.A2}
	ant.Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})
	antenna = &ant

	visor := gopherbot.NewVisor()
	backpack := gopherbot.NewBackpack()

	left := machine.GPIO{machine.BUTTONA}
	left.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLDOWN})

	right := machine.GPIO{machine.BUTTONB}
	right.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLDOWN})

	mode := redVisor

	go blinker()

	for {
		if !right.Get() {
			mode++
			if mode > xmasVisor {
				mode = greenVisor
			}
		}

		if !left.Get() {
			mode--
			if mode < greenVisor {
				mode = xmasVisor
			}
		}

		switch mode {
		case greenVisor:
			visor.Green()
			backpack.Clear()
		case redVisor:
			visor.Red()
			backpack.Red()
		case cylonVisor:
			visor.Cylon()
			backpack.Alternate(color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0x00})
		// case tiltVisor:
		// 	tilt()
		case xmasVisor:
			visor.Xmas()
			backpack.Xmas()
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func blinker() {
	for {
		antenna.High()
		time.Sleep(500 * time.Millisecond)
		antenna.Low()
		time.Sleep(500 * time.Millisecond)
	}
}
