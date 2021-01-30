// +build circuitplay_express

package gopherbot

import (
	"machine"

	"tinygo.org/x/drivers/buzzer"
)

// Speaker returns the SpeakerDevice.
func Speaker() *SpeakerDevice {
	enable := machine.PA30
	enable.Configure(machine.PinConfig{Mode: machine.PinOutput})
	enable.Set(true)

	speaker := machine.A0
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})

	bzr := buzzer.New(speaker)
	return &SpeakerDevice{bzr}
}
