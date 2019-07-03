# Gopherbot

A programmable robot gopher plushie.

## Assembling Your Kit

Did you receive a Gopherbot kit? Awesome! For assembly instructions, go here:

[Assembly Instructions](./assembly/README.md)

## Putting code on Gopherbot

To put code on Gopherbot, you need to install some software on your own machine.

### Go 1.12

If you have not installed Go 1.12 on your computer already, you can download it here:

https://golang.org/dl/

### TinyGo

Follow the instructions here to install TinyGo:

https://tinygo.org/getting-started/

### TinyGo drivers

To install the various drivers and other code dependencies run these commands:

```
go get -u tinygo.org/x/drivers
go get -u github.com/eclipse/paho.mqtt.golang
```

### TinyGo drivers

Lastly, get the code from this repository:

```
git clone https://github.com/hybridgroup/gopherbot.git
cd gopherbot
```

OK great, we're ready to write our first TinyGo program and put it on Gopherbot.

## Hello, Gopherbot

The Circuit Playground Express comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

- Plug your Circuit Playground Express into your computer's USB port.
- Press the "RESET" button on the board two times to get the Circuit Playground Express board ready to receive code. You have to do this every time you want to put new code on to the board.
- The Circuit Playground Express board will appear to your computer like a USB drive. Determine the path to the board, for example on Linux it will be something like `/media/[USERNAME]/[NAME OF THE BOARD]`.
- Build your TinyGo program to the board in `.uf2` format using this command:

```shell
tinygo build -size short -o /media/yourname/CPLAYBOOT/flash.uf2 -target=circuitplay-express examples/blinky1
```

- The Circuit Playground Express board should restart and then begin running your program. This program just causes the small LED labelled "D13" on the Circuit Playground Express board to start blinking on and off.

Now you are ready to try something more flashy.

- Once again, press the "RESET" button on the board two times to get the Circuit Playground Express board ready to receive code.
- Build the demo TinyGo program to the board in format using this command:

```shell
tinygo build -size short -o /media/yourname/CPLAYBOOT/flash.uf2 -target=circuitplay-express ./examples/blink/main.go
```

Now THAT is a blink!

## What To Do Next?

If you want to load one of the example programs on Gopherbot, check out our [examples](./examples/README.md).

For a series of programming activities to learn TinyGo using Gopherbot, check out the [learn folder](./learn/README.md).
