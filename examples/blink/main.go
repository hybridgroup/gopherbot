package main

import (
	"time"

	"github.com/hybridgroup/gopherbot"
)

func main() {
	visor := gopherbot.NewVisor()

	for {
		visor.Green()
		time.Sleep(time.Millisecond * 500)

		visor.Clear()
		time.Sleep(time.Millisecond * 500)
	}
}
