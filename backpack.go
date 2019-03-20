package gopherbot

import (
	"image/color"
	"machine"

	"github.com/tinygo-org/drivers/ws2812"
)

// Backpack controls the Gopherbot Backpack LED.
type Backpack struct {
	d   *ws2812.Device
	LED []color.RGBA
	alt bool
}

// NewBackpack returns a new Backpack to control Gopherbot Backpack.
func NewBackpack() *Backpack {
	neo := machine.GPIO{machine.NEOPIXELS}
	neo.Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})
	v := ws2812.New(neo)

	return &Backpack{
		d:   &v,
		LED: make([]color.RGBA, backpackLEDCount),
	}
}

// Show sets the Backpack to display the current LED array state.
func (b *Backpack) Show() {
	b.d.WriteColors(b.LED)
}

// // SetColors sets the Backpack colors to the new array
// func (b *Backpack) SetColors(colors []color.RGBA) {
// 	copy(b.LED, colors)
// }

// Clear clears the Backpack LEDs.
func (b *Backpack) Clear() {
	b.SetColor(color.RGBA{R: 0x00, G: 0x00, B: 0x00})
}

// SetColor sets the Backpack LEDs to a single color.
func (b *Backpack) SetColor(color color.RGBA) {
	for i := range b.LED {
		b.LED[i] = color
	}

	b.Show()
}

// Red turns all of the Backpack LEDs red.
func (b *Backpack) Red() {
	b.SetColor(color.RGBA{R: 0xff, G: 0x00, B: 0x00})
}

// Green turns all of the Backpack LEDs green.
func (b *Backpack) Green() {
	b.SetColor(color.RGBA{R: 0x00, G: 0xff, B: 0x00})
}

// Xmas light style that rotates
func (b *Backpack) Xmas() {
	b.Alternate(color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0xff, B: 0x00})
}

// Alternate the 2 colors.
func (b *Backpack) Alternate(color1, color2 color.RGBA) {
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
