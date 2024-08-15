package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("Hotmail_german_cloud.txt")
	if err != nil {
		log.Fatalln(err)
	}

	for _, i := range file {
		fmt.Println(i)

	}
}
