package core

import (
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/imview"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sort"
)

func Width(i image.Image) int {
	return i.Bounds().Max.X - i.Bounds().Min.X
}

func Height(i image.Image) int {
	return i.Bounds().Max.Y - i.Bounds().Min.Y
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

type Size struct {
	width  uint
	height uint
}

type ImageShape string

type ImagePositionAndSize struct {
	sp   image.Point
	size Size
}

const (
	RectangleShape ImageShape = "Rectangle"
)

func (bgImg *MyImage) drawRaw(innerImg image.Image, sp image.Point, width uint, height uint) {
	resizedImg := imaging.Resize(innerImg, int(width), int(height), imaging.Lanczos)
	w := int(Width(resizedImg))
	h := int(Height(resizedImg))
	draw.Draw(bgImg, image.Rectangle{sp, image.Point{sp.X + w, sp.Y + h}}, resizedImg, image.ZP, draw.Src)
}

func makeImageCollage(desiredWidth int, numberOfRows int, images ...image.Image) *MyImage {

	sort.Slice(images, func(i, j int) bool {
		return Height(images[i]) > Height(images[j])
	})

	numberOfColumns := len(images) / numberOfRows
	imagesMatrix := make([][]image.Image, numberOfRows)

	numberOfColumnsAdded := 0
	maxNumberOfColumns := 0
	for idx := 0; idx < numberOfRows; idx++ {
		columnsInRow := numberOfColumns
		if len(images)%numberOfRows > 0 && (numberOfRows-idx)*numberOfColumns < len(images)-numberOfColumnsAdded {
			columnsInRow++
		}

		if columnsInRow > maxNumberOfColumns {
			maxNumberOfColumns = columnsInRow
		}

		imagesMatrix[idx] = images[numberOfColumnsAdded : numberOfColumnsAdded+columnsInRow]
		numberOfColumnsAdded += columnsInRow
	}

	maxWidth := uint(0)
	imagesSize := make([][]Size, numberOfRows)
	for row := 0; row < numberOfRows; row++ {
		imagesSize[row] = make([]Size, len(imagesMatrix[row]))

		calculatedWidth := math.Floor(float64(desiredWidth) / float64(len(imagesMatrix[row])))

		rowWidth := uint(0)
		rowHeight := uint(0)
		for col := 0; col < len(imagesMatrix[row]); col++ {
			originalWidth := float64(Width(imagesMatrix[row][col]))
			originalHeight := float64(Height(imagesMatrix[row][col]))
			resizeFactor := calculatedWidth / originalWidth

			w := uint(originalWidth * resizeFactor)
			h := uint(originalHeight * resizeFactor)
			imagesSize[row][col] = Size{w, h}

			rowWidth += w
			rowHeight += h

		}

		if rowWidth > maxWidth {
			maxWidth = rowWidth
		}
	}

	maxHeight := uint(0)
	for col := 0; col < maxNumberOfColumns; col++ {
		colHeight := uint(0)
		for row := 0; row < numberOfRows; row++ {
			if len(imagesSize[row]) > col {
				colHeight += imagesSize[row][col].height

			}
		}

		if colHeight > maxHeight {
			maxHeight = colHeight
		}
	}

	output := drawImagesOnBackgroundInParallel(numberOfRows, maxWidth, maxHeight, maxNumberOfColumns, imagesMatrix, desiredWidth)

	return output
}

func calculateImageStartingPointAndSize(img image.Image, imagesMatrix [][]image.Image, padding int, desiredWidth int) (ImagePositionAndSize, error) {
	sp_y := padding
	for row := range imagesMatrix {
		sp_x := padding
		calculatedColumnWidth := math.Floor(float64(desiredWidth) / float64(len(imagesMatrix[row])))
		rowHeight := 0

		for col := range imagesMatrix[row] {
			originalWidth := float64(Width(imagesMatrix[row][col]))
			originalHeight := float64(Height(imagesMatrix[row][col]))
			resizeFactor := calculatedColumnWidth / originalWidth

			w := uint(originalWidth * resizeFactor)
			h := uint(originalHeight * resizeFactor)

			if imagesMatrix[row][col] == img {
				return ImagePositionAndSize{image.Point{sp_x, sp_y}, Size{w, h}}, nil
			} else {
				sp_x += int(w) + padding
			}

			if int(h) > rowHeight {
				rowHeight = int(h)
			}
		}

		sp_y += rowHeight + padding
	}

	return ImagePositionAndSize{image.Point{-1, -1}, Size{0, 0}}, errors.New("Image not found in matrix")
}

func drawSingleImageOnBackground(img image.Image, imagesMatrix [][]image.Image, padding int, desiredWidth int, background *MyImage) {
	imageDetails, _ := calculateImageStartingPointAndSize(img, imagesMatrix, padding, desiredWidth)
	sp := imageDetails.sp
	size := imageDetails.size

	background.drawRaw(img, sp, size.width, size.height)
}

func drawImagesOnBackgroundInParallel(numberOfRows int, maxWidth uint, maxHeight uint, maxNumberOfColumns int, imagesMatrix [][]image.Image, desiredWidth int) *MyImage {
	padding := 1
	rectangleEnd := image.Point{int(maxWidth) + (maxNumberOfColumns-1)*padding + 2*padding, int(maxHeight) + (numberOfRows-1)*padding + 2*padding}

	output := MyImage{image.NewRGBA(image.Rectangle{image.ZP, rectangleEnd})}

	for r := range imagesMatrix {
		for c := range imagesMatrix[r] {
			go drawSingleImageOnBackground(imagesMatrix[r][c], imagesMatrix, padding, desiredWidth, &output)
		}
	}

	return &output
}

// imagecollager will make a collage of images by combining them onto black background
// Script parameters are:
// 1. image shape - share for each inner image inside background - 'Rectangle' or 'Circle'
// 2. number of rows in which images are displayed
// 3. path to the directory where images are stored on file system

func loadImageChannelNew(path string, info os.FileInfo, e error, images chan image.Image, errors chan error) {
	if e != nil {
		errors <- e
		return
	}

	if !info.IsDir() {
		fimg, _ := os.Open(path)
		defer fimg.Close()
		img, _, imageError := image.Decode(fimg)

		if imageError == nil {
			images <- img
		} else {
			errors <- imageError
		}
	}
}

func countFiles(dirPath string) (int, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return 0, err
	}

	counter := 0
	for _, file := range files {
		if !file.IsDir() {
			counter++
		}
	}

	return counter, nil
}

func loadImagesChannel(dirName string, images chan image.Image, quit chan int, errors chan error) {
	err := filepath.Walk(dirName, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			errors <- e
		}

		if !info.IsDir() {
			fimg, _ := os.Open(path)
			defer fimg.Close()
			img, _, imageError := image.Decode(fimg)
			if imageError == nil {
				images <- img
			}
		}
		return nil
	})
	if err != nil {
		errors <- err
	} else {
		quit <- 1
	}
}

func CreateTestWallpaper(returnFrames []image.Image) {

	output := makeImageCollage(12000, 2, returnFrames...)
	f, err := os.Create("./core/frames/samples/test.png")
	if err != nil {
		// Handle error
	}
	defer f.Close()

	// Specify the quality, between 0-100
	// Higher is better
	opt := jpeg.Options{
		Quality: 100,
	}
	err = jpeg.Encode(f, output, &opt)
	if err != nil {
		fmt.Println(err)
	}
	imview.Show(output.value)

}
