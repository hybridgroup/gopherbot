package gopherbot

import (
	"machine"

	"tinygo.org/x/drivers/lis3dh"
	"tinygo.org/x/drivers/thermistor"
)

// AccelerometerDevice controls the Gopherbot built-in LIS3DH.
type AccelerometerDevice struct {
	lis3dh.Device
}

// Accelerometer returns the AccelerometerDevice.
func Accelerometer() *AccelerometerDevice {
	EnsureI2CInit()

	accel := lis3dh.New(machine.I2C1)
	accel.Address = lis3dh.Address1 // address on the Circuit Playground Express
	accel.Configure()
	accel.SetRange(lis3dh.RANGE_2_G)

	return &AccelerometerDevice{
		Device: accel,
	}
}

// ThermometerDevice controls the Gopherbot built-in thermistor.
type ThermometerDevice struct {
	thermistor.Device
}

// Thermometer returns the ThermometerDevice.
func Thermometer() *ThermometerDevice {
	EnsureADCInit()

	s := thermistor.New(machine.TEMPSENSOR)
	s.Configure()

	return &ThermometerDevice{
		Device: s,
	}
}
