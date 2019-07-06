package gopherbot

import (
	"machine"
	"time"
)

// AntennaDevice controls the LED of the Gopherbot Helmet's antenna.
type AntennaDevice struct {
	machine.Pin
	Speed time.Duration
}

// Antenna returns a the Antenna to control the Gopherbot Antenna LED.
func Antenna() *AntennaDevice {
	led := machine.A2
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return &AntennaDevice{
		Pin:   led,
		Speed: 500 * time.Millisecond,
	}
}

// On turns on the Antenna LED.
func (a *AntennaDevice) On() {
	a.High()
}

// Off turns off the Antenna LED.
func (a *AntennaDevice) Off() {
	a.Low()
}

// Blink starts the antenna blinking repeatedly.
func (a *AntennaDevice) Blink() {
	for {
		a.On()
		time.Sleep(a.Speed)

		a.Off()
		time.Sleep(a.Speed)
	}
}
