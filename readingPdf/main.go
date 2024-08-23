package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("simple.pdf")
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range file {
		fmt.Println(i)
	}
}
