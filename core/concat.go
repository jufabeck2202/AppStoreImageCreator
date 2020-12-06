package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func loadImage(dirPath string) (image.Image, image.Point) {
	path, err := os.Open(dirPath)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	image, err := png.Decode(path)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer path.Close()
	return image, image.Bounds().Size()
}

func main() {
	screenshot, size := loadImage("test.png")

	frame, _ := loadImage("Frame-X.png")
	//newImage := resize.Resize(160, 0, original_image, resize.Lanczos3)

	//Create Image the Size of a Frame:
	frameSize := frame.Bounds()
	output := image.NewRGBA(frameSize)
	offset := image.Pt(100, 90)
	//combine
	draw.Draw(output, frame.Bounds().Add(offset), screenshot, image.ZP, draw.Src)
	draw.Draw(output, frameSize, frame, image.ZP, draw.Over)

	//make same size as Input:
	newImage := resize.Resize(uint(size.X), uint(size.Y), output, resize.Lanczos3)

	third, err := os.Create("result.jpg")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, newImage, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()
}
