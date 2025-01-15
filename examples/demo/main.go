package main

import (
	"image/color"
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
	accel   *gopherbot.AccelerometerDevice
	visor   *gopherbot.VisorDevice
	speaker *gopherbot.SpeakerDevice

	pos = 0
)

func main() {
	accel = gopherbot.Accelerometer()
	visor = gopherbot.Visor()
	speaker = gopherbot.Speaker()

	antenna := gopherbot.Antenna()
	backpack := gopherbot.Backpack()

	left := gopherbot.LeftButton()
	right := gopherbot.RightButton()

	mode := redVisor

	go antenna.Blink()

	for {
		if right.Pushed() {
			mode++
			if mode > xmasVisor {
				mode = greenVisor
			}
		}

		if left.Pushed() {
			mode--
			if mode < greenVisor {
				mode = xmasVisor
			}
		}

		switch mode {
		case greenVisor:
			visor.Green()
			if gopherbot.Slider().IsOn() {
				backpack.Alternate(color.RGBA{R: 0x00, G: 0xff, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0xff})
			} else {
				backpack.Off()
			}
		case redVisor:
			visor.Red()
			if gopherbot.Slider().IsOn() {
				backpack.Red()
			} else {
				backpack.Off()
			}
		case cylonVisor:
			visor.Cylon()
			if gopherbot.Slider().IsOn() {
				backpack.Alternate(color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0x00})
			} else {
				backpack.Off()
			}
		case tiltVisor:
			tilt()
		case xmasVisor:
			visor.Xmas()
			if gopherbot.Slider().IsOn() {
				backpack.Xmas()
			} else {
				backpack.Off()
			}
		}

		time.Sleep(200 * time.Millisecond)
	}
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
			speaker.Bleep()
		}
	case x < -300000:
		pos--
		if pos < 0 {
			pos = 0
			speaker.Bleep()
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
