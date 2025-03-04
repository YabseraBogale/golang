package main

import (
	"encoding/json"
	"fmt"
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

func extractMCQs(text string) []MCQ {
	var mcqs []MCQ
	pattern := regexp.MustCompile(`(?m)(?P<question>\d+\.\s.*?\?)\s*\n\s*(?P<choices>(?:[a-d]\..*?\n)+)\s*Answer:\s*(?P<answer>[a-d])\..*?\n\s*Explanation:\s*(?P<explanation>.*?)\n\n`)

	matches := pattern.FindAllStringSubmatch(text, -1)
	if matches == nil {
		return mcqs
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

		mcqs = append(mcqs, MCQ{
			Question:    strings.TrimSpace(match[1]),
			Choices:     choices,
			Answer:      strings.TrimSpace(match[3]),
			Explanation: strings.TrimSpace(match[4]),
		})
	}
	return mcqs
}

func main() {
	fileContent, err := os.ReadFile("mcqs.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(len(fileContent))
	mcqs := extractMCQs(string(fileContent))
	jsonData, err := json.MarshalIndent(mcqs, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	os.WriteFile("mcqs.json", jsonData, 0644)
	fmt.Println("MCQs extracted and saved as mcqs.json")
}
