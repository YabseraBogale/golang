package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	r float32 = 0.299
	g float32 = 0.587
	b float32 = 0.114
)

func main() {
	file, err := os.Open("./wall.png")
	if err != nil {

	}

	fileimage, err := png.Decode(file)
	b := fileimage.Bounds()
	newimage := image.NewRGBA(b)
	for x := 0; x < fileimage.Bounds().Max.X; x++ {
		for y := 0; y < fileimage.Bounds().Max.Y; y++ {
			oldpx := fileimage.At(x, y)
			r,b,g,_:=color.Gray16Model.Convert(oldpx).RGBA()
			print(r,g,b)
			newimage.Set(x, y, color.Gray16Model.Convert(oldpx))
		}
	}
	outFile, err := os.Create("changed.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	png.Encode(outFile, newimage)

}
