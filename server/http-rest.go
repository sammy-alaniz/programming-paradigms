package main

import (
	"encoding/json"
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
	log.Println("http-rest.processQuery hit!")
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Println("error occured", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//hr.handleQuery(data)
	uuid, results, er := hr.cd.query(data["statement"].(string))
	log.Println("UUID : ", uuid)
	log.Println("RESULTS : ", results)
	log.Println("ERROR : ", er)

	w.WriteHeader(http.StatusAccepted)
}

func (hr *HttpRest) handleQuery(data map[string]interface{}) {
	log.Println("http-rest.handleQuery hit!")
	log.Println("http-rest.Handling data:", data)
	hr.cd.query(data["statement"].(string))
}
