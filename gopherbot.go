package gopherbot

import "github.com/tinygo-org/tinygo/src/machine"

const (
	BackpackLEDCount = 10
	VisorLEDCount    = 12
)

var (
	adcInitComplete = false
	i2cInitComplete = false
)

// EnsureADCInit makes sure that the Gopherbot ADC has been initialized.
func EnsureADCInit() {
	if !adcInitComplete {
		machine.InitADC()
		adcInitComplete = true
	}
}

// EnsureI2CInit makes sure that the Gopherbot I2C has been initialized.
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
func Rescale(input, fromMin, fromMax, toMin, toMax int32) int32 {
	return (input-fromMin)*(toMax-toMin)/(fromMax-fromMin) + toMin
}
