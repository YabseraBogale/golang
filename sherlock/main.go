package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/YabseraBogale/golang/sherlock/site"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no username provided")
		return
	}
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

	for i := 1; i < len(os.Args); i++ {
		for key, value := range full_map {

			if key == "$schema" {
				continue
			}

			var site site.SiteData

			if err := json.Unmarshal(value, &site); err != nil {
				fmt.Println(err)
				continue
			}
			username_url := strings.Replace(site.URL, "{}", os.Args[i], -1)
			c, err := client.Get(username_url)
			if err != nil {
				log.Println(err)
			}
			defer c.Body.Close()
			if c.StatusCode == 200 {
				fmt.Println(username_url)
			}
		}
	}

}
