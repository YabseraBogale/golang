package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadDir(".")
	if err != nil {

	}
	fmt.Println(file)
}
