package gopherbot

import (
	"image/color"
	"machine"

	"github.com/tinygo-org/drivers/ws2812"
)

const max = 10

// Visor controls the Gopherbot Visor Neopixel LED.
type Visor struct {
	d   *ws2812.Device
	LED []color.RGBA
	rg  bool
}

// NewVisor returns a new Visor to control Gopherbot Visor.
func NewVisor() *Visor {
	// TODO: point to the visor's neopixels
	neo := machine.GPIO{machine.NEOPIXELS}
	neo.Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})
	v := ws2812.New(neo)

	return &Visor{
		d:   &v,
		LED: make([]color.RGBA, max),
	}
}

// Show sets the visor to display the current LED array state.
func (v *Visor) Show() {
	v.d.WriteColors(v.LED)
}

// Clear clears the visor.
func (v *Visor) Clear() {
	for i := range v.leds {
		v.LED[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
	}

	v.Show()
}

// Red turns all of the Visor LEDs red.
func (v *Visor) Red() {
	for i := range v.leds {
		v.LED[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
	}

	v.Show()
}

// Green turns all of the Visor LEDs green.
func (v *Visor) Green() {
	for i := range v.LED {
		v.LED[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
	}

	v.Show()
}

// Xmas light style
func (v *Visor) Xmas() {
	v.rg = !v.rg
	for i := range v.LED {
		v.rg = !v.rg
		if v.rg {
			v.LED[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
		} else {
			v.LED[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
		}
	}

	v.Show()
}
