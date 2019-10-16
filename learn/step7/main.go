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
)

const (
	forward = iota
	backward
)

var (
	visor  *gopherbot.VisorDevice
	thermo *gopherbot.ThermometerDevice
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

	mode := roomTemperature

	go antenna.Blink()

	for {
		statusOn := false

		if left.Pushed() {
			statusOn = true
			mode--
			if mode < greenJet {
				mode = roomTemperature
			}
		}

		if right.Pushed() {
			statusOn = true
			mode++
			if mode > roomTemperature {
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
