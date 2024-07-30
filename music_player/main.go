package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Music struct {
	Name string `json:"Name"`
}

func files(w http.ResponseWriter, r *http.Request) {
	fs, err := os.ReadDir("../../../../Music")
	if err != nil {
		log.Println(err)
	}
	data := []Music{}
	for _, i := range fs {
		data = append(data, Music{Name: i.Name()})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", files)
	http.ListenAndServe(":8080", nil)
}
