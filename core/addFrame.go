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

func CutFrame(image image.Image, frame *DeviceFrame) image.Image {
	croppedImg, err := cutter.Crop(image, cutter.Config{
		Width:   image.Bounds().Size().X - frame.XBorderInv,
		Height:  image.Bounds().Size().Y - frame.YBorderInv,
		Mode:    cutter.Centered,
		Options: cutter.Copy,
	})
	if err != nil {
		log.Printf("failed to decode: %s", err)
	}
	return croppedImg
}

func AddFrame(task *addFrameTask, errChan chan error, returnFrame chan image.Image) {
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

			var output image.Image

			//remove border
			if frameStruct.HasBorder {
				output = CutFrame(canvas, &frameStruct)
				fmt.Printf("Removed Boarder: %v x %v \n", output.Bounds().Size().X, output.Bounds().Size().Y)
			} else {
				output = canvas
			}

			var resizedImage image.Image

			if task.resizeToOriginal {
				//resize to screenshot size
				resizedImage = imaging.Resize(output, screenshotSize.Size().X, 0, imaging.Lanczos)
			} else {
				resizedImage = output
			}
			//make same size as Input:
			outputSize := resizedImage.Bounds()
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
			background := <-backgroundChannel
			fmt.Printf("Got Gradient with with size: %v x %v \n", background.Bounds().Size().X, background.Bounds().Size().Y)
			draw.Draw(background, frameSize.Add(offsetOutput), resizedImage, image.ZP, draw.Over)

			//add text
			if task.hasText() {
				dc := gg.NewContextForRGBA(background)
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
			returnFrame <- background

		}
	case err := <-errChannel:
		errChan <- err
	}

}
