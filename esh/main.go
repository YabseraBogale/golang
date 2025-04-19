package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/add_job", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/add_employee", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/add_item", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/add_candate", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/add_emergency_contact", func(w http.ResponseWriter, r *http.Request) {

	})
	err = http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
