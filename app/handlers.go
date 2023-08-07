package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"nombre"`
	City    string `json:"ciudad"`
	Zipcode string `json:"zip"`
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Alvaro", "ciudad1", "123"},
		{"Juan", "ciudad1", "123"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
