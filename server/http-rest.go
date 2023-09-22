package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func start() string {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(":8080", nil)
	return "yup"
}

func start_http_request_handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/data", processData).Methods("POST")
	router.HandleFunc("/query", processQuery).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func processData(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	go handleData(data)
	w.WriteHeader(http.StatusAccepted)
}

func processQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processQuery hit!")
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("error occured", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	go handleQuery(data)
	w.WriteHeader(http.StatusAccepted)
}

func handleQuery(data map[string]interface{}) {
	log.Println("Handling data:", data)

}

func handleData(data map[string]interface{}) {
	log.Println("Handling data:", data)
}
