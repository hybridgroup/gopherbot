package main

import (
	"machine"
	"time"

	"github.com/hybridgroup/gopherbot"
)

const (
	greenVisor = iota
	redVisor
	// cylonVisor
	// tiltVisor
	xmasVisor
)

const (
	forward = iota
	backward
)

func main() {
	println("hi")

	visor := gopherbot.NewVisor()

	left := machine.GPIO{machine.BUTTONA}
	left.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLDOWN})

	right := machine.GPIO{machine.BUTTONB}
	right.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLDOWN})

	mode := redVisor

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
		case redVisor:
			visor.Red()
		// case cylonVisor:
		// 	cylon()
		// case tiltVisor:
		// 	tilt()
		case xmasVisor:
			visor.Xmas()
		}

		time.Sleep(200 * time.Millisecond)
	}
}
