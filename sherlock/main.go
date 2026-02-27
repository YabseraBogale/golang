package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "https://wakatime.com/@"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)
}
