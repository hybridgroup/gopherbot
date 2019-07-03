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
