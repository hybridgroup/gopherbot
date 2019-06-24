package main

import (
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

func main() {
	visor := gopherbot.NewVisor()

	left := machine.BUTTONA
	left.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	right := machine.BUTTONB
	right.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

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
		case cylonVisor:
			visor.Cylon()
		// case tiltVisor:
		// 	tilt()
		case xmasVisor:
			visor.Xmas()
		}

		time.Sleep(200 * time.Millisecond)
	}
}
