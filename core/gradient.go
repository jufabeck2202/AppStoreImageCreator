package core

import (
	"image"
	"image/color"
)

var (
	colorB = [3]float64{248, 54, 0}
	colorA = [3]float64{254, 140, 0}
)

// var (
// 	width  = 256
// 	height = 256
// 	max    = float64(width)
// )

func linearGradient(x, y, maxSize float64) (uint8, uint8, uint8) {
	d := x / maxSize
	r := colorA[0] + d*(colorB[0]-colorA[0])
	g := colorA[1] + d*(colorB[1]-colorA[1])
	b := colorA[2] + d*(colorB[2]-colorA[2])
	return uint8(r), uint8(g), uint8(b)
}

func CreateGradient(width, height int, images chan image.Image) {

	var w, h int = width, height
	image := image.NewRGBA(image.Rect(0, 0, w, h)) //*NRGBA (image.Image interface)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			r, g, b := linearGradient(float64(x), float64(y), float64(height))
			c := color.RGBA{

				r,
				g,
				b,
				255,
			}
			image.Set(x, y, c)
		}
	}
	images <- image

}
