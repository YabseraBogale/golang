package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("public/index.html")
		t.Execute(w, nil)
	})
	err := http.ListenAndServe(":8020", nil)
	if err != nil {
		log.Println(err)
	}
}
