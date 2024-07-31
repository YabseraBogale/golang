package main

import (
	"fmt"
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
	defer file.Close()
	fileimage, err := png.Decode(file)
	ascii := make([][]byte, fileimage.Bounds().Max.X)
	for x := 0; x < fileimage.Bounds().Max.X; x++ {
		ascii[x] = make([]byte, fileimage.Bounds().Max.Y)
		for y := 0; y < fileimage.Bounds().Max.Y; y++ {
			oldcolor := fileimage.At(x, y)
			R, B, G, _ := oldcolor.RGBA()
			average := r*float32(R) + g*float32(G) + b*float32(B)
			if uint8(average) == 0 {
				ascii[x][y] = byte('O')
			} else if uint8(average) == 255 {
				ascii[x][y] = byte('X')
			} else if uint8(average) < 150 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x][y] = byte('W')
			} else if uint8(average) > 150 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x][y] = byte('Y')
			}

		}
	}

	fmt.Println(ascii)

}
