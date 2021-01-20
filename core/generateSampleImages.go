package core

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func GenerateImages() {
	var wg sync.WaitGroup
	frames := Frames{}.get()
	path := "./core/frames/samples"
	for _, frame := range frames {
		wg.Add(1)
		go sampleImage(frame, path, &wg)
	}
	wg.Wait()
}

func sampleImage(frame DeviceFrame, path string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Generating:", frame.Name)

	dc := gg.NewContext(frame.screenshotWidth, frame.screenshotHeight)
	for i := 0.0; i < 10; i++ {

		dc.DrawRectangle(0+(i*20), 0+(i*20), float64(frame.screenshotWidth)-(i*40), float64(frame.screenshotHeight)-(i*40))
		dc.SetRGB(float64(rand.Intn(255)), float64(rand.Intn(255)), float64(rand.Intn(255)))
		dc.Fill()
	}

	err := dc.SavePNG(filepath.Join(path, strings.ReplaceAll(frame.Name, " ", "_")+".png"))
	if err != nil {
		log.Printf("failed to decode: %s", err)
	}

}

func GenerateTestFrames() {
	files, _ := FilePathWalkDir(filepath.Join("./core/frames/samples"))
	frames := make(chan ReturnFrame, len(files))
	var wg sync.WaitGroup

	for _, path := range files {
		wg.Add(1)
		go generateTestFrame(path, frames, &wg)
	}

	wg.Wait()
	var allFrames []image.Image
	for {
		select {
		case elem := <-frames:

			f, err := os.Create(elem.path)
			if err != nil {
				// Handle error
			}
			defer f.Close()

			// Specify the quality, between 0-100
			// Higher is better
			opt := jpeg.Options{
				Quality: 100,
			}
			err = jpeg.Encode(f, elem.Frame, &opt)
			if err != nil {
				fmt.Println(err)
			}
			allFrames = append(allFrames, elem.Frame)

		default:
			fmt.Println("No value ready, moving on.")
			CreateTestWallpaper(allFrames)
			return
		}
	}

}

func generateTestFrame(path string, newFrame chan ReturnFrame, wg *sync.WaitGroup) {
	defer wg.Done()
	frame := make(chan image.Image)
	error := make(chan error)
	task := CreateNewFrameTask(path,"#FF0000","#00FF00","Iphone",false)
	go AddFrame(task, error, frame)
	select {
	case frame2 := <-frame:
		newFrame <- ReturnFrame{
			Frame: frame2,
			path:  path,
		}
	case err := <-error:
		log.Fatal(err.Error())
	}

}
