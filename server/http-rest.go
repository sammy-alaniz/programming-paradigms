package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpRest struct {
	cd *ConnectionData
}

func (hr *HttpRest) start_http_request_handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/query", hr.processQuery).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func (hr *HttpRest) processQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processQuery hit!")
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("error occured", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	go hr.handleQuery(data)
	w.WriteHeader(http.StatusAccepted)
}

func (hr *HttpRest) handleQuery(data map[string]interface{}) {
	log.Println("Handling data:", data)
	hr.cd.query(data["statement"].(string))
}
