package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"time"

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

func loadImageChannel(pathPicture string, images chan image.Image, errors chan error) {
	path, err := os.Open(pathPicture)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	image, err := png.Decode(path)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}

	if err == nil {
		images <- image
	} else {
		errors <- err
	}

}

func main() {

	screenshotImage := make(chan image.Image)
	frameImage := make(chan image.Image)
	errChannel := make(chan error)

	go loadImageChannel("test.png", screenshotImage, errChannel)
	go loadImageChannel("Frame-X.png", frameImage, errChannel)
	select {
	case screenshot := <-screenshotImage:
		size := screenshot.Bounds().Size()
		select {
		case frame := <-frameImage:
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
			concatDuration := time.Since(startTime)
			log.Print("Making image collage took " + concatDuration.String())
			defer third.Close()
		}
	case <-errChannel:
		log.Fatal("Specified directory with images inside does not exists")
	}

}
