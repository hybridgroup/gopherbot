# Learning TinyGo Using Gopherbot

Here are a series of activities to learn about TinyGo by using Gopherbot.

[![GoDoc](https://godoc.org/github.com/hybridgroup/gopherbot?status.svg)](https://godoc.org/github.com/hybridgroup/gopherbot)

## Installation

If you have not done it yet, follow the installation instructions here:

https://github.com/hybridgroup/gopherbot#installation

## Code

### step0.go - Built-in LED

This tests that you can compile and flash your Gopherbot with TinyGo code, by blinking the built-in LED.

- Plug your Circuit Playground Express into your computer's USB port.
- Build your TinyGo program to the board in `.uf2` format using this command:

```shell
tinygo flash -target=circuitplay-express ./learn/step0/main.go
```

- The Circuit Playground Express board should restart and then begin running your program. This program just causes the small LED labelled "D13" on the Circuit Playground Express board to start blinking on and off.

### step1.go - LED, Button A

Now, instead of just using the low-level TinyGo functions, we are going to use the more advanced "Gopherbot" API to program our furrie robotic friend. Gopherbot has more pre-defined behaviors, but of course all of these are built on top of the lower-level TinyGo "machine" API.

Want to see complete documentation of the Gopherbot API? Check out our GoDocs here at:
https://godoc.org/github.com/hybridgroup/gopherbot

First thing we are going to do is to program the left button on the backpack. Let's make it so when the button is pressed, the small built-in LED on the board turns on. Also, the code will make it so that when the button is released that the LED turns off.

```shell
tinygo flash -target=circuitplay-express ./learn/step1/main.go
```

### step2.go - LED, Button A, Button B

Next, let's do the same thing for the right button that we did for the left one. When the button is pressed, the small built-in LED on the board turns on. Also, the code will make it so that when the button is released that the LED turns off.

```shell
tinygo flash -target=circuitplay-express ./learn/step2/main.go
```

### step3.go - LED, Button A, Button B, Backpack LEDs

Now we will get brightly colored, by adding the ring of Red Green Blue (RGB) LEDs that are on the Circuit Playground Express board on the `Backpack`.

The Gopherbot `Backpack` has several built-in functions for these LEDs to set colors and patterns. We will add three of these color settings to the program.

We will also use the buttons to select which of the `Backpack` colors to display. Pushing the right or left button will scroll thru these three color presets.

```shell
tinygo flash -target=circuitplay-express ./learn/step3/main.go
```

### step4.go - LED, Button A, Button B, Backpack LEDs, Slider switch

The RGB LEDs are really bright! Let's add the ability to turn them on and off using the built-in slider switch.

```shell
tinygo flash -target=circuitplay-express ./learn/step4/main.go
```

### step5.go - LED, Button A, Button B, Backpack LEDs, Slider switch, Antenna LED

Now we need to get the LED on the end of the `Antenna` blinking. We can do this so that it automatically blinks over and over by using a goroutine to call the `Antenna.Blink()` function.

```shell
tinygo flash -target=circuitplay-express ./learn/step5/main.go
```

### step6.go - LED, Button A, Button B, Backpack LEDs, Slider switch, Antenna LED, Visor LEDs

Time to fully activate Gopherbot by controlling the RGB LEDs that are in the helmet's `Visor`.

Similarly to the `Backpack` does, the `Visor` has several built-in functions for the LEDs to set colors and patterns. We will add three of these color settings to the program.

We will also use the buttons to select which of the `Visor` colors to display. Pushing the right or left button will scroll thru these three color presets at the same time that it chooses the presets for the `Backpack`.

One difference will be that we will program Gopherbot so that the `Visor` stays on regardless of state of the `Slider` switch.

```shell
tinygo flash -target=circuitplay-express ./learn/step6/main.go
```

### step7.go - LED, Button A, Button B, Backpack LEDs, Slider switch, Antenna LED, Visor LEDs, Thermistor

Now that we have all the LEDs working all at the same time, let's explore one of the sensors that is built-in. The Circuit Playground Express board has an onboard `Thermometer` which is actually a kind of analog temperature sensor known as a thermistor.

We will add to our program an additional `mode` which will read the current temperature, and then modify the `Visor` depending on the reading from the `Thermometer`.

```shell
tinygo flash -target=circuitplay-express ./learn/step7/main.go
```

### step8.go - LED, Button A, Button B, Backpack LEDs, Slider switch, Antenna LED, Visor LEDs, Thermistor, Accelerometer

Let try another of the built-in sensors on the Circuit Playground Express board. It has an onboard `Accelerometer` which is a sensor that can detect motion on any of 3 different axes.

We will add to our program another `mode` which will read the `Accelerometer`, and then based on the tilt position of the Gopherbot, modify the `Visor` to make it act like a "level".

```shell
tinygo flash -target=circuitplay-express ./learn/step8/main.go
```

### step9.go - LED, Button A, Button B, Backpack LEDs, Slider switch, Antenna LED, Visor LEDs, Thermistor, Accelerometer, Buzzer

Out final modification to the program will add the built-in `Speaker` to make a bleeping sound when the `tilt()` function detects that the Gopherbot is laying entirely on its side.

```shell
tinygo flash -target=circuitplay-express ./learn/step9/main.go
```
