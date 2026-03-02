package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/YabseraBogale/golang/sherlock/site"
)

func main() {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(site.Data)

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

		var site site.SiteData

		if err := json.Unmarshal(value, &site); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(site.URL)

	}

}
