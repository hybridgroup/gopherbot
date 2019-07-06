package gopherbot

import (
	"machine"

	"tinygo.org/x/drivers/lis3dh"
	"tinygo.org/x/drivers/thermistor"
)

// Accelerometer controls the Gopherbot built-in LIS3DH.
type Accelerometer struct {
	Device *lis3dh.Device
}

// NewAccelerometer returns a new Accelerometer.
func NewAccelerometer() *Accelerometer {
	EnsureI2CInit()

	accel := lis3dh.New(machine.I2C1)
	accel.Address = lis3dh.Address1 // address on the Circuit Playground Express
	accel.Configure()
	accel.SetRange(lis3dh.RANGE_2_G)

	return &Accelerometer{
		Device: &accel,
	}
}

// ReadAcceleration returns the adjusted x, y and z axis in milli-Gs.
func (a *Accelerometer) ReadAcceleration() (x int32, y int32, z int32, err error) {
	x, y, z, err = a.Device.ReadAcceleration()
	return
}

// Thermometer controls the Gopherbot built-in thermistor.
type Thermometer struct {
	Device *thermistor.Device
}

// NewThermometer returns a new Thermometer.
func NewThermometer() *Thermometer {
	EnsureADCInit()

	s := thermistor.New(machine.TEMPSENSOR)
	s.Configure()

	return &Thermometer{
		Device: &s,
	}
}

// ReadTemperature returns the temperature in celsius milli degrees (ºC/1000)
func (t *Thermometer) ReadTemperature() (temperature int32, err error) {
	return t.Device.ReadTemperature()
}
