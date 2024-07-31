package main

import (
	"fmt"
	"image/png"
	"math/rand"
	"os"
	"strconv"
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
	ascii := make([][]string, fileimage.Bounds().Max.Y)
	for x := 0; x < fileimage.Bounds().Max.X; x++ {
		ascii[x] = make([]string, fileimage.Bounds().Max.X)
		for y := 0; y < fileimage.Bounds().Max.Y; y++ {
			oldcolor := fileimage.At(x, y)
			R, B, G, _ := oldcolor.RGBA()
			average := r*float32(R) + g*float32(G) + b*float32(B)
			if uint8(average) == 0 {
				ascii[x][y] = string('Ω')
			} else if uint8(average) == 255 {
				ascii[x][y] = string('ω')
			} else if uint8(average) < 127 || uint8(average) > 32 {
				ascii[x][y] = string(uint8(average))
			} else if uint8(average) >= 127 || uint8(average) <= 32 {
				ascii[x][y] = strconv.Itoa(rand.Intn(127-32) + 32)
			}
		}
	}
	fmt.Println(ascii)
}
