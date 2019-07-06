package gopherbot

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ws2812"
)

// BackpackDevice controls the Gopherbot Backpack LED.
type BackpackDevice struct {
	ws2812.Device
	LED []color.RGBA
	alt bool
}

// Backpack returns a BackpackDevice to control Gopherbot Backpack.
func Backpack() *BackpackDevice {
	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	v := ws2812.New(neo)

	return &BackpackDevice{
		Device: v,
		LED:    make([]color.RGBA, BackpackLEDCount),
	}
}

// Show sets the Backpack to display the current LED array state.
func (b *BackpackDevice) Show() {
	b.WriteColors(b.LED)
}

// Clear clears the Backpack LEDs.
func (b *BackpackDevice) Clear() {
	b.SetColor(color.RGBA{R: 0x00, G: 0x00, B: 0x00})
}

// SetColor sets the Backpack LEDs to a single color.
func (b *BackpackDevice) SetColor(color color.RGBA) {
	for i := range b.LED {
		b.LED[i] = color
	}

	b.Show()
}

// Red turns all of the Backpack LEDs red.
func (b *BackpackDevice) Red() {
	b.SetColor(color.RGBA{R: 0xff, G: 0x00, B: 0x00})
}

// Green turns all of the Backpack LEDs green.
func (b *BackpackDevice) Green() {
	b.SetColor(color.RGBA{R: 0x00, G: 0xff, B: 0x00})
}

// Xmas light style that rotates
func (b *BackpackDevice) Xmas() {
	b.Alternate(color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0xff, B: 0x00})
}

// Alternate the 2 colors.
func (b *BackpackDevice) Alternate(color1, color2 color.RGBA) {
	b.alt = !b.alt
	for i := range b.LED {
		b.alt = !b.alt
		if b.alt {
			b.LED[i] = color1
		} else {
			b.LED[i] = color2
		}
	}

	b.Show()
}
