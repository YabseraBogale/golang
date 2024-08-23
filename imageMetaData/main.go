package main

import (
	"fmt"
	"log"
	"os"

	"github.com/evanoberholster/imagemeta"
)

func main() {
	file, err := os.Open("image.jpg")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	img, err := imagemeta.Decode(file)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(img.String())

}
