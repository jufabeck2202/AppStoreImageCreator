package core

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
)

func LoadImageChannel(pathPicture string, images chan image.Image, errors chan error) {
	file, err := os.Open(pathPicture)
	if err != nil {
		log.Printf("failed to open: %s", err)
	}
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		errors <- err
	}

	// Reset the read pointer if necessary.
	file.Seek(0, 0)
	contentType := http.DetectContentType(buffer)
	switch contentType {
	case "image/png":
		im, err := png.Decode(file)
		if err != nil {
			log.Printf("failed to decode: %s", err)
		}

		if err == nil {
			images <- im
		} else {
			errors <- err
		}
	case "image/jpeg":
		im, err := jpeg.Decode(file)
		if err != nil {
			log.Printf("failed to decode: %s", err)
		}
		if err == nil {
			images <- im
		} else {
			errors <- err
		}
	default:
		errors <- fmt.Errorf("Image not PNG or Jpeg %g")
	}
}
