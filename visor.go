package gopherbot

import (
	"image/color"
	"machine"

	"github.com/tinygo-org/drivers/ws2812"
)

// Visor controls the Gopherbot Visor Neopixel LED.
type Visor struct {
	d       *ws2812.Device
	LED     []color.RGBA
	rg      bool
	forward bool
	pos     int
}

// NewVisor returns a new Visor to control Gopherbot Visor.
func NewVisor() *Visor {
	// TODO: point to the visor's neopixels
	neo := machine.A3
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	v := ws2812.New(neo)

	return &Visor{
		d:   &v,
		LED: make([]color.RGBA, VisorLEDCount),
	}
}

// Show sets the visor to display the current LED array state.
func (v *Visor) Show() {
	v.d.WriteColors(v.LED)
}

// Clear clears the visor.
func (v *Visor) Clear() {
	for i := range v.LED {
		v.LED[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
	}

	v.Show()
}

// Red turns all of the Visor LEDs red.
func (v *Visor) Red() {
	for i := range v.LED {
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

// Cylon visor mode.
func (v *Visor) Cylon() {
	if v.forward {
		v.pos += 2
		if v.pos >= VisorLEDCount {
			v.pos = VisorLEDCount - 2
			v.forward = false
		}
	} else {
		v.pos -= 2
		if v.pos < 0 {
			v.pos = 0
			v.forward = true
		}
	}

	for i := 0; i < VisorLEDCount; i += 2 {
		if i == v.pos {
			v.LED[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
			v.LED[i+1] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
		} else {
			v.LED[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
			v.LED[i+1] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
		}
	}

	// for i := range v.LED {
	// }

	v.Show()
}
