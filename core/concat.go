package core

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/disintegration/imaging"
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

func StartConcat(center bool) {
	startTime := time.Now()
	screenshotImage := make(chan image.Image)
	frameImage := make(chan image.Image)
	errChannel := make(chan error)
	gradientChannel := make(chan *image.RGBA)
	go loadImageChannel("test.png", screenshotImage, errChannel)
	go loadImageChannel("Frame-X.png", frameImage, errChannel)

	select {
	case screenshot := <-screenshotImage:
		screenshotSize := screenshot.Bounds()
		fmt.Printf("Loaded Screenshot with size: %v x %v \n", screenshotSize.Size().X, screenshotSize.Size().Y)
		go CreateGradient(screenshotSize.Size().X, screenshotSize.Size().Y, gradientChannel)

		select {
		case frame := <-frameImage:

			//Create Image the Size of a Frame:
			frameSize := frame.Bounds()
			fmt.Printf("Loaded Frame with size: %v x %v \n", frameSize.Size().X, frameSize.Size().Y)

			output := image.NewRGBA(frameSize)
			offset := image.Pt(100, 90)
			//combine
			gradient := <-gradientChannel

			draw.Draw(output, frameSize.Add(offset), screenshot, image.ZP, draw.Src)
			draw.Draw(output, frameSize, frame, image.ZP, draw.Over)

			//make same size as Input:
			newImage := imaging.Resize(output, screenshotSize.Size().X, 0, imaging.NearestNeighbor)
			outputSize := newImage.Bounds()
			offsetOutput := image.Pt(0, 0)

			if center {
				//calculate middle:
				YOffset := (screenshotSize.Size().Y - outputSize.Size().Y) / 2
				offsetOutput = image.Pt(0, YOffset)
			} else {
				//put image at Bottom
				YOffset := (screenshotSize.Size().Y - outputSize.Size().Y)
				offsetOutput = image.Pt(0, YOffset)
			}

			draw.Draw(gradient, frameSize.Add(offsetOutput), newImage, image.ZP, draw.Over)

			third, err := os.Create("result.jpg")
			if err != nil {
				log.Fatalf("failed to create: %s", err)
			}

			png.Encode(third, gradient)
			concatDuration := time.Since(startTime)
			fmt.Printf("Output Frame with size: %v x %v \n", gradient.Bounds().Size().X, gradient.Bounds().Size().Y)
			log.Print("Making image collage took " + concatDuration.String())
			defer third.Close()
		}
	case <-errChannel:
		log.Fatal("Specified directory with images inside does not exists")
	}

}

type MyImage struct {
	value *image.RGBA
}

func (i *MyImage) Set(x, y int, c color.Color) {
	i.value.Set(x, y, c)
}

func (i *MyImage) ColorModel() color.Model {
	return i.value.ColorModel()
}

func (i *MyImage) Bounds() image.Rectangle {
	return i.value.Bounds()
}

func (i *MyImage) At(x, y int) color.Color {
	return i.value.At(x, y)
}
