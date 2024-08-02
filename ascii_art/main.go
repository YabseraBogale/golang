package main

import (
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
	defer file.Close()
	fileimage, err := png.Decode(file)
	ascii := make([][]byte, fileimage.Bounds().Max.Y)
	for x := 0; x < fileimage.Bounds().Max.Y; x++ {
		ascii[x] = make([]byte, fileimage.Bounds().Max.X)
		for y := 0; y < fileimage.Bounds().Max.X; y++ {
			oldcolor := fileimage.At(x, y)
			R, B, G, _ := oldcolor.RGBA()
			average := r*float32(R) + g*float32(G) + b*float32(B)
			if uint8(average) == 0 {
				ascii[x][y] = byte('0')
			} else if uint8(average) == 255 {
				ascii[x][y] = byte('1')
			} else if uint8(average) < 50 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x][y] = byte('#')
			} else if uint8(average) >= 50 && uint8(average) < 87 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x][y] = byte('(')
			} else if uint8(average) >= 87 && uint8(average) < 162 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x][y] = byte(')')
			} else if uint8(average) >= 162 && uint8(average) < 209 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x][y] = byte('_')
			} else if uint8(average) >= 209 && uint8(average) < 255 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x][y] = byte('$')
			}

		}
	}

	fs, err := os.Create("image.txt")
	if err != nil {
		log.Println(err)
	}
	defer fs.Close()
	for _, i := range ascii {
		fs.Write(i)
	}

}
