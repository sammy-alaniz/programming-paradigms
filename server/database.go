package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type ConnectionData struct {
	uri      string
	username string
	password string
	driver   neo4j.Driver
	session  neo4j.Session
	err      error
	uuid     string
}

func (cd *ConnectionData) init() {
	log.Println("database.initalizing connection data")

	cd.uuid = uuid.New().String()
	log.Println(cd.uuid)

	cd.driver, cd.err = neo4j.NewDriver(cd.uri, neo4j.BasicAuth(cd.username, cd.password, ""))

	if cd.err != nil {
		panic(cd.err)
	}

	log.Println("database.Driver established!")

	cd.session = cd.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})

	if cd.session == nil {
		log.Println("database.Session failed had an error")
	}

	log.Println("database.Session established!")
}

func (cd *ConnectionData) query(query string) (hash string, results map[string]interface{}, err error) {

	log.Println("database.query hit!")

	rs, er := cd.session.Run(query, nil)

	if er != nil {
		log.Print("Query had an error! Error: ", er)
		return cd.uuid, nil, er
	}

	results = cd.neo4j_to_http(rs)

	return cd.uuid, results, er
}

// this runs but is untested
func (cd *ConnectionData) neo4j_to_http(data neo4j.Result) (rtn_data map[string]interface{}) {

	log.Println("database.neo4j_to_http hit!!")

	rtn_data = make(map[string]interface{})

	for data.Next() {
		record := data.Record()
		log.Println("record loop")

		for _, key := range record.Keys {
			value, valid := record.Get(key)
			log.Println(("key loop"))

			if !valid {
				// maybe do more error handling?
				continue
			}

			rtn_data[key] = value

			log.Printf("Key: %s, Value: %v\n", key, value)
		}

	}

	return rtn_data
}

func (cd *ConnectionData) shutdown() {
	if cd.driver == nil {
		log.Println("Driver is nil")
	} else {
		cd.driver.Close()
	}
}
