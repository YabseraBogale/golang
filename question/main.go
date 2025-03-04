package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		log.Println(err)
	}

	text := string(file)
	patter := regexp.MustCompile(`(?s)\.(.*?)(?:\n(?:[a-d]\..*?\n?)+)\s*Answer:\s*([a-d]\..*?)\nExplanation:\s*(.*?)(?:\n\n|\z)`)
	matches := patter.FindAllStringSubmatch(text, -1)
	if matches == nil {
		fmt.Println("No matches found")
		return
	}

	// Loop through all matched MCQs
	for i, match := range matches {
		fmt.Printf("Question %d: %s\n", i+1, strings.TrimSpace(match[1]))
	}
}
