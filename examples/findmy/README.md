# Find My Gopherbot

Your Gopherbot is now a FindMy beacon.

Requires CircuitPlayground Bluefruit.

Uses "Go Haystack" see [https://github.com/hybridgroup/go-haystack](https://github.com/hybridgroup/go-haystack)

## Flashing

```shell
tinygo flash -target circuitplay-bluefruit -ldflags="-X main.AdvertisingKey='YOURKEYHERE'" ./examples/demo
```
