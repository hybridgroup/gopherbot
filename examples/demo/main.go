package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/hybridgroup/gopherbot"
)

const (
	greenVisor = iota
	redVisor
	cylonVisor
	tiltVisor
	xmasVisor
)

const (
	forward = iota
	backward
)

var (
	antenna *gopherbot.Antenna
	accel   *gopherbot.Accelerometer
	visor   *gopherbot.Visor

	pos = 0
	dir = 0
)

func main() {
	machine.I2C1.Configure(machine.I2CConfig{})

	antenna = gopherbot.NewAntenna()
	visor = gopherbot.NewVisor()
	backpack := gopherbot.NewBackpack()

	left := machine.BUTTONA
	left.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	right := machine.BUTTONB
	right.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	accel = gopherbot.NewAccelerometer()

	mode := redVisor

	go antenna.Blink()

	for {
		if !right.Get() {
			mode++
			if mode > xmasVisor {
				mode = greenVisor
			}
		}

		if !left.Get() {
			mode--
			if mode < greenVisor {
				mode = xmasVisor
			}
		}

		switch mode {
		case greenVisor:
			visor.Green()
			backpack.Alternate(color.RGBA{R: 0x00, G: 0xff, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0xff})
		case redVisor:
			visor.Red()
			backpack.Red()
		case cylonVisor:
			visor.Cylon()
			backpack.Alternate(color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0x00})
		case tiltVisor:
			tilt()
		case xmasVisor:
			visor.Xmas()
			backpack.Xmas()
		}

		time.Sleep(200 * time.Millisecond)
	}
}

// func blinker() {
// 	for {
// 		antenna.High()
// 		time.Sleep(500 * time.Millisecond)
// 		antenna.Low()
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

func tilt() {
	x, y, z, _ := accel.ReadAcceleration()
	println(x, y, z)
	//println(int64((float64(x) * 9.806)), int64((float64(y) * 9.806)), int64((float64(z) * 9.806)))

	switch {
	case x < 300000 && x > -300000:
		switch {
		case pos <= gopherbot.VisorLEDCount/2-1:
			pos++
		case pos >= gopherbot.VisorLEDCount/2+1:
			pos--
		}
	case x > 300000:
		pos++
		if pos == gopherbot.VisorLEDCount {
			pos = gopherbot.VisorLEDCount - 1
		}
	case x < -300000:
		pos--
		if pos < 0 {
			pos = 0
		}
	}

	for i := range visor.LED {
		if i == pos {
			visor.LED[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0x77}
		} else {
			visor.LED[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x77}
		}
	}

	visor.Show()
}
