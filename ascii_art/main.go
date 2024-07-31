package main

import (
	"image/png"
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
	countWhite := 0
	countBlack := 0
	other := 0
	for x := 0; x < fileimage.Bounds().Max.X; x++ {
		for y := 0; y < fileimage.Bounds().Max.Y; y++ {
			oldcolor := fileimage.At(x, y)
			R, B, G, _ := oldcolor.RGBA()
			average := r*float32(R) + g*float32(G) + b*float32(B)
			if uint8(average) == 0 {
				countWhite += 1
			} else if uint8(average) == 255 {
				countBlack += 1
			} else {
				other += 1
			}
		}
	}
	println("there is", countBlack, "black pixel and", countWhite, "white pixel in the image and other pixel is", other)

}
