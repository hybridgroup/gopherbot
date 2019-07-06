package gopherbot

import (
	"machine"
	"time"
)

// Antenna controls the LED of the Gopherbot Helmet's antenna.
type Antenna struct {
	LED   *machine.Pin
	Speed time.Duration
}

// NewAntenna returns a new Antenna to control the Gopherbot Antenna LED.
func NewAntenna() *Antenna {
	led := machine.A2
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return &Antenna{
		LED:   &led,
		Speed: 500 * time.Millisecond,
	}
}

// On turns on the Antenna LED.
func (a *Antenna) On() {
	a.LED.High()
}

// Off turns off the Antenna LED.
func (a *Antenna) Off() {
	a.LED.Low()
}

// Blink starts the antenna blinking repeatedly.
func (a *Antenna) Blink() {
	for {
		a.On()
		time.Sleep(a.Speed)

		a.Off()
		time.Sleep(a.Speed)
	}
}
