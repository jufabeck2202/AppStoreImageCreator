package main
/*
#cgo CGO_CFLAGS_ALLOW=-Wl,(-framework|CoreFoundation)
*/

import (
	"fmt"
	"os"
	"github.com/h2non/bimg"
  )

func main(){
	buffer, err := bimg.Read("test.png")
	if err != nil {
	  fmt.Fprintln(os.Stderr, err)
	}
	
	newImage, err := bimg.NewImage(buffer).Resize(100, 600)
	if err != nil {
	  fmt.Fprintln(os.Stderr, err)
	}
	
	size, err := bimg.NewImage(newImage).Size()
	if size.Width == 800 && size.Height == 600 {
	  fmt.Println("The image size is valid")
	}
	
	bimg.Write("new.jpg", newImage)
}
