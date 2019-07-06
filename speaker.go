package gopherbot

import (
	"machine"

	"tinygo.org/x/drivers/buzzer"
)

// SpeakerDevice is the Gopherbot speaker.
type SpeakerDevice struct {
	buzzer.Device
}

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

// Bleep makes a bleep sound using the speaker.
func (s *SpeakerDevice) Bleep() {
	// do bleep
	s.Tone(buzzer.C3, buzzer.Eighth)
}

// Bloop makes a bloop sound using the speaker.
func (s *SpeakerDevice) Bloop() {
	// do bloop
	s.Tone(buzzer.C5, buzzer.Quarter)
}
