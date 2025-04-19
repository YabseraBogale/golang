package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	css := http.FileServer(http.Dir("./public/style"))
	js := http.FileServer(http.Dir("./public/js"))

	http.Handle("/style/", http.StripPrefix("/style/", css))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("public/index.html")
		if err != nil {
			log.Fatalln(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/add_job", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("public/add_job.html")
		if err != nil {
			log.Fatalln(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/add_employee", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("public/add_employee.html")
		if err != nil {
			log.Fatalln(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/add_item", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("public/add_item.html")
		if err != nil {
			log.Fatalln(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/add_candate", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("public/add_candate.html")
		if err != nil {
			log.Fatalln(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/add_emergency_contact", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("public/add_emergency_contact.html")
		if err != nil {
			log.Fatalln(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	err = http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
