package gopherbot

import (
	"machine"
)

// SliderDevice is the slider switch.
type SliderDevice struct {
	machine.Pin
}

// IsOn checks to see if the slider is in on position.
func (s *SliderDevice) IsOn() bool {
	return !s.Get()
}

// Slider returns the SliderDevice.
func Slider() *SliderDevice {
	slider := machine.SLIDER
	slider.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	return &SliderDevice{slider}
}
