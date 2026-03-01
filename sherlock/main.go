package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const data = "https://raw.githubusercontent.com/sherlock-project/sherlock/master/sherlock_project/resources/data.json"

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(data)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalln(resp.StatusCode)
	}

	body_byte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var full_map map[string]json.RawMessage
	json.Unmarshal(body_byte, full_map)

	for key, _ := range full_map {
		if key == "$schema" {
			continue
		} else {
			fmt.Println(key)
		}
	}

}
