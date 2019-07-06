package gopherbot

import (
	"machine"
	"time"
)

// LEDDevice controls any of the standard LEDs on Gopherbot.
type LEDDevice struct {
	machine.Pin
	Speed time.Duration
}

// StatusLED returns the built-in LED of the Circuit Playground Express.
func StatusLED() *LEDDevice {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return &LEDDevice{
		Pin:   led,
		Speed: 500 * time.Millisecond,
	}
}

// On turns on the LED.
func (a *LEDDevice) On() {
	a.High()
}

// Off turns off the LED.
func (a *LEDDevice) Off() {
	a.Low()
}

// Blink starts the LED blinking repeatedly.
func (a *LEDDevice) Blink() {
	for {
		a.On()
		time.Sleep(a.Speed)

		a.Off()
		time.Sleep(a.Speed)
	}
}
