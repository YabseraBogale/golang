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

type SteData struct {
	ErrorMsg        interface{} `json:"errorMsg,omitempty"`
	ErrorType       string      `json:"errorType"`
	IsNSFW          bool        `json:"isNSFW,omitempty"`
	URL             string      `json:"url"`
	URLMain         string      `json:"urlMain"`
	UsernameClaimed string      `json:"username_claimed"`
}

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
	err = json.Unmarshal(body_byte, &full_map)

	if err != nil {
		log.Fatalln(err)
	}

	for key, value := range full_map {

		if key == "$schema" {
			continue
		}

		var site SteData

		if err := json.Unmarshal(value, &site); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(site.URL)

	}

}
