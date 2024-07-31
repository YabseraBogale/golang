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
			R, B, G, _ := color.RGBA()
			average := 0.299*R + 0.587*G + 0.114*B
			print(average, " ")
		}
		println()
	}

}
