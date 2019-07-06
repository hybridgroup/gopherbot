package gopherbot

import (
	"machine"
)

// Button is a button.
type Button struct {
	machine.Pin
}

// Pushed checks to see if the button is being pushed.
func (b *Button) Pushed() bool {
	return !b.Get()
}

// NewLeftButton returns a new left Button.
func NewLeftButton() *Button {
	left := machine.BUTTONA
	left.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	return &Button{left}
}

// NewRightButton returns a new right Button.
func NewRightButton() *Button {
	right := machine.BUTTONB
	right.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	return &Button{right}
}
