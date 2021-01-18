package core

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/oliamb/cutter"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/draw"
	"log"
	"path/filepath"
)

const framePath = "./core/frames"

func CutFrame(image image.Image) image.Image {
	croppedImg, err := cutter.Crop(image, cutter.Config{
		Width:   image.Bounds().Size().X - 240,
		Height:  image.Bounds().Size().Y - 120,
		Mode:    cutter.Centered,
		Options: cutter.Copy,
	})
	if err != nil {
		log.Printf("failed to decode: %s", err)
	}
	return croppedImg
}

func AddFrame(task *addFrameTask, errChan chan error, returnFrame chan image.Image) {
	//wait group. End when finished.
	screenshotImage := make(chan image.Image)
	errChannel := make(chan error)
	backgroundChannel := make(chan *image.RGBA)
	frames := Frames{}

	go LoadImageChannel(task.inputImagePath, screenshotImage, errChannel)

	select {
	case screenshot := <-screenshotImage:
		screenshotSize := screenshot.Bounds()

		fmt.Printf("Loaded Screenshot with size: %v x %v \n", screenshotSize.Size().X, screenshotSize.Size().Y)
		frameStruct := frames.GetForSize(screenshotSize.Size().X, screenshotSize.Size().Y)
		fmt.Printf("Found Frames for %s \n", frameStruct.Name)

		frameImage := make(chan image.Image)
		go LoadImageChannel(filepath.Join(framePath, frameStruct.path), frameImage, errChannel)

		if task.hasGradient() {
			go CreateGradient(screenshotSize.Size().X, screenshotSize.Size().Y, task.hexColor1, task.hexColor2, backgroundChannel)
		} else {
			go SingleColorBackground(screenshotSize.Size().X, screenshotSize.Size().Y, task.hexColor1, backgroundChannel)

		}

		select {
		case frame := <-frameImage:

			//Create Image the Size of a Frame:
			frameSize := frame.Bounds()
			fmt.Printf("Resized Frame with size: %v x %v \n", frameSize.Size().X, frameSize.Size().Y)

			canvas := image.NewRGBA(frameSize)
			offset := image.Pt(frameStruct.xOffset, frameStruct.YOffset)
			//combine

			//draw screenshot on Output
			draw.Draw(canvas, frameSize.Add(offset), screenshot, image.ZP, draw.Src)
			//draw frame on output
			draw.Draw(canvas, frameSize, frame, image.ZP, draw.Over)

			//resize frame
			//output := CutFrame(canvas)
			output := canvas
			//fmt.Printf("Resized new Frame with size: %v x %v \n", output.Bounds().Size().X, output.Bounds().Size().Y)

			//make same size as Input:
			newImage := imaging.Resize(output, screenshotSize.Size().X, 0, imaging.Lanczos)
			outputSize := newImage.Bounds()
			offsetOutput := image.Pt(0, 0)

			if task.hasText() {
				//put image at Bottom
				YOffset := screenshotSize.Size().Y - outputSize.Size().Y
				offsetOutput = image.Pt(0, YOffset)
			} else {

				//calculate middle:
				YOffset := (screenshotSize.Size().Y - outputSize.Size().Y) / 2
				offsetOutput = image.Pt(0, YOffset)
			}
			//fetch gradient
			gradient := <-backgroundChannel
			fmt.Printf("Got Gradient with with size: %v x %v \n", gradient.Bounds().Size().X, gradient.Bounds().Size().Y)
			draw.Draw(gradient, frameSize.Add(offsetOutput), newImage, image.ZP, draw.Over)

			if task.hasText() {
				dc := gg.NewContextForRGBA(gradient)
				dc.SetRGB(1, 1, 1)

				font, err := truetype.Parse(goregular.TTF)
				if err != nil {
					errChan <- err
				}
				face := truetype.NewFace(font, &truetype.Options{
					Size: 90,
				})
				dc.SetFontFace(face)
				dc.Stroke()
				dc.DrawStringWrapped(task.heading, 0, 100, 0.0, 0.0, float64(outputSize.Size().X), 0, gg.AlignCenter)
				returnFrame <- dc.Image()
			}
			returnFrame <- gradient

		}
	case err := <-errChannel:
		errChan <- err
	}

}
