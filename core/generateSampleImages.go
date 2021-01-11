package core

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/jufabeck2202/AppStoreImageCreator/server"
	"image"
	"log"
	"math/rand"
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
	for i:= 0.0; i < 10; i++ {

		dc.DrawRectangle(0+(i*20),0+(i*20), float64(frame.screenshotWidth)-(i*40), float64(frame.screenshotHeight)-(i*40))
		dc.SetRGB(float64(rand.Intn(255)),float64(rand.Intn(255)),float64(rand.Intn(255)))
		dc.Fill()
	}

	err := dc.SavePNG(filepath.Join(path, strings.ReplaceAll(frame.Name, " ", "_")+".png"))
	if err != nil {
		log.Printf("failed to decode: %s", err)
	}

}

func GenerateTestFrames() {
	files, _ := server.FilePathWalkDir(filepath.Join("./core/frames/samples"))
	var wg sync.WaitGroup
	frames := make(chan string, len(files))

	for _, path := range files {
		go generateTestFrame(path, frames, &wg)
	}
	wg.Wait()
}

func generateTestFrame(path string, returnFrame chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	frame := make(chan image.Image)
	error := make(chan error)

	go AddFrameNew(path,"", "","Iphone",false,true,error,frame)


}
