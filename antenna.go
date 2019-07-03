package gopherbot

import (
	"machine"
	"time"
)

// Antenna controls the LED of the Gopherbot Helmet's antenna.
type Antenna struct {
	LED *machine.Pin
}

// NewAntenna returns a new Antenna to control the Gopherbot Antenna LED.
func NewAntenna() *Antenna {
	led := machine.A2
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return &Antenna{
		LED: &led,
	}
}

// Blink starts the Antenna blinking.
func (a *Antenna) Blink() {
	for {
		a.LED.High()
		time.Sleep(500 * time.Millisecond)
		a.LED.Low()
		time.Sleep(500 * time.Millisecond)
	}
}
