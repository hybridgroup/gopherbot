package gopherbot

import (
	"tinygo.org/x/drivers/buzzer"
)

// SpeakerDevice is the Gopherbot speaker.
type SpeakerDevice struct {
	buzzer.Device
}

// Bleep makes a bleep sound using the speaker.
func (s *SpeakerDevice) Bleep() {
	s.Tone(buzzer.C3, buzzer.Eighth)
}

// Bloop makes a bloop sound using the speaker.
func (s *SpeakerDevice) Bloop() {
	s.Tone(buzzer.C5, buzzer.Quarter)
}

// Blip makes a blip sound using the speaker.
func (s *SpeakerDevice) Blip() {
	s.Tone(buzzer.C6, buzzer.Eighth/8)
}
