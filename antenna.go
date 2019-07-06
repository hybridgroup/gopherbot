package gopherbot

import (
	"machine"
	"time"
)

// AntennaDevice controls the LED of the Gopherbot Helmet's antenna.
type AntennaDevice struct {
	LEDDevice
}

// Antenna returns a the Antenna to control the Gopherbot Antenna LED.
func Antenna() *AntennaDevice {
	led := machine.A2
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return &AntennaDevice{
		LEDDevice: LEDDevice{
			Pin:   led,
			Speed: 500 * time.Millisecond,
		},
	}
}
