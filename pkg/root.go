package chigo

import "math"

type RGB struct {
	Red, Green, Blue int
}

func NewRGB(i float64) RGB {
	return RGB{
		int(math.Sin(0.1*i)*127 + 128),
		int(math.Sin(0.1*i+2*math.Pi/3)*127 + 128),
		int(math.Sin(0.1*i+4*math.Pi/3)*127 + 128),
	}
}
