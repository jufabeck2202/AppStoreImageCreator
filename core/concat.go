package core

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	"github.com/oliamb/cutter"
)

type (
	//Label is a struct
	Label struct {
		Text     string
		FontPath string
		FontType string
		Size     float64
		Color    image.Image
		DPI      float64
		Spacing  float64
		XPos     int
		YPos     int
	}
)


func loadImageChannel(pathPicture string, images chan image.Image, errors chan error) {
	path, err := os.Open(pathPicture)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	image, err := jpeg.Decode(path)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}

	if err == nil {
		images <- image
	} else {
		errors <- err
	}

}

func loadImagePNG(pathPicture string, images chan image.Image, errors chan error) {
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

func CutFrame(image image.Image) image.Image {
	croppedImg, err := cutter.Crop(image, cutter.Config{
		Width: image.Bounds().Size().X -240,
		Height: image.Bounds().Size().Y -120,
		Mode: cutter.Centered,
		Options: cutter.Copy,
	})
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	return croppedImg

}

func AddFrame(wg *sync.WaitGroup, inputImagePath string, userID string) {
	//wait group. End when finished.
	wg.Add(1)
	defer wg.Done()

	startTime := time.Now()
	inputFileName := filepath.Base(inputImagePath)
	screenshotImage := make(chan image.Image)
	errChannel := make(chan error)
	gradientChannel := make(chan *image.RGBA)
	frames := Frames{}

	go loadImageChannel(inputImagePath, screenshotImage, errChannel)

	select {
	case screenshot := <-screenshotImage:
		screenshotSize := screenshot.Bounds()
		fmt.Printf("Loaded Screenshot with size: %v x %v \n", screenshotSize.Size().X, screenshotSize.Size().Y)
		frameStruct := frames.GetForSize(screenshotSize.Size().X, screenshotSize.Size().Y)
		fmt.Printf("Found Frames for %s \n", frameStruct.Name)

		frameImage := make(chan image.Image)
		go loadImagePNG(filepath.Join("core", "frames", frameStruct.path), frameImage, errChannel)
		go CreateGradient(screenshotSize.Size().X, screenshotSize.Size().Y, gradientChannel)

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
			output := CutFrame(canvas)
			fmt.Printf("Loaded Frame with size: %v x %v \n", output.Bounds().Size().X, output.Bounds().Size().Y)



			//make same size as Input:
			newImage := imaging.Resize(output, screenshotSize.Size().X, 0, imaging.Lanczos)
			outputSize := newImage.Bounds()
			offsetOutput := image.Pt(0, 0)

			center := false
			if center {
				//calculate middle:
				YOffset := (screenshotSize.Size().Y - outputSize.Size().Y) / 2
				offsetOutput = image.Pt(0, YOffset)
			} else {
				//put image at Bottom
				YOffset := screenshotSize.Size().Y - outputSize.Size().Y
				offsetOutput = image.Pt(0, YOffset)
			}

			gradient := <-gradientChannel
			draw.Draw(gradient, frameSize.Add(offsetOutput), newImage, image.ZP, draw.Over)
			const S = 400
			dc := gg.NewContextForRGBA(gradient)
			dc.SetRGB(1, 1, 1)
			font, err := truetype.Parse(goregular.TTF)
			if err != nil {
				panic("")
			}
			face := truetype.NewFace(font, &truetype.Options{
				Size: 90,
			})
			dc.SetFontFace(face)
			text := "Download my App "
			dc.Stroke()
			dc.DrawStringWrapped(text, 0, 100, 0.0, 0.0, float64(outputSize.Size().X), 0, gg.AlignCenter)
			dc.SavePNG(path.Join("./Storage", "live", userID, inputFileName))


			concatDuration := time.Since(startTime)

			log.Printf("Output Frame with size: %v x %v \n", gradient.Bounds().Size().X, gradient.Bounds().Size().Y)
			log.Print("Making image collage took " + concatDuration.String())
		}
	case <-errChannel:
		log.Fatal("Specified directory with images inside does not exists")
	}

}

// GenerateBanner is a function that combine images and texts into one image
func GenerateBanner(labels []Label, background *image.RGBA) (*image.RGBA, error) {
	//create image's background

	//set the background color
	//add label(s)
	bgImg, err := addLabel(background, labels)
	if err != nil {
		return nil, err
	}

	return bgImg, nil
}

func addLabel(img *image.RGBA, labels []Label) (*image.RGBA, error) {
	//initialize the context
	c := freetype.NewContext()

	for _, label := range labels {
		//read font data
		fontBytes, err := ioutil.ReadFile(label.FontPath + label.FontType)
		if err != nil {
			return nil, err
		}
		f, err := freetype.ParseFont(fontBytes)
		if err != nil {
			return nil, err
		}
		//set label configuration
		c.SetDPI(label.DPI)
		c.SetFont(f)
		c.SetFontSize(label.Size)
		c.SetClip(img.Bounds())
		c.SetDst(img)
		c.SetSrc(label.Color)

		//positioning the label
		pt := freetype.Pt(label.XPos, label.YPos+int(c.PointToFixed(label.Size)>>6))
		// Calculate the widths and print to image

		//draw the label on image
		_, err = c.DrawString(label.Text, pt)
		if err != nil {
			log.Println(err)
			return img, nil
		}
		pt.Y += c.PointToFixed(label.Size * label.Spacing)
	}

	return img, nil
}
