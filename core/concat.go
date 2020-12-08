package core

import (
	"image"
	"image/draw"
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

func StartConcat() {

	screenshotImage := make(chan image.Image)
	frameImage := make(chan image.Image)
	errChannel := make(chan error)
	gradientChannel := make(chan image.Image)
	go loadImageChannel("test.png", screenshotImage, errChannel)
	go loadImageChannel("Frame-X.png", frameImage, errChannel)
	go CreateGradient(1325, 2616, gradientChannel)

	select {
	case frame := <-frameImage:
		select {
		case screenshot := <-screenshotImage:
			size := screenshot.Bounds().Size()
			//newImage := resize.Resize(160, 0, original_image, resize.Lanczos3)

			//Create Image the Size of a Frame:
			frameSize := frame.Bounds()
			output := image.NewRGBA(frameSize)
			offset := image.Pt(100, 90)
			//combine
			gradient := <-gradientChannel
			draw.Draw(output, frameSize, gradient, image.ZP, draw.Over)
			draw.Draw(output, frame.Bounds().Add(offset), screenshot, image.ZP, draw.Src)
			draw.Draw(output, frameSize, frame, image.ZP, draw.Over)

			//make same size as Input:
			newImage := resize.Resize(uint(size.X), uint(size.Y), output, resize.Lanczos3)

			third, err := os.Create("result.jpg")
			if err != nil {
				log.Fatalf("failed to create: %s", err)
			}
			startTime := time.Now()
			png.Encode(third, newImage)
			concatDuration := time.Since(startTime)
			log.Print("Making image collage took " + concatDuration.String())
			defer third.Close()
		}
	case <-errChannel:
		log.Fatal("Specified directory with images inside does not exists")
	}

}
