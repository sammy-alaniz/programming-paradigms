package main

import (
	"fmt"
	"encoding/json"
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

func start_two() {
	router := mux.NewRouter()
    router.HandleFunc("/data", processData).Methods("POST")
    router.HandleFunc("/datatwo", printHello).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func processData(w http.ResponseWriter, r *http.Request) {
    var data map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    go handleData(data) // Start a new goroutine to handle the data
    w.WriteHeader(http.StatusAccepted)
}

func printHello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("dataTwo")
}

func handleData(data map[string]interface{}) {
    // Simulate processing the data (in reality, this would be more complex)
    log.Println("Processing data:", data)
}