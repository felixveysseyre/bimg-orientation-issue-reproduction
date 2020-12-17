package main

import (
	"fmt"
	"github.com/h2non/bimg"
	"io/ioutil"
	"os"
	"path"
)

func main()  {
	cD, _ := os.Getwd()
	fP := path.Join(cD, "data/Landscape_6.jpg")
	fB, _ := ioutil.ReadFile(fP)
	bI := bimg.NewImage(fB)
	iS, _ := bI.Size()
	iM, _ := bI.Metadata()

	fmt.Printf("Image size: %+v\n", iS) // 1200 x 1800
	fmt.Printf("Image orientation: %+v\n", iM.EXIF.Orientation) // 6

	options := bimg.Options{
		Top: 0,
		Left: 0,
		AreaWidth: 1000,
		AreaHeight: 1800,
	}

	_, err := bI.Process(options)

	fmt.Printf("Error: %+v\n", err) // extract_area: bad extract area
}