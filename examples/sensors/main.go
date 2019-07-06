package main

import (
	"image/color"
	"time"

	"github.com/hybridgroup/gopherbot"
)

const (
	temperature = iota
	lightlevel
)

const (
	forward = iota
	backward
)

var (
	thermo *gopherbot.ThermometerDevice
	photo  *gopherbot.LightMeterDevice

	backpack *gopherbot.BackpackDevice

	pos = 0
	dir = 0
)

func main() {
	backpack = gopherbot.Backpack()

	thermo = gopherbot.Thermometer()
	photo = gopherbot.LightMeter()

	left := gopherbot.LeftButton()
	right := gopherbot.RightButton()

	mode := temperature

	for {
		if right.Pushed() {
			mode++
			if mode > lightlevel {
				mode = temperature
			}
		}

		if left.Pushed() {
			mode--
			if mode < temperature {
				mode = lightlevel
			}
		}

		switch mode {
		case temperature:
			showTemperature()
		case lightlevel:
			showLumens()
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func showTemperature() {
	temp, _ := thermo.ReadTemperature()

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

func showLumens() {
	lumens := photo.Get()

	lvl := gopherbot.Rescale(int32(lumens), 0, 35000, 0, 255)
	println(lumens, lvl)

	var r, g, b byte
	switch {
	case lumens < 12000:
		b = byte(lvl)
	case lumens > 12000 && lumens < 25000:
		g = byte(lvl)
	case lumens > 25000:
		r = byte(lvl)
	}
	backpack.SetColor(color.RGBA{R: r, G: g, B: b})
}
