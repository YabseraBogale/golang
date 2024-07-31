package main

import (
	"image/png"
	"os"
)

func main() {
	file, err := os.Open("./wall.png")
	if err != nil {

	}

	fileimage, err := png.Decode(file)

	for x := 0; x < fileimage.Bounds().Max.X; x++ {
		for y := 0; y < fileimage.Bounds().Max.Y; y++ {
			color := fileimage.At(x, y)
			print(color.RGBA())
		}
		println()
	}

}
