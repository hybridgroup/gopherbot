# Learning TinyGo Using Gopherbot

Here are a series of activities to learn about TinyGo by using Gopherbot.

## Installation

If you have not done it yet, follow the installation instructions here:

LINK

## Code

### step0.go - Built-in LED

![Gopherbot](./images/step0.png)

This tests that you can compile and flash your Gopherbot with TinyGo code, by blinking the built-in LED.

- Plug your Circuit Playground Express into your computer's USB port.
- Press the "RESET" button on the board two times to get the Circuit Playground Express board ready to receive code. You have to do this every time you want to put new code on to the board.
- The Circuit Playground Express board will appear to your computer like a USB drive. Determine the path to the board, for example on Linux it will be something like `/media/[USERNAME]/[NAME OF THE BOARD]`.
- Build your TinyGo program to the board in `.uf2` format using this command:

```shell
tinygo build -size short -o /media/yourname/CPLAYBOOT/flash.uf2 -target=circuitplay-express ./learn/step0/main.go
```

- The Circuit Playground Express board should restart and then begin running your program. This program just causes the small LED labelled "D13" on the Circuit Playground Express board to start blinking on and off.

### step1.go - LED, Button A

![Gopherbot](./images/step1.png)

### step2.go - LED, Button A, Antenna LED, Button B

![Gopherbot](./images/step2.png)

### step3.go - LED, Button A, Antenna LED, Button B, Backpack LEDs

![Gopherbot](./images/step3.png)

### step4.go - LED, Button A, Antenna LED, Button B, Backpack LEDs, Sider switch

![Gopherbot](./images/step4.png)

### step5.go - LED, Button A, Antenna LED, Button B, Backpack LEDs, Sider switch, Visor LEDs

![Gopherbot](./images/step5.png)

### step6.go - LED, Button A, Antenna LED, Button B, Backpack LEDs, Sider switch, Thermistor

![Gopherbot](./images/step6.png)

### step7.go - LED, Button A, Antenna LED, Button B, Backpack LEDs, Sider switch, Thermistor, Accelerometer

![Gopherbot](./images/step7.png)

### step8.go - LED, Button A, Antenna LED, Button B, Backpack LEDs, Sider switch, Thermistor, Accelerometer, Buzzer

![Gopherbot](./images/step8.png)
