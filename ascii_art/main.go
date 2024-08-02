package main

import (
	"image/png"
	"os"
	"strings"
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
	ascii := make([]string, fileimage.Bounds().Max.X)
	for x := 0; x < fileimage.Bounds().Max.X; x++ {
		ascii[x] = ""
		for y := 0; y < fileimage.Bounds().Max.Y; y++ {
			oldcolor := fileimage.At(x, y)
			R, B, G, _ := oldcolor.RGBA()
			average := r*float32(R) + g*float32(G) + b*float32(B)
			if uint8(average) == 0 {
				ascii[x] += "#"
			} else if uint8(average) == 255 {
				ascii[x] += "[]"
			} else if uint8(average) < 50 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x] += "-"
			} else if uint8(average) >= 50 && uint8(average) < 87 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x] += "("
			} else if uint8(average) >= 87 && uint8(average) < 162 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x] += "+"
			} else if uint8(average) >= 162 && uint8(average) < 209 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x] += "%"
			} else if uint8(average) >= 209 && uint8(average) < 255 && uint8(average) != 0 && uint8(average) != 255 {
				ascii[x] += "$"
			}

		}
	}
	for j, i := range ascii {
		ascii[j] = transforme(i)
	}

	for _, i := range ascii {
		println(i)
	}
	// fs, err := os.Create("image.txt")
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer fs.Close()
	// for _, i := range ascii {

	// }

}

func transforme(oldstr string) string {
	var newstr string
	resize := 2
	for i := 0; i < strings.Count(oldstr, "-")/resize; i++ {
		newstr += "-"
	}
	for i := 0; i < strings.Count(oldstr, "[]")/resize; i++ {
		newstr += "[]"
	}
	for i := 0; i < strings.Count(oldstr, "#")/resize; i++ {
		newstr += "#"
	}
	for i := 0; i < strings.Count(oldstr, "+")/resize; i++ {
		newstr += "+"
	}
	for i := 0; i < strings.Count(oldstr, "%")/resize; i++ {
		newstr += "%"
	}
	for i := 0; i < strings.Count(oldstr, "$")/resize; i++ {
		newstr += "$"
	}
	if len(newstr) < 108 {
		for i := len(newstr); i < 108; i++ {
			newstr += "$"
		}
	}
	return newstr
}
