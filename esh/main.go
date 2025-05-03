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

	"github.com/YabseraBogale/golang/esh/database"
	"github.com/jackc/pgx/v5"
)

var templates *template.Template

func init() {

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
	defer logFile.Close() // Close the log file after use.
	log.SetOutput(logFile)

	config, err := pgx.ParseConfig(os.Getenv("postgres"))
	if err != nil {
		log.Println(err)
	}
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close(context.Background()) // Close the database after use.

	// --- Static File Handling ---
	// Serve CSS and JavaScript files.
	// `http.FileServer` creates a handler that serves files from the given directory.

	css := http.FileServer(http.Dir("./public/style"))
	js := http.FileServer(http.Dir("./public/js"))

	http.Handle("/style/", http.StripPrefix("/style/", css))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		err = templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			log.Println(err)
		}
	})
	// API for human resources to add job
	http.HandleFunc("/add_job", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			job_title := r.PostFormValue("job_title")
			department := r.PostFormValue("department")
			job_description := r.PostFormValue("job_description")
			_, err := conn.Exec(context.Background(), `Insert into Job(job_title,Department,job_description) values($1,$2,$3)`, job_title, department, job_description)
			if err != nil {
				log.Println(err)
			}
		}

		err = templates.ExecuteTemplate(w, "add_job.html", nil)
		if err != nil {
			log.Println(err)
		}
	})

	// API Endpoint to get Job Titles
	http.HandleFunc("/job_title", func(w http.ResponseWriter, r *http.Request) {
		// Query the database to get all job titles.
		row, err := conn.Query(context.Background(), "Select job_title from job")
		if err != nil {
			log.Println(err)
		}
		defer row.Close()       // Close the rows after use.
		job_title := []string{} // Set the Content-Type header to indicate JSON
		for row.Next() {
			var title string
			if err := row.Scan(&title); err != nil {
				log.Println(err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			job_title = append(job_title, title)
		}

		// Set the Content-Type header to indicate JSON
		w.Header().Set("Content-Type", "application/json")
		// Encode the job titles as JSON and send the response.
		if err := json.NewEncoder(w).Encode(job_title); err != nil {
			log.Printf("Error during row iteration: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// API Endpoint to get Departments
	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		// Query the database to get all departments.
		row, err := conn.Query(context.Background(), "Select department from job")
		if err != nil {
			log.Println(err)
		}
		defer row.Close() // Close the rows after use.
		departments := []string{}
		for row.Next() {
			var department string
			if err := row.Scan(&department); err != nil {
				log.Println(err)
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

	// Display the "add_employee.html" form.
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
				log.Println(err)
			}
			salary := r.PostFormValue("salary")
			emergency_firstname := r.PostFormValue("emergency_firstname")
			emergency_middlename := r.PostFormValue("emergency_middlename")
			emergency_lastname := r.PostFormValue("emergency_lastname")
			emergency_phonenumber := r.PostFormValue("emergency_phonenumber")
			emergency_email := r.PostFormValue("emergency_email")
			emergency_fyda_id := r.PostFormValue("emergency_fyda_id")
			_, err = conn.Exec(context.Background(), `Insert into Employee(firstname, middlename, lastname, phonenumber, fyda_id, email, 
					department, job_title, hire_date,salary, emergency_firstname,
					emergency_middlename, emergency_lastname, emergency_phonenumber,
					emergency_email, emergency_fyda_id) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)`,
				firstname, middlename, lastname, phonenumber, fyda_id, email,
				department, job_title, hire_date, salary, emergency_firstname,
				emergency_middlename, emergency_lastname, emergency_phonenumber,
				emergency_email, emergency_fyda_id)
			if err != nil {
				log.Println(err)
			}
		}

		// Display the "add_employee.html" form.
		err = templates.ExecuteTemplate(w, "add_employee.html", nil)
		if err != nil {
			log.Println(err)
		}

	})
	http.HandleFunc("/item_track", func(w http.ResponseWriter, r *http.Request) {

	})

	// Add Item Page and API
	http.HandleFunc("/add_item", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {

			employee_id := r.PostFormValue("employee_id")
			item_name := r.PostFormValue("item_name")
			item_description := r.PostFormValue("item_description")
			quantity := r.PostFormValue("quantity")
			item_status := r.PostFormValue("item_status")
			item_date, err := time.Parse("2006-01-02", r.PostFormValue("item_date"))
			if err != nil {
				log.Println(err)
			}
			_, err = conn.Exec(context.Background(), `Insert into item(employee_id, item_name, item_description,
								quantity, item_status, item_date) values($1,$2,$3,$4,$5,$6)`,
				employee_id, item_name, item_description, quantity, item_status, item_date)
			if err != nil {
				log.Println(err)
			}
		}

		err = templates.ExecuteTemplate(w, "add_item.html", nil)
		if err != nil {
			log.Println(err)
		}
	})

	// API for purchase team to see approved item to be bought and then they will add it the store
	http.HandleFunc("/purchase_request/purchase_team", func(w http.ResponseWriter, r *http.Request) {
		// Query the database to get approved purchase requests.

		purchase_request_list := []database.PurchaseRequest{} // Slice to store purchase requests.
		row, err := conn.Query(context.Background(), `Select * from purchase_request where item_purchase_request='Approved'`)
		if err != nil {
			log.Println(err)
		}
		defer row.Close() // Close the rows after use.
		for row.Next() {
			var employeeid string
			var itemid string
			var itemname string
			var itemdescription string
			var itemquanitiy int
			var itemstatus string
			var itempurchaserequest string
			var itemdate string
			err = row.Scan(&employeeid, &itemid, &itemname, &itemdescription, &itemquanitiy, &itemstatus, &itempurchaserequest, &itemdate)
			if err != nil {
				log.Println(err)
			}
			date, err := time.Parse("2006-01-02", itemdate)
			if err != nil {
				log.Println(err)
			}
			purchase_request_list = append(purchase_request_list, database.InsertPurchaseRequest(employeeid, itemid, itemname, itemdescription,
				itemquanitiy, itemstatus, itempurchaserequest, date))

		}
		// go to frontend to list purchase request list
		// Display the "purchase_request_manager.html" template with the list of requests.
		err = templates.ExecuteTemplate(w, "purchase_request_manager.html", purchase_request_list)
		if err != nil {
			log.Println(err)
		}
	})

	// API for Managers to see purchase request to be approved
	http.HandleFunc("/purchase_request/manager", func(w http.ResponseWriter, r *http.Request) {
		// Query the database to get purchase requests that are "To be Approved".
		purchase_request_list := []database.PurchaseRequest{} // Slice to store purchase requests.
		row, err := conn.Query(context.Background(), `Select * from purchase_request where item_purchase_request='To be Approved'`)
		if err != nil {
			log.Println(err)
		}
		defer row.Close() // Close the rows after use.
		for row.Next() {
			var employeeid string
			var itemid string
			var itemname string
			var itemdescription string
			var itemquanitiy int
			var itemstatus string
			var itempurchaserequest string
			var itemdate string
			err = row.Scan(&employeeid, &itemid, &itemname, &itemdescription, &itemquanitiy, &itemstatus, &itempurchaserequest, &itemdate)
			if err != nil {
				log.Println(err)
			}
			date, err := time.Parse("2006-01-02", itemdate)
			if err != nil {
				log.Println(err)
			}
			purchase_request_list = append(purchase_request_list, database.InsertPurchaseRequest(employeeid, itemid, itemname, itemdescription,
				itemquanitiy, itemstatus, itempurchaserequest, date))

		}

		err = templates.ExecuteTemplate(w, "purchase_request_manager.html", purchase_request_list)
		if err != nil {
			log.Println(err)
		}
	})

	// Purchase Request -  Create a New Request
	http.HandleFunc("/purchase_request", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Process form submission to create a new purchase request.
			employee_id := r.PostFormValue("employee_id")
			item_name := r.PostFormValue("item_name")
			item_description := r.PostFormValue("item_description")
			quantity := r.PostFormValue("quantity")
			item_status := r.PostFormValue("item_status")
			item_purchase_request := "To be Approved"
			item_date, err := time.Parse("2006-01-02", r.PostFormValue("item_date"))
			if err != nil {
				log.Println(err)
			}
			_, err = conn.Exec(context.Background(), `Insert into purchase_request(employee_id, item_name, item_description,
								quantity, item_status, item_purchase_request,item_date) values($1,$2,$3,$4,$5,$6,$7)`,
				employee_id, item_name, item_description, quantity, item_status, item_purchase_request, item_date)
			if err != nil {
				log.Println(err)
			}
		}

		// Display the "purchase_request.html" form.
		err := templates.ExecuteTemplate(w, "purchase_request.html", nil)
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/add_candate", func(w http.ResponseWriter, r *http.Request) {
		err = templates.ExecuteTemplate(w, "add_candate.html", nil)
		if err != nil {
			log.Println(err)
		}
	})

	// --- Start the HTTP Server ---
	// Start the web server and listen for incoming requests on port 8080.
	err = http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Println(err)
	}

}
