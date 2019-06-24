package gopherbot

const (
	BackpackLEDCount = 10
	VisorLEDCount    = 12
)

// Rescale performs a direct linear rescaling of an integer from one scale to another.
//
// For example:
//
//		val := gopherbot.Rescale(25, 0, 100, 0, 10)
//
func Rescale(input, fromMin, fromMax, toMin, toMax int32) int32 {
	return (input-fromMin)*(toMax-toMin)/(fromMax-fromMin) + toMin
}
