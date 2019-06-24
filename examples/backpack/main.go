package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/hybridgroup/gopherbot"
	"tinygo.org/x/drivers/thermistor"
)

const (
	greenJet = iota
	redJet
	redRotate
	temperature
	xmasJet
)

const (
	forward = iota
	backward
)

var (
	antenna *machine.Pin

	sensor   *thermistor.Device
	backpack *gopherbot.Backpack

	pos = 0
	dir = 0
)

func main() {
	machine.InitADC()

	backpack = gopherbot.NewBackpack()

	s := thermistor.New(machine.TEMPSENSOR)
	s.Configure()
	sensor = &s

	left := machine.BUTTONA
	left.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	right := machine.BUTTONB
	right.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	mode := redJet

	for {
		if !right.Get() {
			mode++
			if mode > xmasJet {
				mode = greenJet
			}
		}

		if !left.Get() {
			mode--
			if mode < greenJet {
				mode = xmasJet
			}
		}

		switch mode {
		case greenJet:
			backpack.Green()
		case redJet:
			backpack.Red()
		case redRotate:
			backpack.Alternate(color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0x00})
		case temperature:
			showTemperature()
		case xmasJet:
			backpack.Xmas()
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func showTemperature() {
	temp, _ := sensor.ReadTemperature()

	ts := gopherbot.Rescale(temp, 0, 35000, 0, 255)
	println(temp, ts)

	var r, g, b byte
	switch {
	case temp < 12000:
		b = byte(ts)
	case temp > 12000 && temp < 25000:
		g = byte(ts)
	case temp > 25000:
		r = byte(ts)
	}
	backpack.SetColor(color.RGBA{R: r, G: g, B: b})
}
