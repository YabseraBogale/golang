package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
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
	var wg sync.WaitGroup
	for i := 1; i < len(os.Args); i++ {
		wg.Add(1)
		go func(username string) {
			defer wg.Done()
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
					if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
						continue
					}
					continue // Move to the next site since 'c' is nil
				}
				if c.StatusCode == 200 {
					data, err := io.ReadAll(c.Body)
					if err != nil {
						fmt.Println(err)
					}
					if strings.Contains(string(data), os.Args[i]) {
						fmt.Println(key, ":", os.Args[i])
					}
				}
				c.Body.Close()

			}
		}(os.Args[i])
	}
	wg.Wait()
}
