package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	//starting server
	log.Println(http.ListenAndServe("localhost:8000", router))
}
