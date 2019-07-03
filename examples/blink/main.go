package main

import (
	"machine"
	"time"
)

func main() {
	antenna := machine.A2
	antenna.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		antenna.High()
		time.Sleep(time.Millisecond * 500)

		antenna.Low()
		time.Sleep(time.Millisecond * 500)
	}
}
