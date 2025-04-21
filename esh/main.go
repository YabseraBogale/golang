package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var templates *template.Template

func init() {
	// Parse all HTML files in the "templates" directory
	var err error
	templates, err = template.ParseGlob(filepath.Join("public", "*.html"))
	if err != nil {
		log.Fatal(err)
	}
}
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

		err = templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/add_job", func(w http.ResponseWriter, r *http.Request) {
		job_title := r.PostFormValue("job_title")
		job_description := r.PostFormValue("job_description")
		// don't forget to delete the fmt line below
		fmt.Println(job_title, job_description)
		err = templates.ExecuteTemplate(w, "add_job.html", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/add_employee", func(w http.ResponseWriter, r *http.Request) {
		err = templates.ExecuteTemplate(w, "add_employee.html", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/add_item", func(w http.ResponseWriter, r *http.Request) {
		employee_id := r.PostFormValue("employee_id")
		item_id := r.PostFormValue("item_id")
		item_name := r.PostFormValue("item_name")
		item_description := r.PostFormValue("item_description")
		quantity := r.PostFormValue("quantity")
		item_status := r.PostFormValue("item_status")
		item_date := r.PostFormValue("item_date")
		fmt.Println(item_date, employee_id, item_id, item_description, item_name, quantity, item_status)
		// don't forget to delete the fmt line below
		err = templates.ExecuteTemplate(w, "add_item.html", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/add_candate", func(w http.ResponseWriter, r *http.Request) {
		err = templates.ExecuteTemplate(w, "add_candate.html", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/add_emergency_contact", func(w http.ResponseWriter, r *http.Request) {
		err = templates.ExecuteTemplate(w, "add_emergency_contact.html", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	err = http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
