package gopherbot

import (
	"image/color"
	"machine"

	"github.com/tinygo-org/drivers/ws2812"
)

const max = 10

// Visor controls the Gopherbot Visor.
type Visor struct {
	//ws2812.New(neo)
	d    *ws2812.Device
	leds []color.RGBA
}

// NewVisor returns a new Visor to control Gopherbot Visor.
func NewVisor() *Visor {
	// TODO: point to the visor's neopixels
	neo := machine.GPIO{machine.NEOPIXELS}
	neo.Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})

	return &Visor{
		d: ws2812.New(neo),
		leds: make([]color.RGBA, max)
	}
}

// Red turns all of the Visor LEDs red.
func (v *Visor) Red() {
	for i := range leds {
		v.leds[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
	}

	v.d.WriteColors(leds)
}

// func green() {
// 	for i := range leds {
// 		leds[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0x77}
// 	}

// 	visor.WriteColors(leds)
// }

// func xmas() {
// 	rg = !rg
// 	for i := range leds {
// 		rg = !rg
// 		if rg {
// 			leds[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0x77}
// 		} else {
// 			leds[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0x77}
// 		}
// 	}

// 	visor.WriteColors(leds)
// }
