package core

import (
	"image"
	"image/color"
	"strconv"
)

var (
	colorB = [3]float64{248, 54, 0}
	colorA = [3]float64{254, 140, 0}
)

type Hex string

type RGB struct {
	Red   float64
	Green float64
	Blue  float64
}

func (h Hex) toRGB() (RGB, error) {
	return Hex2RGB(h)
}

func Hex2RGB(hex Hex) (RGB, error) {
	var rgb RGB
	values, err := strconv.ParseUint(string(hex), 16, 32)

	if err != nil {
		return RGB{}, err
	}

	rgb = RGB{
		Red:   float64(uint8(values >> 16)),
		Green: float64(uint8((values >> 8) & 0xFF)),
		Blue:  float64(uint8(values & 0xFF)),
	}

	return rgb, nil
}

func linearGradient(x, y, maxSize float64, color1, color2 string) (uint8, uint8, uint8) {
	//convert string to hex RGB
	//remove #
	var color1hex Hex = Hex(trimLeftChar(color1))
	var color2hex Hex = Hex(trimLeftChar(color2))
	rgb1, _ := Hex2RGB(color1hex)
	rgb2, _ := Hex2RGB(color2hex)

	d := x / maxSize
	r := rgb1.Red + d*(rgb2.Red-rgb1.Red)
	g := rgb1.Green + d*(rgb2.Green-rgb1.Green)
	b := rgb1.Blue + d*(rgb2.Blue-rgb1.Blue)
	return uint8(r), uint8(g), uint8(b)
}

func CreateGradient(width, height int, color1, color2 string, images chan *image.RGBA) {

	var w, h int = width, height
	gradient := image.NewRGBA(image.Rect(0, 0, w, h)) //*NRGBA (image.Image interface)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b := linearGradient(float64(x), float64(y), float64(height), color1, color2)
			c := color.RGBA{
				r,
				g,
				b,
				255,
			}
			gradient.Set(x, y, c)
		}
	}
	images <- gradient
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}
