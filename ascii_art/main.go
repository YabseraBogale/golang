package main

import (
	"fmt"
	"image"
	"os"
)

func main() {
	file, err := os.Open("./wall.png")
	if err != nil {

	}
	defer file.Close()
	_, filestring, err := image.Decode(file)
	fmt.Println(filestring)
}
