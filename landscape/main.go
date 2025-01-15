package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello\n")
	})
	err := http.ListenAndServe(":8020", nil)
	if err != nil {
		log.Println(err)
	}
}
