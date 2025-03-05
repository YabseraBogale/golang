package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("question.json")
	if err != nil {
		log.Println(err)
	}
	for _, i := range file {
		fmt.Println(string(i))
	}

}
