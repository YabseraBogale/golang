package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	file, err := os.ReadFile("mcqs.txt")
	if err != nil {
		log.Println(err)
	}

	text := string(file)
	patter := regexp.MustCompile(`| P a g e`)
	fmt.Println(patter.FindAllStringIndex(text, -1))

}
