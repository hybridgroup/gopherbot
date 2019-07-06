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
	xmasJet
)

const (
	forward = iota
	backward
)

func main() {
	backpack := gopherbot.NewBackpack()
	left := gopherbot.NewLeftButton()
	right := gopherbot.NewRightButton()

	mode := redJet

	for {
		if right.Pushed() {
			mode++
			if mode > xmasJet {
				mode = greenJet
			}
		}

		if left.Pushed() {
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
		case xmasJet:
			backpack.Xmas()
		}

		time.Sleep(200 * time.Millisecond)
	}
}
