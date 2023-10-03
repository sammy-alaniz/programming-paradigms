package main

import (
	"fmt"

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
	fmt.Println("initalizing connection data")

	cd.uuid = uuid.New().String()
	fmt.Println(cd.uuid)

	cd.driver, cd.err = neo4j.NewDriver(cd.uri, neo4j.BasicAuth(cd.username, cd.password, ""))

	//defer cd.driver.Close()

	if cd.err != nil {
		panic(cd.err)
	}

	fmt.Println("Driver established!")

	cd.session = cd.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})

	//defer cd.session.Close()

	if cd.session == nil {
		fmt.Println("Session failed had an error")
	}

	//cd.query("MATCH (a) RETURN a")

	fmt.Println("Session established!")
}

func (cd *ConnectionData) query(query string) (hash string, results map[string]interface{}, err error) {

	fmt.Println("query hit!")
	fmt.Println("Data", query)

	rs, er := cd.session.Run(query, nil)

	fmt.Println("Query ran!")

	if er != nil {
		fmt.Print("Query had an error! Error: ", er)
		return cd.uuid, nil, er
	}

	fmt.Println("Query did not have an error!")

	results = cd.neo4j_to_http(rs)

	return cd.uuid, results, er
}

// this runs but is untested
func (cd *ConnectionData) neo4j_to_http(data neo4j.Result) (rtn_data map[string]interface{}) {

	fmt.Println("neo4j_to_http hit!!")
	fmt.Println("Data: ", data)

	rtn_data = make(map[string]interface{})

	for data.Next() {
		record := data.Record()

		for _, key := range record.Keys {
			value, valid := record.Get(key)

			if !valid {
				// maybe do more error handling?
				continue
			}

			rtn_data[key] = value

			fmt.Printf("Key: %s, Value: %v\n", key, value)
		}

	}

	return rtn_data
}

func (cd *ConnectionData) shutdown() {
	if cd.driver == nil {
		fmt.Println("Driver is nil")
	} else {
		cd.driver.Close()
	}
}
