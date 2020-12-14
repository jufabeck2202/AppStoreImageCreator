package core

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/disintegration/imaging"
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
			newImage := imaging.Resize(output, screenshotSize.Size().X, 0, imaging.Lanczos)
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
			labels := []Label{
				Label{
					FontPath: "../../golang/freetype/testdata/",
					Size:     48,
					FontType: "luximr.ttf" ,
					Color:    image.Black,
					DPI:      72,
					Spacing:  1.5,
					Text:     "Download my beautiful App",
					XPos:     10,
					YPos:     0,
				},
				Label{
					FontPath: "./core/fonts/",
					Size:     80,
					FontType: "SFProDisplay-Regular.ttf",
					Color:    image.Black,
					DPI:      72,
					Spacing:  1.5,
					Text:     "The Real App SUCKS!",
					XPos:     10,
					YPos:     50,
				},
			}


			const S = 400
			dc := gg.NewContextForRGBA(gradient)
			dc.SetRGB(0, 0, 0)
			font, err := truetype.Parse(goregular.TTF)
			if err != nil {
				panic("")
			}
			face := truetype.NewFace(font, &truetype.Options{
				Size: 40,
			})
			dc.SetFontFace(face)
			text := "Hello, world! This text is a bit centered. help my i will call your mother if you are not a good boy. Yeaaa goood booooy "
 			dc.Stroke()
			dc.DrawStringWrapped(text, 0, 100, 0.0, 0.0, float64(outputSize.Size().X),  0, gg.AlignCenter)
			//dc.Image()
			dc.SavePNG("out.png")

			textImage, error := GenerateBanner(labels, gradient)
			if error != nil {
				log.Fatalf("failed to create: %s", error)
			}

			third, error := os.Create("result.jpg")
			if error != nil {
				log.Fatalf("failed to create: %s", error)
			}

			//png.Encode(third, gradient)
			png.Encode(third, textImage)
			concatDuration := time.Since(startTime)
			fmt.Printf("Output Frame with size: %v x %v \n", gradient.Bounds().Size().X, gradient.Bounds().Size().Y)
			log.Print("Making image collage took " + concatDuration.String())
			defer third.Close()
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
