// +build circuitplay_bluefruit

package gopherbot

import (
	"tinygo.org/x/drivers/buzzer"

	"machine"
)

// Speaker returns the SpeakerDevice.
func Speaker() *SpeakerDevice {
	enable := machine.D11
	enable.Configure(machine.PinConfig{Mode: machine.PinOutput})
	enable.Set(true)

	speaker := machine.D12
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})

	bzr := buzzer.New(speaker)
	return &SpeakerDevice{bzr}
}
