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
			username := r.FormValue("username")
			password := r.FormValue("password")
			fmt.Println(username, password)
		} else if r.Method == http.MethodGet {
			t, _ := template.ParseFiles("public/login.html")
			t.Execute(w, nil)
		}

	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")
			re_password := r.FormValue("re_password")
			if re_password != password {
				t, _ := template.ParseFiles("public/signup.html")
				t.Execute(w, struct{ NotMatched bool }{NotMatched: true})
			} else {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
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
