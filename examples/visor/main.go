package main

import (
	"time"

	"github.com/hybridgroup/gopherbot"
)

const (
	greenVisor = iota
	redVisor
	cylonVisor
	xmasVisor
)

const (
	forward = iota
	backward
)

func main() {
	visor := gopherbot.NewVisor()

	left := gopherbot.NewLeftButton()
	right := gopherbot.NewRightButton()

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
		case xmasVisor:
			visor.Xmas()
		}

		time.Sleep(200 * time.Millisecond)
	}
}
