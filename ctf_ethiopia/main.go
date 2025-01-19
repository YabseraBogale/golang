package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

		} else if r.Method == http.MethodGet {
			t, _ := template.ParseFiles("public/login.html")
			t.Execute(w, nil)
		}

	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

		} else if r.Method == http.MethodGet {
			t, _ := template.ParseFiles("public/signup.html")
			t.Execute(w, nil)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("public/home.html")
		t.Execute(w, nil)
	})

	fmt.Println("Listing on http://localhost:5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Println(err)
	}
}
