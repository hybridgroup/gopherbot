# Gopherbot

Gopherbot is a robot gopher plushie that is programmable using TinyGo (https://tinygo.org)

[![GoDoc](https://godoc.org/github.com/hybridgroup/gopherbot?status.svg)](https://godoc.org/github.com/hybridgroup/gopherbot)

Uses an [Adafruit Circuit Playground Express](https://www.adafruit.com/product/3333) with a 3D printed helmet and backpack.

Here is a TinyGo program that blinks all of the various built-in LEDs all at the same time when the slider switch is in the "on" position.

```go
package main

import (
	"time"

	"github.com/hybridgroup/gopherbot"
)

func main() {
	led := gopherbot.StatusLED()
	antenna := gopherbot.Antenna()
	visor := gopherbot.Visor()
	backpack := gopherbot.Backpack()

	slider := gopherbot.Slider()

	for {
		if slider.IsOn() {
			led.On()
			antenna.On()
			visor.Blue()
			backpack.Blue()
			time.Sleep(500 * time.Millisecond)

			led.Off()
			antenna.Off()
			visor.Off()
			backpack.Off()
			time.Sleep(500 * time.Millisecond)
		}
	}
}
```

## Assembling Your Kit

Did you receive a Gopherbot kit? Awesome! For assembly instructions, go here:

[Gopherbot Assembly Instructions](./assembly/README.md)

The 3D designs for the helmet, ears, and backpack are created by Damen Evans under the Creative Commons - Attribution - Non-Commercial license. The files are located here:

https://www.thingiverse.com/thing:3761937

## Installation

To put code on Gopherbot, you need to install some software on your own machine.

### Go 1.18

If you have not installed Go 1.18 on your computer already, you can download it here:

https://golang.org/dl/

### TinyGo

Follow the instructions here to install TinyGo:

https://tinygo.org/getting-started/

### Gopherbot code

Lastly, get the code from this repository:

```
git clone https://github.com/hybridgroup/gopherbot.git
cd gopherbot
```

OK great, we're ready to write our first TinyGo program and put it on Gopherbot.

## Hello, Gopherbot

Gopherbot uses an [Adafruit Circuit Playground Express](https://www.adafruit.com/product/3333) as its "brain". To put new code on the Circuit Playground Express you can copy a file in the correct format from your computer to the board using a USB connection without having to install any extra flashing software. This is because it comes with a "bootloader" named UF2 already installed, that lets to do the flashing. Here is more information about the [UF2 bootloader](https://github.com/Microsoft/uf2) if you are interested.

Here is what to do:

- Plug your Circuit Playground Express into your computer's USB port.
- Build and flash your TinyGo program to the board in `.uf2` format using this command:

```shell
tinygo flash -target=gopherbot examples/blinky1
```

- The Circuit Playground Express board should restart and then begin running your program. This program just causes the small LED labelled "D13" on the Circuit Playground Express board to start blinking on and off.

Now you are ready to try something a little more flashy.

- Build and flash the demo TinyGo program to the board using this command:

```shell
tinygo flash -target=circuitplay-express ./examples/blink/
```

Now THAT is a blink!

## What To Do Next?

If you want to load one of the example programs on Gopherbot, check out our [examples located here](./examples).

For a series of activities to learn programming TinyGo using Gopherbot, check out the [learn folder](./learn/README.md).

Have fun!

## Gopherbot 2

The Gopherbot 2 is the same as the Gopherbot but uses a Circuit Playground Bluefruit board, which has a Nordic Semiconductors nrf52840 processor with built-in Bluetooth radio.

To flash code on the Gopherbot 2, use the same commands but with the `target=gopherbot2`.

```shell
tinygo flash -target=gopherbot2 ./examples/blink/
```
