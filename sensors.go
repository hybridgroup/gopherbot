package gopherbot

import (
	"machine"

	"tinygo.org/x/drivers/lis3dh"
)

// Accelerometer controls the Gopherbot built-in LIS3DH.
type Accelerometer struct {
	Device *lis3dh.Device
}

// NewAccelerometer returns a new Accelerometer.
func NewAccelerometer() *Accelerometer {
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
