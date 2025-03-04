package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.ReadFile("question.json")
	if err != nil {
		log.Println(err)
	}

}
