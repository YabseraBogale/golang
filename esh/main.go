package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/jackc/pgx/v5"
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

	config, err := pgx.ParseConfig(os.Getenv("postgres"))
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(context.Background())

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
		if r.Method == "POST" {
			job_title := r.PostFormValue("job_title")
			department := r.PostFormValue("department")
			job_description := r.PostFormValue("job_description")
			_, err := conn.Exec(context.Background(), `Insert into Job(job_title,Department,job_description) values($1,$2,$3)`, job_title, department, job_description)
			if err != nil {
				log.Fatalln(err)
			}
		}

		err = templates.ExecuteTemplate(w, "add_job.html", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})

	http.HandleFunc("/job_title", func(w http.ResponseWriter, r *http.Request) {
		row, err := conn.Query(context.Background(), "Select job_title from job")
		if err != nil {
			log.Fatalln(err)
		}
		defer row.Close()
		job_title := []string{}
		for row.Next() {
			var title string
			if err := row.Scan(&title); err != nil {
				log.Fatalln(err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			job_title = append(job_title, title)
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(job_title); err != nil {
			log.Printf("Error during row iteration: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		row, err := conn.Query(context.Background(), "Select department from job")
		if err != nil {
			log.Fatalln(err)
		}
		defer row.Close()
		departments := []string{}
		for row.Next() {
			var department string
			if err := row.Scan(&department); err != nil {
				log.Fatalln(err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			departments = append(departments, department)

		}
		if err := row.Err(); err != nil {
			log.Printf("Error during row iteration: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(departments); err != nil {
			log.Printf("Error during row iteration: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/add_employee", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			firstname := r.PostFormValue("firstname")
			middlename := r.PostFormValue("middlename")
			lastname := r.PostFormValue("lastname")
			phonenumber := r.PostFormValue("phonenumber")
			fyda_id := r.PostFormValue("fyda_id")
			email := r.PostFormValue("email")
			department := r.PostFormValue("department")
			job_title := r.PostFormValue("job_title")
			hire_date, err := time.Parse("2006-01-02", r.PostFormValue("hire_date"))
			if err != nil {
				log.Fatalln(err)
			}
			emergency_firstname := r.PostFormValue("emergency_firstname")
			emergency_middlename := r.PostFormValue("emergency_middlename")
			emergency_lastname := r.PostFormValue("emergency_lastname")
			emergency_phonenumber := r.PostFormValue("emergency_phonenumber")
			emergency_email := r.PostFormValue("emergency_email")
			emergency_fyda_id := r.PostFormValue("emergency_fyda_id")
			_, err = conn.Exec(context.Background(), `Insert into Employee(firstname, middlename, lastname, phonenumber, fyda_id, email, 
					department, job_title, hire_date, emergency_firstname,
					emergency_middlename, emergency_lastname, emergency_phonenumber,
					emergency_email, emergency_fyda_id) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`,
				firstname, middlename, lastname, phonenumber, fyda_id, email,
				department, job_title, hire_date, emergency_firstname,
				emergency_middlename, emergency_lastname, emergency_phonenumber,
				emergency_email, emergency_fyda_id)
			if err != nil {
				log.Fatalln(err)
			}
		}
		err = templates.ExecuteTemplate(w, "add_employee.html", nil)
		if err != nil {
			log.Fatalln(err)
		}

	})

	http.HandleFunc("/add_item", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {

			employee_id := r.PostFormValue("employee_id")

			item_name := r.PostFormValue("item_name")
			item_description := r.PostFormValue("item_description")
			quantity := r.PostFormValue("quantity")
			item_status := r.PostFormValue("item_status")
			item_date, err := time.Parse("2006-01-02", r.PostFormValue("item_date"))
			if err != nil {
				log.Fatalln(err)
			}
			_, err = conn.Exec(context.Background(), `Insert into item(employee_id, item_name, item_description,
								quantity, item_status, item_date) values($1,$2,$3,$4,$5,$6)`,
				employee_id, item_name, item_description, quantity, item_status, item_date)
			if err != nil {
				log.Fatalln(err)
			}
		}

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

	err = http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
