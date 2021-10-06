package chigo

import "math"

func GetRGB(i int) (red int, green int, blue int) {
	red = int(math.Sin(0.1*float64(i)+0)*127 + 128)
	green = int(math.Sin(0.1*float64(i)+2*math.Pi/3)*127 + 128)
	blue = int(math.Sin(0.1*float64(i)+4*math.Pi/3)*127 + 128)
	return
}
