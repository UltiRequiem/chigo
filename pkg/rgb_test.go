package chigo

import (
	"testing"
)

func TestNewRGB(t *testing.T) {
	rgb := NewRGB(0)

	if rgb.Red != 128 || rgb.Green != 237 || rgb.Blue != 18 {
		t.Errorf("NewRGB(0, 0, 0) failed")
	}
}
