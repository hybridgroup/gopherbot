package main

import (
	"image/color"
	"time"

	"github.com/hybridgroup/gopherbot"
)

const (
	greenJet = iota
	redJet
	redRotate
	roomTemperature
	tiltVisor
)

const (
	forward = iota
	backward
)

var (
	visor  *gopherbot.VisorDevice
	thermo *gopherbot.ThermometerDevice
	accel  *gopherbot.AccelerometerDevice

	pos = 0
)

func main() {
	left := gopherbot.LeftButton()
	right := gopherbot.RightButton()
	slider := gopherbot.Slider()

	led := gopherbot.StatusLED()
	backpack := gopherbot.Backpack()
	antenna := gopherbot.Antenna()

	visor = gopherbot.Visor()
	thermo = gopherbot.Thermometer()
	accel = gopherbot.Accelerometer()

	mode := tiltVisor

	go antenna.Blink()

	for {
		statusOn := false

		if left.Pushed() {
			statusOn = true
			mode--
			if mode < greenJet {
				mode = tiltVisor
			}
		}

		if right.Pushed() {
			statusOn = true
			mode++
			if mode > tiltVisor {
				mode = greenJet
			}
		}

		if statusOn {
			led.On()
		} else {
			led.Off()
		}

		switch mode {
		case greenJet:
			visor.Green()
		case redJet:
			visor.Red()
		case redRotate:
			visor.Cylon()
		case roomTemperature:
			showTemperature()
		case tiltVisor:
			tilt()
		}

		if slider.IsOn() {
			switch mode {
			case greenJet:
				backpack.Green()
			case redJet:
				backpack.Red()
			case redRotate:
				backpack.Alternate(color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0x00})
			case roomTemperature:
				backpack.Off()
			}
		} else {
			backpack.Off()
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func showTemperature() {
	temp, _ := thermo.ReadTemperature()
	ts := gopherbot.Rescale(temp, 0, 35000, 0, 255)
	var c color.RGBA

	switch {
	case temp < 20000:
		c = color.RGBA{R: 0x00, G: 0x00, B: byte(ts)}
	case temp < 23000:
		c = color.RGBA{R: 0x00, G: byte(ts), B: byte(ts)}
	case temp < 26000:
		c = color.RGBA{R: 0x00, G: byte(ts), B: 0x00}
	case temp < 29000:
		c = color.RGBA{R: byte(ts), G: byte(ts), B: 0x00}
	default:
		c = color.RGBA{R: byte(ts), G: 0x00, B: 0x00}
	}

	println(temp, ts)
	visor.Alternate(c, color.RGBA{R: 0x00, G: 0x00, B: 0x00})
}

func tilt() {
	x, y, z, _ := accel.ReadAcceleration()
	println(x, y, z)

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
