package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type MCQ struct {
	Question    string            `json:"question"`
	Choices     map[string]string `json:"choices"`
	Answer      string            `json:"answer"`
	Explanation string            `json:"explanation"`
}

func main() {
	file, err := os.ReadFile("mcqs.txt")
	if err != nil {
		log.Println(err)
	}
	pattern := regexp.MustCompile(`(?m)(?P<question>\d+\.\s.*?\?)\s*\n\s*(?P<choices>(?:[a-d]\..*?\n)+)\s*Answer:\s*(?P<answer>[a-d])\..*?\n\s*Explanation:\s*(?P<explanation>.*?)\n\n`)

	text := string(file)
	matches := pattern.FindAllStringSubmatch(text, -1)
	if matches == nil {
		fmt.Println("no match")
	}
	for _, match := range matches {
		choicesText := strings.TrimSpace(match[2])
		choicesLines := strings.Split(choicesText, "\n")
		choices := make(map[string]string)
		for _, line := range choicesLines {
			if len(line) > 2 {
				choices[string(line[0])] = strings.TrimSpace(line[2:])
			}
		}

		fmt.Println(MCQ{
			Question:    strings.TrimSpace(match[1]),
			Choices:     choices,
			Answer:      strings.TrimSpace(match[3]),
			Explanation: strings.TrimSpace(match[4]),
		})
	}

}
