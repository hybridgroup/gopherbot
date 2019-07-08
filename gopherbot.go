// Package gopherbot is the TinyGo code for the Gopherbot programmable robotic plushie.
package gopherbot

import "machine"

const (
	// BackpackLEDCount is how many WS2812 LEDs are on the backpack, aka
	// the Circuit Playground Express.
	BackpackLEDCount = 10

	// VisorLEDCount is how many WS2812 LEDs are on the helmet visor.
	VisorLEDCount = 12
)

var (
	adcInitComplete = false
	i2cInitComplete = false
)

// EnsureADCInit makes sure that the Gopherbot ADC has been initialized, but
// is only initialized once.
func EnsureADCInit() {
	if !adcInitComplete {
		machine.InitADC()
		adcInitComplete = true
	}
}

// EnsureI2CInit makes sure that the Gopherbot I2C has been initialized, but
// is only initialized once.
func EnsureI2CInit() {
	if !i2cInitComplete {
		machine.I2C1.Configure(machine.I2CConfig{})
		i2cInitComplete = true
	}
}

// Rescale performs a direct linear rescaling of an integer from one scale to another.
//
// For example:
//
//		val := gopherbot.Rescale(25, 0, 100, 0, 10)
//
// This re-scales the number 25 from a scale that ranges from 0-100,
// to a scale that ranges from 0-10.
func Rescale(input, fromMin, fromMax, toMin, toMax int32) int32 {
	return (input-fromMin)*(toMax-toMin)/(fromMax-fromMin) + toMin
}
